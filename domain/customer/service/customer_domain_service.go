package service

import (
	"github.com/garfcat/go-ddd/domain/customer/entity"
	"github.com/garfcat/go-ddd/domain/customer/repository"
	"github.com/google/uuid"
)

type CustomerDomainService interface {
	GetCustomer(id uuid.UUID) (entity.Customer, error)
	CreateCustomer(entity.Customer) error
	UpdateCustomer(entity.Customer) error
}

type customerDomainService struct {
	r repository.Repository
}

func NewCustomerDomainService(r repository.Repository) CustomerDomainService {
	return &customerDomainService{r: r}
}

func (cds *customerDomainService) GetCustomer(id uuid.UUID) (entity.Customer, error) {
	customerPo, err := cds.r.Get(id)
	if err != nil {
		return entity.Customer{}, err
	}
	// po to do
	return CreateCustomer(customerPo), nil
}

func (cds *customerDomainService) CreateCustomer(customer entity.Customer) error {
	customerPo := CreateCustomerPo(customer)
	return cds.r.Add(customerPo)
}

func (cds *customerDomainService) UpdateCustomer(customer entity.Customer) error {
	customerPo := CreateCustomerPo(customer)
	return cds.r.Update(customerPo)
}
