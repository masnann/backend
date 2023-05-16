package migration

import (
	"backend/cms/book/book_domain"
	"backend/cms/photo/photo_domain"
	"backend/cms/user/user_domain"
	"backend/utils/database"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&user_domain.UserModel{}, &book_domain.BookModels{}, &photo_domain.Category{}, &photo_domain.Photo{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")

}
