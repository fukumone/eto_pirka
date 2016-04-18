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
