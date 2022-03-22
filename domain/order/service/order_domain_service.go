package service

import "github.com/garfcat/go-ddd/domain/order/repository"

type OrderDomainService interface {
}

type orderDomainService struct {
	r repository.Repository
}

func NewOrderDomainService(r repository.Repository) OrderDomainService {
	return &orderDomainService{r: r}
}