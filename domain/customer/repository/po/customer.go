package po

import (
	"time"

	"github.com/google/uuid"
)

type CustomerPo struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the person
	Name string
	// Age is the age of the person
	Age int
	// CreateTime is the time of the person create
	CreateTime time.Time
	UpdateTIme time.Time
}
