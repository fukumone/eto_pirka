package models

import (
	"time"
)

type User struct {
	ID        int
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
