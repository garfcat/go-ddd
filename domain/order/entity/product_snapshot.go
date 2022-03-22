package entity

import "github.com/google/uuid"

type ProductSnapshot struct {
	Id    uuid.UUID
	Name  string
	price float64
}
