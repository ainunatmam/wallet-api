package repositories

import (
	"context"
	"database/sql"
	"wallet-api/app/entity"
)


const (
	tableOwner = "owners"
	tableWallet = "wallets"
	tableCurrency = "currencies"
	tableWalletMutation = "wallets_mutations"
	tableTransaction = "transactions"
)

type OwnerRepository interface {
	Create(ctx context.Context, name string) error
}

type WalletRepository interface {
	Create(ctx context.Context, ownerId int, currency string) (*entity.Wallet, error)
	FindByWalletId(ctx context.Context,walletId string) (*entity.Wallet, error)
	FindByWalletIds(ctx context.Context, walletIds []string) (*[]entity.Wallet, error)
	FindById(ctx context.Context,id uint64) (*entity.Wallet, error)
	FindByIds(ctx context.Context, ids []uint64) (*[]entity.Wallet, error)
	FindByCurrencyCode(ctx context.Context, currency string) (*entity.Wallet, error)
	Update(ctx context.Context, tx *sql.Tx, wallet *entity.Wallet) (*entity.Wallet, error)
	LockRowById(ctx context.Context, tx *sql.Tx, Id uint64) (*entity.Wallet, error)
	LockRowByIds(ctx context.Context, tx *sql.Tx, ids []uint64) (*[]entity.Wallet, error)
	ListByOwnerId(ctx context.Context, ownerId uint64) (*[]entity.Wallet, error)
}

type WalletMutationRepository interface {
	Create(ctx context.Context, tx *sql.Tx, walletMutation *entity.WalletMutation) (*int64, error)
}

type CurrencyRepository interface {
	FindByCode(ctx context.Context, code string) (*entity.Currency, error)
	FindByCodes(ctx context.Context,codes []string) (*[]entity.Currency, error)
}

type TransactionRepository interface {
	LockRowByTrxId(ctx context.Context, tx *sql.Tx, trxId string) (*entity.Transaction, error)
	Update(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error)
	FindByTrxId(ctx context.Context, trxId string) (*entity.Transaction, error)
	Create(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error)
}