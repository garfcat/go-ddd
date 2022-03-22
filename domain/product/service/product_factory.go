package service

import (
	"github.com/garfcat/go-ddd/domain/product/entity"
	"github.com/garfcat/go-ddd/domain/product/repository/po"
)

func CreateProduct(po po.ProductPo) entity.Product {
	return entity.Product{}
}

func CreateProductPo(Product entity.Product) po.ProductPo {
	return po.ProductPo{}
}
