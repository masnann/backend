package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mime/multipart"
	"os"
)

const DefaultPathAssetImage = "./utils/images/"

func HandleSingleFile(ctx *fiber.Ctx) error {

	//Upload file
	file, errFile := ctx.FormFile("images")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}
	//Logic jika tidak upload file

	var filename *string
	if file != nil {
		errCheckContentType := CheckContentType(file, "image/jpg", "image/png")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}
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

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}
	return nil
}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {

	if len((contentTypes)) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed file type")
	} else {
		return errors.New("Nothing file to be checking")
	}
}
