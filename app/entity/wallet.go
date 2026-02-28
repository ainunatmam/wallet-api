package entity

import (
	"time"
)

type WalletStatus string

const (
	WalletStatusActive    WalletStatus = "ACTIVE"
	WalletStatusSuspended WalletStatus = "SUSPENDED"
)


type Wallet struct {
	ID        uint64    `json:"id" db:"id"`
	WalletID  string    `json:"wallet_id" db:"wallet_id"`
	OwnerID   int       `json:"owner_id" db:"owner_id"`
	Currency  string    `json:"currency" db:"currency"`
	Balance   uint64    `json:"balance" db:"balance"`
	Status    string    `json:"status" db:"status"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

func DefaultBalanceWallet() uint64 {
	return 0
}
