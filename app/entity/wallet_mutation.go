package entity

import "time"

const (
	WalletMutationTypeTopUp    = "TOP_UP"
	WalletMutationTypePayment  = "PAYMENT"
	WalletMutationTypeTransferIn = "TRANSFER_IN"
	WalletMutationTypeTransferOut = "TRANSFER_OUT"
)

type WalletMutation struct {
	ID            uint64     `json:"id" db:"id"`
	WalletID      uint64     `json:"wallet_id" db:"wallet_id"`
	MutationType  string     `json:"mutation_type" db:"mutation_type"`
	Amount        int64     `json:"amount" db:"amount"`
	BalanceBefore uint64     `json:"balance_before" db:"balance_before"`
	BalanceAfter  uint64     `json:"balance_after" db:"balance_after"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
}