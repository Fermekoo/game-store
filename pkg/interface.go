package pkg

type ApiGameInterface interface {
	Profile() (ProfileResponse, error)
	Order(order OrderCall) (OrderResponse, error)
	ListService(filter FilterListService) (ServiceResponse, error)
	Game() ([]GameResponse, error)
}
