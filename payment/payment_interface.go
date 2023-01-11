package payment

type Payment interface {
	Charge(payloads *ChargeRequest) (*ChargeResponse, error)
}
