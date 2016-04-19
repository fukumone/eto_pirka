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

type CommunityForm struct {
	Community
	Token string
	Errors []string
}

func ValidCommunity(c *CommunityForm) bool {
	if c.Community.Name == "" {
		c.Errors = append(c.Errors, "コミュニティ名を入力してください")
	}
	if c.Community.Description == "" {
		c.Errors = append(c.Errors, "説明を入力してください")
	}
	if c.Community.AdministratorId == "" {
		c.Errors = append(c.Errors, "管理者IDを入力してください")
	}
	if len(c.Errors) > 0 {
		return false
	}
	return true
}

func CommnunityValidAdmin(c Community, AdministratorId string) bool {
	return c.AdministratorId == AdministratorId
}
