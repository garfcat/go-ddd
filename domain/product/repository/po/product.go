package po

import (
	"time"

	"github.com/google/uuid"
)

type ProductPo struct {
	// item is the root entity which is an item
	ID          uuid.UUID
	Name        string
	Description string
	price       float64
	// Quantity is the number of products in stock
	quantity   int
	CreateTime time.Time
	UpdateTIme time.Time
}
