package models

import (
	"time"
)

type Message struct {
	ID        int
	UserId    int
	CommunityId int
	Name      string `sql:"size:255"`
	Body      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageForm struct {
	Message
}
