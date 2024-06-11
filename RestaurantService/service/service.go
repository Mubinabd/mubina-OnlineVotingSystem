package service

type Service struct {
	RS  *RestaurantService
}

func NewService(rs *RestaurantService) *Service {
	return &Service{RS: rs}
}