package pkg

type ApiGameInterface interface {
	Profile() (ProfileResponse, error)
	Order(order OrderCall) (OrderResponse, error)
	ListService(filter FilterListService) (ListServiceResponse, error)
	DetailService(service_code string) (DetailServiceResponse, error)
	Game() ([]GameResponse, error)
}
