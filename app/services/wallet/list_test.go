package wallet

import (
	"context"
	"testing"
	"time"
	"wallet-api/app/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mockWalletRepo := new(MockWalletRepository)
	mockCurrencyRepo := new(MockCurrencyRepository)

	walletService := NewWalletService(
		mockWalletRepo,
		nil,
		mockCurrencyRepo,
		nil,
		nil,
	)

	ownerId := uint64(1)
	wallets := []entity.Wallet{
		{
			WalletID: "W-123",
			Currency: "IDR",
			Balance:  100000,
			Status:   "ACTIVE",
		},
	}

	currencies := []entity.Currency{
		{
			Code:      "IDR",
			Precision: 0,
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockWalletRepo.On("ListByOwnerId", mock.Anything, ownerId).Return(&wallets, nil)
		mockCurrencyRepo.On("FindByCodes", mock.Anything, []string{"IDR"}).Return(&currencies, nil)

		res, err := walletService.List(ctx, ownerId)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 1, len(*res))
		assert.Equal(t, "W-123", (*res)[0].WalletID)
		assert.Equal(t, "100000", (*res)[0].Balance)
	})

	t.Run("Wallet Repo Error", func(t *testing.T) {
		mockWalletRepo.On("ListByOwnerId", mock.Anything, ownerId).Return(nil, assert.AnError).Once()

		res, err := walletService.List(ctx, ownerId)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestDetail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mockWalletRepo := new(MockWalletRepository)
	mockCurrencyRepo := new(MockCurrencyRepository)

	walletService := NewWalletService(
		mockWalletRepo,
		nil,
		mockCurrencyRepo,
		nil,
		nil,
	)

	walletId := "W-123"
	wallet := &entity.Wallet{
		WalletID: walletId,
		Currency: "IDR",
		Balance:  100000,
		Status:   "ACTIVE",
	}

	currency := &entity.Currency{
		Code:      "IDR",
		Precision: 0,
	}

	t.Run("Success", func(t *testing.T) {
		mockWalletRepo.On("FindByWalletId", mock.Anything, walletId).Return(wallet, nil)
		mockCurrencyRepo.On("FindByCode", mock.Anything, "IDR").Return(currency, nil)

		res, err := walletService.Detail(ctx, walletId)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "W-123", res.WalletID)
		assert.Equal(t, "100000", res.Balance)
	})

	t.Run("Wallet Not Found", func(t *testing.T) {
		mockWalletRepo.On("FindByWalletId", mock.Anything, walletId).Return(nil, assert.AnError).Once()

		res, err := walletService.Detail(ctx, walletId)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}