package models

import (
	"time"
)

type Message struct {
	ID        int
	UserId    string
	CommunityId int
	Name      string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ValidMessage(c Message) bool {
	if c.UserId == "" {
		return false
	} else if c.CommunityId == 0 {
		return false
	} else if c.Name == "" {
		return false
	} else if c.Body == "" {
		return false
	}
	return true
}
