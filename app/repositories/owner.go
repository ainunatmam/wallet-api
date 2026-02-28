package repositories

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type ownerRepository struct {
	db *goqu.Database
}

func NewOwnerRepository(db *goqu.Database) OwnerRepository {
	return &ownerRepository{db: db}
}

func (r *ownerRepository) Create(ctx context.Context, name string) error {
	_, err := r.db.From(tableOwner).Insert().Rows(goqu.Record{
		"name": name,
	}).Executor().Exec()

	if err != nil {
		return err
	}
	return nil
}
