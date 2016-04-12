package models

import (
	"time"
)

type UserCommunity struct {
	ID            int
	UserId        int
	CommunityId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
