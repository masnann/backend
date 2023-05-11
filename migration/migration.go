package migration

import (
	"backend/cms/user/user_domain"
	"backend/utils/database"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&user_domain.UserModel{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")

}
