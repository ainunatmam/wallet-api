package transaction

import (
	"database/sql"
	"errors"
	"log"
	"time"
	"wallet-api/app/entity"
	"wallet-api/app/presentation"

	"github.com/gofiber/fiber/v2"
)

func (s *transactionService) Confirmation(c *fiber.Ctx, req *presentation.TransactionConfirmationRequest) (err error) {
	
	tx, err := s.transactionManager.BeginTx(c.Context())
	if err != nil {
		return err
	}

	defer func() {
		commitErr := s.transactionManager.CommitOrRollback(err, tx)
		if commitErr != nil {
			log.Default().Println(commitErr)
		}
	}()
	

	transaction, err := s.transactionRepo.LockRowByTrxId(c.Context(), tx, req.TrxID)
	if err != nil {
		return err
	}

	if transaction == nil {
		return errors.New("Transaction not found")
	}

	if transaction.Status == string(entity.TransactionStatusSuccess) {
		return errors.New("Transaction already confirmed")
	}

	if transaction.Status == string(entity.TransactionStatusFailed) {
		return errors.New("Transaction already failed")
	}

	if transaction.TransactionType != string(entity.TransactionTypeTransfer) {
		var wallet *entity.Wallet
		wallet, err = s.walletRepo.LockRowById(c.Context(), tx, transaction.WalletID)
		if err != nil {
			return err
		}
	
		if wallet == nil {
			return errors.New("Wallet not found")
		}
	
		if wallet.Status != string(entity.WalletStatusActive) {
			return errors.New("Wallet is "+ wallet.Status)
		}
	
		var currency *entity.Currency
		currency, err = s.currencyRepo.FindByCode(c.Context(),wallet.Currency)
		if err != nil {
			return err
		}
	
		if currency == nil {
			return errors.New("Invalid Currency")
		}
		
		switch transaction.TransactionType {
		case string(entity.TransactionTypeTopUp):
			err = s.confirmationTopUp(c, transaction, wallet, currency, tx)
		case string(entity.TransactionTypePayment):
			err = s.confirmationPayment(c, transaction, wallet, currency, tx)
		default:
			err = errors.New("Invalid Transaction Type")
		}
	} else {
		var wallets *[]entity.Wallet
		wallets, err = s.walletRepo.LockRowByIds(c.Context(), tx, []uint64{transaction.WalletID, *transaction.TransferedWalletID})
		if err != nil {
			return err
		}
		if wallets == nil || len(*wallets) < 2 {
			return errors.New("Invalid Wallet")
		}

		senderWallet := entity.Wallet{}
		receiverWallet := entity.Wallet{}
		for _, w := range *wallets {
			if w.ID == transaction.WalletID {
				senderWallet = w
			}
			if w.ID == *transaction.TransferedWalletID {
				receiverWallet = w
			}
		}
		err = s.confirmationTransfer(c, transaction, &senderWallet, &receiverWallet, tx)
	}

	if err != nil {
		return err
	}
	
	return nil
}

func (s *transactionService) confirmationTopUp(c *fiber.Ctx, transaction *entity.Transaction, wallet *entity.Wallet, currency *entity.Currency, tx *sql.Tx) error {
	
	walletBefore := wallet.Balance
	walletAfer := wallet.Balance + uint64(transaction.Amount)
	wallet.Balance = walletAfer
	wallet, err := s.walletRepo.Update(c.Context(), tx, wallet)
	if err != nil {
		return err
	}

	now := time.Now()
	walletMutation := entity.WalletMutation{
		WalletID: wallet.ID,
		MutationType: string(entity.WalletMutationTypeTopUp),
		Amount: transaction.Amount,
		BalanceBefore: walletBefore,
		BalanceAfter: walletAfer,
		CreatedAt: &now,
	}

	if wallet.Balance != walletMutation.BalanceAfter {
		return errors.New("Invalid Balance")
	}
	
	_, err = s.walletMutationRepo.Create(c.Context(), tx, &walletMutation)
	if err != nil {
		return err
	}

	transaction.Status = string(entity.TransactionStatusSuccess)
	transaction, err = s.transactionRepo.Update(c.Context(), tx, transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) confirmationPayment(c *fiber.Ctx, transaction *entity.Transaction, wallet *entity.Wallet, currency *entity.Currency, tx *sql.Tx) error {
	
	if wallet.Balance < uint64(transaction.Amount) {
		return errors.New("Insufficient Balance")
	}

	walletBefore := wallet.Balance
	walletAfer := wallet.Balance - uint64(transaction.Amount)
	wallet.Balance = walletAfer

	
	wallet, err := s.walletRepo.Update(c.Context(), tx, wallet)
	if err != nil {
		return err
	}

	now := time.Now()
	walletMutation := entity.WalletMutation{
		WalletID: wallet.ID,
		MutationType: string(entity.WalletMutationTypePayment),
		Amount: -transaction.Amount,
		BalanceBefore: walletBefore,
		BalanceAfter: walletAfer,
		CreatedAt: &now,
	}

	if wallet.Balance != walletMutation.BalanceAfter {
		return errors.New("Invalid Balance")
	}
	
	_, err = s.walletMutationRepo.Create(c.Context(), tx, &walletMutation)
	if err != nil {
		return err
	}
	
	transaction.Status = string(entity.TransactionStatusSuccess)
	transaction, err = s.transactionRepo.Update(c.Context(), tx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) confirmationTransfer(c *fiber.Ctx, transaction *entity.Transaction, senderWallet *entity.Wallet, receiverWallet *entity.Wallet, tx *sql.Tx) error {
	
	if transaction.TransferedWalletID == nil {
		return errors.New("Invalid Transfered Wallet ID")
	}

	if senderWallet.Currency != receiverWallet.Currency {
		return errors.New("Cannot transfer to different currency")
	}

	if senderWallet.Status != string(entity.WalletStatusActive) {
		return errors.New("Sender Wallet is "+ senderWallet.Status)
	}

	if receiverWallet.Status != string(entity.WalletStatusActive) {
		return errors.New("Receiver Wallet is "+ receiverWallet.Status)
	}

	currencies, err := s.currencyRepo.FindByCodes(c.Context(), []string{senderWallet.Currency, receiverWallet.Currency})
	if err != nil {
		return err
	}

	if currencies == nil || len(*currencies) > 1 {
		return errors.New("Cannot transfer to different currency")
	}

	// sender wallet 
	if senderWallet.Balance < uint64(transaction.Amount) {
		return errors.New("Insufficient Balance")
	}
	senderWalletBalanceBefore := senderWallet.Balance
	senderWalletBalanceAfter := senderWallet.Balance - uint64(transaction.Amount)
	senderWallet.Balance = senderWalletBalanceAfter
	
	_, err = s.walletRepo.Update(c.Context(), tx, senderWallet)
	if err != nil {
		return err
	}

	now := time.Now()
	senderWalletMutation := entity.WalletMutation{
		WalletID: senderWallet.ID,
		MutationType: string(entity.WalletMutationTypeTransferOut),
		Amount: -transaction.Amount,
		BalanceBefore: senderWalletBalanceBefore,
		BalanceAfter: senderWalletBalanceAfter,
		CreatedAt: &now,
	}

	if senderWallet.Balance != senderWalletMutation.BalanceAfter {
		return errors.New("Invalid Balance")
	}
	
	_, err = s.walletMutationRepo.Create(c.Context(), tx, &senderWalletMutation)
	if err != nil {
		return err
	}

	// receiver wallet 
	receiverWalletBalanceBefore := receiverWallet.Balance
	receiverWalletBalanceAfter := receiverWallet.Balance + uint64(transaction.Amount)
	receiverWallet.Balance = receiverWalletBalanceAfter
	_, err = s.walletRepo.Update(c.Context(), tx, receiverWallet)
	if err != nil {
		return err
	}

	receiverWalletMutation := entity.WalletMutation{
		WalletID: receiverWallet.ID,
		MutationType: string(entity.WalletMutationTypeTransferIn),
		Amount: transaction.Amount,
		BalanceBefore: receiverWalletBalanceBefore,
		BalanceAfter: receiverWalletBalanceAfter,
		CreatedAt: &now,
	}

	if receiverWallet.Balance != receiverWalletMutation.BalanceAfter {
		return errors.New("Invalid Balance")
	}
	
	_, err = s.walletMutationRepo.Create(c.Context(), tx, &receiverWalletMutation)
	if err != nil {
		return err
	}

	transaction.Status = string(entity.TransactionStatusSuccess)
	transaction, err = s.transactionRepo.Update(c.Context(), tx, transaction)
	if err != nil {
		return err
	}
	return nil
}
