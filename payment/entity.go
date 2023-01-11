package payment

type ChargeRequest struct {
	PaymentType string
	Bank        string
	Amount      int64
	Name        string
	Phone       string
	OrderID     string
}

type ChargeResponse struct {
	TransactionID string
	OrderID       string
	PaymentType   string
	PaymentCode   string
	PaymentUrl    string
	QRString      string
}
