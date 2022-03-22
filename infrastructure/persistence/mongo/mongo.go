// Package mongo is a mongo implementation of the Customer Repository
package mongo

import (
	"context"
	"time"

	"github.com/garfcat/go-ddd/domain/customer/repository/po"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
	// customer is used to store customers
	customer *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customeraggregate.
// Mongo uses bson so we add tags for that
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromCustomer(c po.CustomerPo) mongoCustomer {
	return mongoCustomer{
		ID:   c.ID,
		Name: c.Name,
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() po.CustomerPo {
	c := po.CustomerPo{}

	c.ID = m.ID
	c.Name = m.Name

	return c
}

// New Create a new mongodb repository
func New(ctx context.Context, connectionString string) (*Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &Repository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *Repository) Get(id uuid.UUID) (po.CustomerPo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return po.CustomerPo{}, err
	}
	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (mr *Repository) Add(c po.CustomerPo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *Repository) Update(c po.CustomerPo) error {
	panic("to implement")
}
