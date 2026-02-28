package repositories

import (
	"context"
	"database/sql"
	"time"
	"wallet-api/app/entity"
	"wallet-api/app/utilities"

	"github.com/doug-martin/goqu/v9"
)

type walletRepository struct {
	db *goqu.Database
}

func NewWalletRepository(db *goqu.Database) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) LockRowById(ctx context.Context, tx *sql.Tx, Id uint64) (*entity.Wallet, error) {
	var wallet entity.Wallet
	sql, args, err := r.db.From(tableWallet).
		Select(
			"id",
			"wallet_id",
			"owner_id",
			"currency",
			"balance",
			"status",
			"created_at",
			"updated_at",
			"deleted_at",
		).
	Where(goqu.Ex{"id": Id}).
	ForUpdate(goqu.Wait).Prepared(true).ToSQL()

	row := tx.QueryRowContext(ctx, sql, args...)

	err = row.Scan(
		&wallet.ID,
		&wallet.WalletID,
		&wallet.OwnerID,
		&wallet.Currency,
		&wallet.Balance,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
		&wallet.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) LockRowByIds(ctx context.Context, tx *sql.Tx, ids []uint64) (*[]entity.Wallet, error) {
	var wallets []entity.Wallet
	sql, args, err := r.db.From(tableWallet).
		Select(
			"id",
			"wallet_id",
			"owner_id",
			"currency",
			"balance",
			"status",
			"created_at",
			"updated_at",
			"deleted_at",
		).
	Where(goqu.I("id").In(ids)).
	ForUpdate(goqu.Wait).Prepared(true).ToSQL()

	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
			var wallet entity.Wallet
			err = rows.Scan(
			&wallet.ID,
			&wallet.WalletID,
			&wallet.OwnerID,
			&wallet.Currency,
			&wallet.Balance,
			&wallet.Status,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeletedAt,
		)

		if err != nil {
			return nil, err
		}
		
		wallets = append(wallets, wallet)
	}

	return &wallets, nil
}

func (r *walletRepository) FindByWalletId(ctx context.Context, walletId string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	_, err := r.db.From(tableWallet).
		Select(
			"id",
			"wallet_id",
			"owner_id",
			"currency",
			"balance",
			"status",
			"created_at",
			"updated_at",
			"deleted_at",
		).
	Where(goqu.Ex{"wallet_id": walletId}).ScanStruct(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindById(ctx context.Context,id uint64) (*entity.Wallet, error) {
	var wallet entity.Wallet
	_, err := r.db.From(tableWallet).
	Select(
		"id",
		"wallet_id",
		"owner_id",
		"currency",
		"balance",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	).
	Where(goqu.Ex{"id": id}).ScanStruct(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindByCurrencyCode(ctx context.Context,currency string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	_, err := r.db.From(tableWallet).
	Select(
		"id",
		"wallet_id",
		"owner_id",
		"currency",
		"balance",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	).
	Where(goqu.Ex{"currency": currency}).ScanStruct(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindByIds(ctx context.Context, ids []uint64) (*[]entity.Wallet, error) {
	var wallets []entity.Wallet
	err := r.db.From(tableWallet).
	Select(
		"id",
		"wallet_id",
		"owner_id",
		"currency",
		"balance",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	).
	Where(
		goqu.I("id").In(ids),
	).ScanStructs(&wallets)
	if err != nil {
		return nil, err
	}
	return &wallets, nil
}

func (r *walletRepository) FindByWalletIds(ctx context.Context,walletIds []string) (*[]entity.Wallet, error) {
	var wallets []entity.Wallet
	err := r.db.From(tableWallet).
	Select(
		"id",
		"wallet_id",
		"owner_id",
		"currency",
		"balance",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	).
	Where(
		goqu.I("wallet_id").In(walletIds),
	).ScanStructs(&wallets)
	if err != nil {
		return nil, err
	}
	return &wallets, nil
}

func (r *walletRepository) Update(ctx context.Context, tx *sql.Tx, wallet *entity.Wallet) (*entity.Wallet, error) {
	sql, args, err := r.db.From(tableWallet).
	Update().
	Set(wallet).
	Where(goqu.Ex{"wallet_id": wallet.WalletID}).
	Prepared(true).
	ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *walletRepository) ListByOwnerId(ctx context.Context, ownerId uint64) (*[]entity.Wallet, error) {
	var wallets []entity.Wallet
	err := r.db.From(tableWallet).
	Select(
		"id",
		"wallet_id",
		"owner_id",
		"currency",
		"balance",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	).
	Where(
		goqu.Ex{"owner_id": ownerId},
	).ScanStructs(&wallets)
	if err != nil {
		return nil, err
	}
	return &wallets, nil
}

func (r *walletRepository) Create(ctx context.Context,ownerId int, currency string) (*entity.Wallet, error) {
	
	walletId, err := r.generateWalletId()
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
	insert := entity.Wallet{
			OwnerID: ownerId,
			Currency: currency,
			Balance: entity.DefaultBalanceWallet(),
			Status: string(entity.WalletStatusActive),
			WalletID: walletId,
			CreatedAt: &now,
			UpdatedAt: &now,
		}
	
	q, err := r.db.From(tableWallet).Insert().Rows(
		insert,
	).Executor().Exec()

	if err != nil {
		return nil, err
	}

	lastInsertId, err := q.LastInsertId()
	if err != nil {
		return nil, err
	}

	insert.ID = uint64(lastInsertId)
	return &insert, nil
}

func (r *walletRepository) generateWalletId() (string, error) {
	walletId :=  utilities.GenerateUUID()

	// check if wallet id exist
	check := uint64(0)
	_, err := r.db.From(tableWallet).
	Select("id").
	Where(goqu.Ex{"wallet_id": walletId}).ScanVal(&check)
	
	
	if err != nil {
		return "", err
	}
	if check > 0 {
		return r.generateWalletId()
	}
	return walletId, nil
}