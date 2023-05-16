package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func HandleSingleFile(ctx *fiber.Ctx) error {

	//Upload file
	file, errFile := ctx.FormFile("images")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}
	//Logic jika tidak upload file

	var filename *string
	if file != nil {
		filename = &file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./utils/images/%s", *filename))

		if errSaveFile != nil {
			log.Println("Failed to store file into images")
		}
	} else {
		log.Println("Nothing file to uploading.")
	}
	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error Read Multiple File, Error = ", errForm)
	}

	files := form.File["photos"]
	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)
			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./utils/images/%s", filename))

			if errSaveFile != nil {
				log.Println("Failed to store file into images")
			}
		} else {
			log.Println("Nothing file to uploading.")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}

	}
	ctx.Locals("filenames", filenames)
	return ctx.Next()
}
