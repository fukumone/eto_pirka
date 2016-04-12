package models

import (
	"time"
)

type Community struct {
	ID               int
	AdministratorId  int
	Name             string `sql:"size:255"`
	Description      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
