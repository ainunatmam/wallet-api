package presentation

type ResponseBase struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (ResponseBase) Failed(status int, message string) ResponseBase {
	return ResponseBase{
		Message: message,
	}
}

func (ResponseBase) Success(message string, data interface{}) ResponseBase {
	return ResponseBase{
		Message: message,
		Data:    data,
	}
}

type DetailWalletResponse struct {
	WalletID string `json:"wallet_id"`
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
	Status   string `json:"status"`
}

type WalletActionConfirmation struct {
	TrxID  string `json:"trx_id"`
	Status string `json:"status"`
}