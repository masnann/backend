package photo_handler

import (
	"backend/cms/photo/photo_domain"
	"backend/utils"
	"backend/utils/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(photo_domain.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	// Validasi request
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal Input Data",
			"error":   errValidate.Error(),
		})
	}

	// Validasi keberadaan gambar (image) yang diunggah
	filenames := ctx.Locals("filenames")
	log.Println("filename", filenames)
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			newPhoto := photo_domain.Photo{
				Images:     filename,
				CategoryID: photo.CategoryId,
			}
			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("Some data not saved properly")
			}
		}

		// Membuat keterangan tentang foto yang diunggah
		photoInfo := make([]map[string]interface{}, 0)
		for _, filename := range filenamesData {
			info := map[string]interface{}{
				"filename": filename,
				"category": photo.CategoryId,
			}
			photoInfo = append(photoInfo, info)
		}

		return ctx.JSON(fiber.Map{
			"message": "success upload photo",
			"photos":  photoInfo,
		})
	}
}

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	var photo photo_domain.Photo
	photoId := ctx.Params("id")
	err := database.DB.Debug().First(&photo, "id = ?", photoId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Photo not found",
		})
	}
	errDeleteFile := utils.HandleRemoveFile(photo.Images)
	if errDeleteFile != nil {
		log.Println("Failed to delete some file")

	}
	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success delete photo",
	})

}
