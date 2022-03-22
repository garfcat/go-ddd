// Package main runs the tavern and performs an Order
package main

import (
	"context"
	"log"
	"os"

	"github.com/garfcat/go-ddd/application/service/customer"
	"github.com/garfcat/go-ddd/application/service/order"
	"github.com/garfcat/go-ddd/application/service/product"
	"github.com/garfcat/go-ddd/domain/customer/service"
	orderDomain "github.com/garfcat/go-ddd/domain/order/service"
	productDomain "github.com/garfcat/go-ddd/domain/product/service"
	"github.com/garfcat/go-ddd/infrastructure/persistence/mongo"
	"github.com/garfcat/go-ddd/interface/customer_interface"
	"github.com/garfcat/go-ddd/interface/order_interface"
	"github.com/garfcat/go-ddd/interface/product_interface"
	"github.com/gin-gonic/gin"
)

func main() {
	url := "mongodb://test:test@127.0.0.1:27017"
	// 仓储初始化
	cRepository, err := mongo.New(context.Background(), url)
	if err != nil {
		panic(err)
	}

	// 领域服务初始化
	customerDomainService := service.NewCustomerDomainService(cRepository)
	orderDomainService := orderDomain.NewOrderDomainService(cRepository)
	productDomainService := productDomain.NewProductDomainService(nil)

	// 应用服务
	// Create Order Service
	cs, err := customer.NewCustomerService(
		customer.WithCustomerDomainService(customerDomainService),
	)
	if err != nil {
		panic(err)
	}
	pt, err := product.NewProductService(
		product.WithProductDomainService(productDomainService),
	)
	if err != nil {
		panic(err)
	}

	or, err := order.NewOrderService(
		order.WithCustomerDomainService(customerDomainService),
		order.WithOrderDomainService(orderDomainService),
		order.WithProductDomainService(productDomainService),
	)
	if err != nil {
		panic(err)
	}

	// 接口层
	csi, err := customer_interface.NewCustomerInterface(
		customer_interface.WithCustomerApplicationService(cs),
	)
	if err != nil {
		panic(err)
	}
	oi, err := order_interface.NewOrderInterface(
		order_interface.WithOrderApplicationService(or),
	)
	if err != nil {
		panic(err)
	}
	pi, err := product_interface.NewProductInterface(
		product_interface.WithProductApplicationService(pt),
	)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	//customer routes
	r.POST("/users", csi.CreateCustomer)
	r.POST("/product", pi.CreateProduct)
	r.POST("/order", oi.CreateOrder)

	//Starting the application
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" //localhost
	}
	log.Fatal(r.Run(":" + port))
}
