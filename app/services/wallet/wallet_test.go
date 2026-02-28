package wallet

import (
	"context"
	"testing"
	"time"
	"wallet-api/app/entity"
	"wallet-api/app/presentation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mockWalletRepo := new(MockWalletRepository)

	walletService := NewWalletService(
		mockWalletRepo,
		nil,
		nil,
		nil,
		nil,
	)

	req := &presentation.WalletCreateRequest{
		OwnerId:  1,
		Currency: "IDR",
	}

	t.Run("Success", func(t *testing.T) {
		mockWalletRepo.On("FindByCurrencyCode", mock.Anything, "IDR").Return(nil, nil).Once()
		mockWalletRepo.On("Create", mock.Anything, req.OwnerId, req.Currency).Return(&entity.Wallet{ID: 1}, nil).Once()

		res, err := walletService.Create(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, uint64(1), res.ID)
	})

	t.Run("Wallet Already Exists", func(t *testing.T) {
		mockWalletRepo.On("FindByCurrencyCode", mock.Anything, "IDR").Return(&entity.Wallet{}, nil).Once()

		res, err := walletService.Create(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already exists")
		assert.Nil(t, res)
	})

	t.Run("Repo Error", func(t *testing.T) {
		mockWalletRepo.On("FindByCurrencyCode", mock.Anything, "IDR").Return(nil, assert.AnError).Once()

		res, err := walletService.Create(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
