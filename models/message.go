package models

import (
	"time"
)

type Message struct {
	ID        int
	UserId    int
	CommunityId int `binding:"required,CommunityId"`
	Name      string `binding:"required,name"`
	Body      string `binding:"required,body"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
