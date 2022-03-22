package service

import (
	"github.com/garfcat/go-ddd/domain/product/entity"
	"github.com/garfcat/go-ddd/domain/product/repository"
	"github.com/google/uuid"
)

type ProductDomainService interface {
	GetAll() ([]entity.Product, error)
	GetByID(uuid.UUID) (entity.Product, error)
	Add(entity.Product) error
	Update(entity.Product) error
	Delete(uuid.UUID) error
}

type productDomainService struct {
	r repository.ProductRepository
}

func NewProductDomainService(r repository.ProductRepository) ProductDomainService {
	return &productDomainService{r: r}
}

func (cds *productDomainService) GetAll() ([]entity.Product, error) {
	return nil, nil
}
func (cds *productDomainService) GetByID(uuid.UUID) (entity.Product, error) {
	return entity.Product{}, nil
}
func (cds *productDomainService) Add(entity.Product) error {
	return nil
}
func (cds *productDomainService) Update(entity.Product) error {
	return nil

}
func (cds *productDomainService) Delete(uuid.UUID) error {
	return nil
}
