package pkg

type GeneralResponse struct {
	Result  bool        `json:"result"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type ProfileResponse struct {
	Result  bool        `json:"result"`
	Data    DataProfile `json:"data"`
	Message string      `json:"message"`
}

type DataProfile struct {
	Fullname   string `json:"full_name"`
	Username   string `json:"user_name"`
	Balance    uint64 `json:"balance"`
	Point      uint64 `json:"point"`
	Level      string `json:"level"`
	Registered string `json:"registered"`
}

type OrderCall struct {
	ServiceCode string `json:"service"`
	AccountID   uint64 `json:"account_id"`
	AccountZone string `json:"account_zone"`
}
type OrderCallRequest struct {
	ServiceCode   string `json:"service" binding:"required"`
	AccountID     string `json:"account_id" binding:"required"`
	AccountZone   string `json:"account_zone"`
	PaymentMethod string `json:"payment_method"`
	Phone         string `json:"phone"`
	Name          string `json:"name"`
}

type OrderResponse struct {
	Result  bool       `json:"result"`
	Data    *DataOrder `json:"data"`
	Message string     `json:"message"`
}

type DataOrder struct {
	TRXID       string `json:"trx_id"`
	AccountID   string `json:"data"`
	AccountZone string `json:"zone"`
	Service     string `json:"service"`
	Status      string `json:"status"`
	Note        string `json:"note"`
	Balance     string `json:"balance"`
	Price       string `json:"price"`
}

type ListServiceResponse struct {
	Result  bool              `json:"result"`
	Data    []DataServiceGame `json:"data"`
	Message string            `json:"message"`
}
type DetailServiceResponse struct {
	Result  bool            `json:"result"`
	Data    DataServiceGame `json:"data"`
	Message string          `json:"message"`
}

type DataServiceGame struct {
	Code   string       `json:"code"`
	Game   string       `json:"game"`
	Name   string       `json:"name"`
	Price  ServicePrice `json:"price"`
	Server string       `json:"server"`
	Status string       `json:"status"`
}

type ServicePrice struct {
	Basic   uint `json:"basic"`
	Premium uint `json:"premium"`
	Special uint `json:"special"`
}

type GameResponse struct {
	Name string `json:"name"`
}

type FilterListService struct {
	FilterType  string
	FilterValue string
}

type FilterRequestListService struct {
	Game string `form:"game"`
}
