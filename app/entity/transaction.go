package entity

import "time"

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeTopUp TransactionType = "TOP_UP"
	TransactionTypePayment TransactionType = "PAYMENT"
	TransactionTypeTransfer TransactionType = "TRANSFER"
	TransactionTypeTransferIn TransactionType = "TRANSFER_IN"
	TransactionTypeTransferOut TransactionType = "TRANSFER_OUT"
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed TransactionStatus = "FAILED"
)

type Transaction struct {
	ID              int64     `json:"id" db:"id"`
	TrxID           string    `json:"trx_id" db:"trx_id"`
	WalletID        uint64    `json:"wallet_id" db:"wallet_id"`
	TransferedWalletID *uint64    `json:"transfered_wallet_id" db:"transfered_wallet_id"`
	TransactionType string    `json:"transaction_type" db:"transaction_type"`
	Amount          int64     `json:"amount" db:"amount"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
}