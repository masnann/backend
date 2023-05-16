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
	// Validasi Required Images
	var filenameString string
	filename := ctx.Locals("filename")
	log.Println("filename", filename)
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := book_domain.BookModels{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
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
