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

func (s *walletService) Payment(c context.Context, walletId string, req *presentation.TransactionPaymentRequest) (*presentation.WalletActionConfirmation, error) {
	
	amountDecimal, err := utilities.StringToDecimal(req.Amount)
	if err != nil {
		return nil, err
	}

	if amountDecimal.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("Invalid Amount")
	}

	wallet, err := s.walletRepo.FindByWalletId(c,walletId)
	if err != nil {
		return nil, err
	}

	if wallet.Status != string(entity.WalletStatusActive) {
		return nil, errors.New("Wallet is "+ wallet.Status)
	}

	currency, err := s.currencyRepo.FindByCode(c, wallet.Currency)
	if err != nil {
		return nil, err
	}

	if currency == nil {
		return nil, errors.New("Invalid Currency")
	}

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

	tx, err := s.transactionManager.BeginTx(c)
	if err != nil {
		return nil, err
	}

	if wallet.Balance < 0 {
		return nil, errors.New("Insufficient Balance")
	}

	defer func() {
		commitErr := s.transactionManager.CommitOrRollback(err, tx)
		if commitErr != nil {
			log.Default().Println(commitErr)
		}
	}()

	transaction := entity.Transaction{
		WalletID: wallet.ID,
		Amount: amountMinor,
		TransactionType: string(entity.TransactionTypePayment),
		Status: string(entity.TransactionStatusPending),
	}

	trx, err := s.transactionRepo.Create(c, tx, &transaction)
	if err != nil {
		return nil, err
	}

	result := presentation.WalletActionConfirmation{
		TrxID: trx.TrxID,
		Status: trx.Status,
	}

	return &result, nil
}
