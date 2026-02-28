package wallet

import (
	"context"
	"log"
	"wallet-api/app/entity"
)

func (s *walletService) Suspend(c context.Context, walletId string) error {
	
	wallet, err := s.walletRepo.FindByWalletId(c, walletId)
	if err != nil {
		return err
	}

	tx, err := s.transactionManager.BeginTx(c)
	if err != nil {
		return err
	}
	defer func() {
		commitErr := s.transactionManager.CommitOrRollback(err, tx)
		if commitErr != nil {
			log.Default().Println(commitErr)
		}
	}()

	wallet.Status = string(entity.WalletStatusSuspended)
	wallet, err = s.walletRepo.Update(c, tx, wallet)
	if err != nil {
		return err
	}

	return nil
}
