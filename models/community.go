package models

import (
	"time"
)

type Community struct {
	ID               int
	AdministratorId  string
	Name             string
	Description      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ValidCommunity(c Community) bool {
	if c.Name == "" {
		return false
	} else if c.AdministratorId == "" {
		return false
	} else if c.Description == "" {
		return false
	}
	return true
}

func CommnunityValidAdmin(c Community, AdministratorId string) bool {
	return c.AdministratorId == AdministratorId
}
