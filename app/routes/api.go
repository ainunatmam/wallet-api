package routes

import (
	"wallet-api/app/handlers/owner"
	"wallet-api/app/handlers/transaction"
	"wallet-api/app/handlers/wallet"
	"wallet-api/app/repositories"
	ownerService "wallet-api/app/services/owner"
	transactionService "wallet-api/app/services/transaction"
	walletService "wallet-api/app/services/wallet"

	"github.com/gofiber/fiber/v2"
)


func (r *router) api() fiber.Router {
	root := r.app.Group("/")
	api := root.Group("/api")

	// init repo 
	ownerRepo := repositories.NewOwnerRepository(r.goquLibrary.DB())
	walletRepo := repositories.NewWalletRepository(r.goquLibrary.DB())
	walletMutationRepo := repositories.NewWalletMutationRepository(r.goquLibrary.DB())
	currencyRepo := repositories.NewCurrencyRepository(r.goquLibrary.DB())
	transactionRepo := repositories.NewTransactionRepository(r.goquLibrary.DB())
	
	// init service
	ownerService := ownerService.NewOwnerService(ownerRepo)
	walletService := walletService.NewWalletService(walletRepo, walletMutationRepo, currencyRepo, transactionRepo, r.transactionManager)
	transactionService := transactionService.NewTransactionService(transactionRepo, walletRepo, walletMutationRepo, currencyRepo, r.transactionManager)
	
	// init handler
	ownerHandler := owner.NewOwnerHandler(ownerService)
	walletHandler := wallet.NewWalletHandler(walletService)	
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	// Owner
	api.Post("/owner", ownerHandler.Create)	
	api.Get("/owner/:ownerId/wallets", walletHandler.List)

	// Wallet
	api.Post("/wallet", walletHandler.Create)
	api.Post("/wallet/:walletId/top-up", walletHandler.TopUp)
	api.Post("/wallet/:walletId/payment", walletHandler.Payment)
	api.Post("/wallet/transfer", walletHandler.Transfer)
	api.Post("/wallet/:walletId/suspend", walletHandler.Suspend)
	api.Get("/wallet/:walletId", walletHandler.Detail)

	// Transaction
	api.Post("/transaction/confirmation", transactionHandler.Confirmation)

	return root
}
