package photo_handler

import (
	"backend/cms/photo/photo_domain"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(photo_domain.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal Input Data",
			"error":   errValidate.Error(),
		})
	}
	// Validasi Required Images
	var filenameString string
	filenames := ctx.Locals("filenames")
	log.Println("filename", filenames)
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenames)
	}
	log.Println(filenameString)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
