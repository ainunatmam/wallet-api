package wallet

import (
	"context"
	"wallet-api/app/presentation"
	"wallet-api/app/utilities"
)

func (s *walletService) Detail(c context.Context, walletId string) (*presentation.DetailWalletResponse, error) {
	
	wallet, err := s.walletRepo.FindByWalletId(c, walletId)
	if err != nil {
		return nil, err
	}

	currency, err := s.currencyRepo.FindByCode(c, wallet.Currency)
	if err != nil {
		return nil, err
	}

	return &presentation.DetailWalletResponse{
		WalletID: wallet.WalletID,
		Currency: wallet.Currency,
		Balance:  utilities.ToMajorString(wallet.Balance, int32(currency.Precision)),
		Status:   wallet.Status,
	}, nil
}
