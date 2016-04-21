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

type MessageForm struct {
	Message
	Token string
	Errors []string
}

func ValidMessage(c *MessageForm, token string) bool {
	if c.Message.UserId == "" {
		c.Errors = append(c.Errors, "ユーザーIDを入力してください")
	}
	if c.Message.CommunityId == 0 {
		c.Errors = append(c.Errors, "コミュニティIDを入力してください")
	}
	if c.Message.Name == "" {
		c.Errors = append(c.Errors, "名前を入力してください")
	}
	if c.Message.Body == "" {
		c.Errors = append(c.Errors, "本文を入力してください")
	}
	if c.Token != token || c.Token == "" {
		c.Errors = append(c.Errors, "トークンが不正です")
	}

	if len(c.Errors) > 0 {
		return false
	}
	return true
}
