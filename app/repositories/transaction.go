package repositories

import (
	"context"
	"database/sql"
	"time"
	"wallet-api/app/entity"
	"wallet-api/app/utilities"

	"github.com/doug-martin/goqu/v9"
)


type transactionRepository struct {
	db *goqu.Database
}

func NewTransactionRepository(db *goqu.Database) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) FindByTrxId(ctx context.Context, trxId string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	_, err := r.db.From(tableTransaction).
	Select(
		"id",
		"trx_id",
		"wallet_id",
		"transfered_wallet_id",
		"transaction_type",
		"amount",
		"status",
		"created_at",
		"updated_at",
	).
	Where(goqu.Ex{"trx_id": trxId}).ScanStruct(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) Update(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error) {
	sql, args, err := r.db.From(tableTransaction).
	Update().
	Set(transaction).
	Where(goqu.Ex{"id": transaction.ID}).
	Prepared(true).
	ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) LockRowByTrxId(ctx context.Context, tx *sql.Tx, trxId string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	sql, args, err := r.db.From(tableTransaction).
		Select(
		"id",
		"trx_id",
		"wallet_id",
		"transfered_wallet_id",
		"transaction_type",
		"amount",
		"status",
		"created_at",
		"updated_at",
	).
	Where(goqu.Ex{"trx_id": trxId}).
	ForUpdate(goqu.Wait).Prepared(true).ToSQL()

	row := tx.QueryRowContext(ctx, sql, args...)
	err = row.Scan(
		&transaction.ID,
		&transaction.TrxID,
		&transaction.WalletID,
		&transaction.TransferedWalletID,
		&transaction.TransactionType,
		&transaction.Amount,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) Create(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) (*entity.Transaction, error) {
	
	trxId, err := r.generateTrxId()
	if err != nil {
		return nil, err
	}
	now := time.Now()


	transaction.TrxID = trxId
	transaction.CreatedAt = &now
	transaction.UpdatedAt = &now
	q, args, err := r.db.From(tableTransaction).Insert().Rows(
		transaction,
	).Prepared(true).
	ToSQL()

	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}


	return &entity.Transaction{
		ID: lastInsertId,
		TrxID: transaction.TrxID,
		WalletID: transaction.WalletID,
		TransferedWalletID: transaction.TransferedWalletID,
		TransactionType: transaction.TransactionType,
		Amount: transaction.Amount,
		CreatedAt: &now,
		UpdatedAt: &now,
	}, nil
}

func (r *transactionRepository) generateTrxId() (string ,error) {
	 utilities.GenerateUUID()
	 trxId := utilities.GenerateUUID()
	
	check := uint64(0)
	_, err := r.db.From(tableTransaction).
	Select("id").
	Where(goqu.Ex{"trx_id": trxId}).ScanVal(&check)
	
	
	if err != nil {
		return "", err
	}
	if check > 0 {
		return r.generateTrxId()
	}

	return trxId, nil
}