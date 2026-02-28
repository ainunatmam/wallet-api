package wallet

import (
	"context"
	"wallet-api/app/entity"
	"wallet-api/app/libraries"
	"wallet-api/app/presentation"
	"wallet-api/app/repositories"
)

type walletService struct {
	walletRepo repositories.WalletRepository
	transactionManager libraries.TransactionManager
	walletMutationRepo repositories.WalletMutationRepository
	currencyRepo repositories.CurrencyRepository
	transactionRepo repositories.TransactionRepository
}

type WalletService interface {
	Create(c context.Context, req *presentation.WalletCreateRequest) (*entity.Wallet, error)
	TopUp(c context.Context, walletId string, req *presentation.TransactionTopUpRequest) (*presentation.WalletActionConfirmation, error)
	Payment(c context.Context, walletId string, req *presentation.TransactionPaymentRequest) (*presentation.WalletActionConfirmation, error)
	Transfer(c context.Context, req *presentation.TransactionTransferRequest) (*presentation.WalletActionConfirmation, error)
	Suspend(c context.Context, walletId string) error
	Detail(c context.Context, walletId string) (*presentation.DetailWalletResponse, error)
	List(c context.Context, ownerId uint64) (*[]presentation.DetailWalletResponse, error)
}

func NewWalletService(
	walletRepo repositories.WalletRepository,
	walletMutationRepo repositories.WalletMutationRepository,
	currencyRepo repositories.CurrencyRepository,
	transactionRepo repositories.TransactionRepository,
	transactionManager libraries.TransactionManager,
	) WalletService {
	return &walletService{
		walletRepo: walletRepo, 
		walletMutationRepo: walletMutationRepo,
		transactionManager: transactionManager,
		currencyRepo: currencyRepo,
		transactionRepo: transactionRepo,
	}
}
