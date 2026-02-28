package wallet

import (
	"context"
	"wallet-api/app/entity"
	"wallet-api/app/presentation"
	"wallet-api/app/utilities"
)

func (s *walletService) List(c context.Context, ownerId uint64) (*[]presentation.DetailWalletResponse, error) {
	
	wallets, err := s.walletRepo.ListByOwnerId(c, ownerId)
	if err != nil {
		return nil, err
	}

	var currencyCodes []string
	for _, w := range *wallets {
		currencyCodes = append(currencyCodes, w.Currency)
	}

	currencies, err := s.currencyRepo.FindByCodes(c, currencyCodes)
	if err != nil {
		return nil, err
	}

	response, err := s.mapWalletToResponse(wallets, currencies)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *walletService) mapWalletToResponse(wallets *[]entity.Wallet, currencies *[]entity.Currency) (*[]presentation.DetailWalletResponse, error) {
	
	currencyMap := make(map[string]entity.Currency)
	for _, currency := range *currencies {
		currencyMap[currency.Code] = currency
	}

	var responses []presentation.DetailWalletResponse

	for _, wallet := range *wallets {
		responses = append(responses, presentation.DetailWalletResponse{
			WalletID: wallet.WalletID,
			Currency: wallet.Currency,
			Balance:  utilities.ToMajorString(wallet.Balance, int32(currencyMap[wallet.Currency].Precision)),
			Status:   wallet.Status,
		})
	}

	return &responses, nil
}