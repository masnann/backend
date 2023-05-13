package migration

import (
	"backend/cms/book/book_domain"
	"backend/cms/user/user_domain"
	"backend/utils/database"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&user_domain.UserModel{}, &book_domain.BookModels{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")

}
