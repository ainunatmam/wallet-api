package transaction

import (
	"wallet-api/app/libraries"
	"wallet-api/app/presentation"

	"wallet-api/app/repositories"

	"github.com/gofiber/fiber/v2"
)

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	walletRepo repositories.WalletRepository
	walletMutationRepo repositories.WalletMutationRepository
	currencyRepo repositories.CurrencyRepository
	transactionManager libraries.TransactionManager
}

type TransactionService interface {
	Confirmation(c *fiber.Ctx, req *presentation.TransactionConfirmationRequest) (err error)
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepository,
	walletRepo repositories.WalletRepository,
	walletMutationRepo repositories.WalletMutationRepository,
	currencyRepo repositories.CurrencyRepository,
	transactionManager libraries.TransactionManager,
) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		walletRepo: walletRepo,
		walletMutationRepo: walletMutationRepo,
		currencyRepo: currencyRepo,
		transactionManager: transactionManager,
	}
}