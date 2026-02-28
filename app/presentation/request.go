package presentation

type OwnerCreateRequest struct {
	Name string `validate:"required" json:"name"`
}

type WalletCreateRequest struct {
	OwnerId  int    `validate:"required" json:"owner_id"`
	Currency string `validate:"required,iso4217" json:"currency"`
}

type TransactionTopUpRequest struct {
	Amount string `validate:"required" json:"amount"`
}

type TransactionPaymentRequest struct {
	Amount string `validate:"required" json:"amount"`
}

type TransactionTransferRequest struct {
	FromWalletId string `validate:"required" json:"from_wallet_id"`
	ToWalletId   string `validate:"required" json:"to_wallet_id"`
	Amount       string `validate:"required" json:"amount"`
}

type TransactionConfirmationRequest struct {
	TrxID string `validate:"required" json:"trx_id"`
}