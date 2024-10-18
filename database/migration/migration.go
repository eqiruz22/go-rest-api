package migration

import (
	"fiber/backend/database"
	"fiber/backend/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Auth{}, &entity.Role{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("migration success")
}