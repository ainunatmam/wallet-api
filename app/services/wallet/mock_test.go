package wallet

import (
	"context"
	"database/sql"
	"wallet-api/app/entity"

	"github.com/stretchr/testify/mock"
)

// MockWalletRepository
type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Create(ctx context.Context, ownerId int, currency string) (*entity.Wallet, error) {
	args := m.Called(ctx, ownerId, currency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) FindByWalletId(ctx context.Context, walletId string) (*entity.Wallet, error) {
	args := m.Called(ctx, walletId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) FindByWalletIds(ctx context.Context, walletIds []string) (*[]entity.Wallet, error) {
	args := m.Called(ctx, walletIds)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) FindById(ctx context.Context, id uint64) (*entity.Wallet, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) FindByIds(ctx context.Context, ids []uint64) (*[]entity.Wallet, error) {
	args := m.Called(ctx, ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) FindByCurrencyCode(ctx context.Context, currency string) (*entity.Wallet, error) {
	args := m.Called(ctx, currency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) Update(ctx context.Context, tx *sql.Tx, wallet *entity.Wallet) (*entity.Wallet, error) {
	args := m.Called(ctx, tx, wallet)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) LockRowById(ctx context.Context, tx *sql.Tx, Id uint64) (*entity.Wallet, error) {
	args := m.Called(ctx, tx, Id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) LockRowByIds(ctx context.Context, tx *sql.Tx, ids []uint64) (*[]entity.Wallet, error) {
	args := m.Called(ctx, tx, ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Wallet), args.Error(1)
}

func (m *MockWalletRepository) ListByOwnerId(ctx context.Context, ownerId uint64) (*[]entity.Wallet, error) {
	args := m.Called(ctx, ownerId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Wallet), args.Error(1)
}

// MockWalletMutationRepository
type MockWalletMutationRepository struct {
	mock.Mock
}

func (m *MockWalletMutationRepository) Create(ctx context.Context, tx *sql.Tx, walletMutation *entity.WalletMutation) (*int64, error) {
	args := m.Called(ctx, tx, walletMutation)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*int64), args.Error(1)
}

// MockCurrencyRepository
type MockCurrencyRepository struct {
	mock.Mock
}

func (m *MockCurrencyRepository) FindByCode(ctx context.Context, code string) (*entity.Currency, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Currency), args.Error(1)
}

func (m *MockCurrencyRepository) FindByCodes(ctx context.Context, codes []string) (*[]entity.Currency, error) {
	args := m.Called(ctx, codes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Currency), args.Error(1)
}

// MockTransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) LockRowByTrxId(ctx context.Context, tx *sql.Tx, trxId string) (*entity.Transaction, error) {
	args := m.Called(ctx, tx, trxId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Update(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error) {
	args := m.Called(ctx, tx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindByTrxId(ctx context.Context, trxId string) (*entity.Transaction, error) {
	args := m.Called(ctx, trxId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Create(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error) {
	args := m.Called(ctx, tx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

// MockTransactionManager
type MockTransactionManager struct {
	mock.Mock
}

func (m *MockTransactionManager) BeginTx(ctx context.Context) (*sql.Tx, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sql.Tx), args.Error(1)
}

func (m *MockTransactionManager) CommitTx(tx *sql.Tx) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockTransactionManager) RollbackTx(tx *sql.Tx) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockTransactionManager) CommitOrRollback(err error, tx *sql.Tx) error {
	args := m.Called(err, tx)
	return args.Error(0)
}
