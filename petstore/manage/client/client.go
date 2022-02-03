// This package exposes functions to manage
// pets for petstore loneliness 2000.

package client

import (
	"time"

	"github.com/gofrs/uuid"
)

type Client struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	Address   string    `db:"address"`
	Gender    string    `db:"gender"`
	Age       int64     `db:"age"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Clients []Client
