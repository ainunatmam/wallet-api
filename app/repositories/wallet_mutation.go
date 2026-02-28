package repositories

import (
	"context"
	"database/sql"
	"wallet-api/app/entity"

	"github.com/doug-martin/goqu/v9"
)

type walletMutationRepository struct {
	db *goqu.Database
}

func NewWalletMutationRepository(db *goqu.Database) WalletMutationRepository {
	return &walletMutationRepository{db: db}
}

func (r *walletMutationRepository) Create(ctx context.Context, tx *sql.Tx, walletMutation *entity.WalletMutation) (*int64, error) {
	
	q, args, err := r.db.From(tableWalletMutation).Insert().Rows(
		walletMutation,
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

	return &lastInsertId, nil
}