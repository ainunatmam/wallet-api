package repositories

import (
	"context"
	"wallet-api/app/entity"

	"github.com/doug-martin/goqu/v9"
)

type currencyRepository struct {
	db *goqu.Database
}

func NewCurrencyRepository(db *goqu.Database) CurrencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) FindByCode(ctx context.Context, code string) (*entity.Currency, error) {
	var currency entity.Currency
	_, err := r.db.From(tableCurrency).
	Select(
		"id",
		"code",
		"name",
		"precision",
		"created_at",
	).
	Where(goqu.Ex{"code": code}).ScanStruct(&currency)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *currencyRepository) FindByCodes(ctx context.Context, codes []string) (*[]entity.Currency, error) {
	var currencies []entity.Currency
	err := r.db.From(tableCurrency).
	Select(
		"id",
		"code",
		"name",
		"precision",
		"created_at",
	).
	Where(
		goqu.I("code").In(codes),
	).ScanStructs(&currencies)
	if err != nil {
		return nil, err
	}
	return &currencies, nil
}