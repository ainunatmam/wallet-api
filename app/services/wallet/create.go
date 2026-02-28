package wallet

import (
	"context"
	"errors"
	"wallet-api/app/entity"
	"wallet-api/app/presentation"
)



func (s *walletService) Create(c context.Context, req *presentation.WalletCreateRequest) (*entity.Wallet, error) {

	checkWallet, err := s.walletRepo.FindByCurrencyCode(c, req.Currency)
	if err != nil {
		return nil, err
	}

	if checkWallet != nil && checkWallet.OwnerID == req.OwnerId {
		return nil, errors.New("Wallet " + req.Currency + " already exists")
	}

	wallet, err := s.walletRepo.Create(c, req.OwnerId, req.Currency)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}