package wallet

import (
	"context"
	"errors"
	"log"
	"wallet-api/app/entity"
	"wallet-api/app/presentation"
	"wallet-api/app/utilities"

	"github.com/shopspring/decimal"
)

func (s *walletService) Transfer(c context.Context, req *presentation.TransactionTransferRequest) (*presentation.WalletActionConfirmation, error) {
	
	amountDecimal, err := utilities.StringToDecimal(req.Amount)
	if err != nil {
		return nil, err
	}

	if amountDecimal.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("Invalid Amount")
	}

	wallets, err := s.walletRepo.FindByWalletIds(c, []string{req.FromWalletId, req.ToWalletId})
	if err != nil {
		return nil, err
	}

	if wallets == nil || len(*wallets) < 2 {
		return nil, errors.New("Invalid Wallet")
	}

	senderWallet := entity.Wallet{}
	receiverWallet := entity.Wallet{}
	for _, w := range *wallets {
		if w.WalletID == req.FromWalletId {
			senderWallet = w
		}
		if w.WalletID == req.ToWalletId {
			receiverWallet = w
		}
	}

	if senderWallet.Currency != receiverWallet.Currency {
		return nil, errors.New("Cannot transfer to different currency")
	}

	if senderWallet.Status != string(entity.WalletStatusActive) {
		return nil, errors.New("Sender Wallet is "+ senderWallet.Status)
	}

	if receiverWallet.Status != string(entity.WalletStatusActive) {
		return nil, errors.New("Receiver Wallet is "+ receiverWallet.Status)
	}

	currencies, err := s.currencyRepo.FindByCodes(c, []string{senderWallet.Currency, receiverWallet.Currency})
	if err != nil {
		return nil, err
	}

	if currencies == nil || len(*currencies) > 1 {
		return nil, errors.New("Cannot transfer to different currency")
	}

	currency := (*currencies)[0]
	if currency.Precision == 0 && amountDecimal.Exponent() < 0 {
		return nil, errors.New(currency.Code + " does not support decimal")
	}
	

	var amountMinor int64
	amountMinor = amountDecimal.BigInt().Int64()
	if currency.Precision > 0 {
		amountMinor, err = utilities.ToMinorUnit(amountDecimal, int32(currency.Precision))
		if err != nil {
			return nil, err
		}
	}

	if senderWallet.Balance < uint64(amountMinor) {
		return nil, errors.New("Insufficient Balance")
	}

	tx, err := s.transactionManager.BeginTx(c)
	if err != nil {
		return nil, err
	}

	defer func() {
		commitErr := s.transactionManager.CommitOrRollback(err, tx)
		if commitErr != nil {
			log.Default().Println(commitErr)
		}
	}()

	transaction := entity.Transaction{
		WalletID: senderWallet.ID,
		TransferedWalletID: &receiverWallet.ID,
		TransactionType: string(entity.TransactionTypeTransfer),
		Amount: amountMinor,
		Status: string(entity.TransactionStatusPending),
	}

	trx, err := s.transactionRepo.Create(c, tx, &transaction)
	if err != nil {
		return nil, err
	}
	
	result := presentation.WalletActionConfirmation{
		TrxID: trx.TrxID,
		Status: string(entity.TransactionStatusPending),
	}

	return &result, nil
}
