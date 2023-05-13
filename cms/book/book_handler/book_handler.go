package book_handler

import (
	"backend/cms/book/book_domain"
	"backend/utils/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(book_domain.BookRequestCreate)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal Input Data",
			"error":   errValidate.Error(),
		})
	}
	//Upload file
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}
	//Logic jika tidak upload file

	var filename string
	if file != nil {
		filename = file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./utils/images/%s", filename))

		if errSaveFile != nil {
			log.Println("Failed to store file into images")
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	newBook := book_domain.BookModels{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
	}

	errCreateUser := database.DB.Create(&newBook).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
