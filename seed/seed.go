package main

import (
	"fmt"
	"strconv"
	"github.com/t-fukui/eto_pirka/db"
	"github.com/t-fukui/eto_pirka/models"
)

func main() {
	db := db.InitDB()
	for i := 0; i < 3; i++ {
		user := models.User{Name: fmt.Sprintf("ホゲの子" + strconv.Itoa(i+1))}
		db.Create(&user)
		community := models.Community{
			Name: fmt.Sprintf("ホゲの会" + strconv.Itoa(i+1)),
			Description: fmt.Sprintf("ホゲの技術交流会" + strconv.Itoa(i+1)),
			AdministratorId: user.ID}
		db.Create(&community)
		user_community := models.UserCommunity{UserId: user.ID, CommunityId: community.ID}
		db.Create(&user_community)
		for i := 0; i < 3; i++ {
			message := models.Message{
				Name:        "名無し",
				Body:        "こんにちは、hogehogehoge",
				UserId:      user.ID,
				CommunityId: community.ID}
			db.Create(&message)
		}
	}
}
