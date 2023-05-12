package user_handler

import (
	"backend/cms/user/user_domain"
	"backend/cms/user/user_response"
	"backend/utils"
	"backend/utils/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []user_domain.UserModel

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user user_domain.UserModel
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	userResponse := user_response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserCreate(ctx *fiber.Ctx) error {
	user := new(user_domain.RequestUserInsert)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// Validasi panjang password minimal 6 karakter
	if len(user.Password) < 6 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Password harus lebih dari 6 huruf",
		})
	}
	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal Input Data",
			"error":   errValidate.Error(),
		})
	}
	newUser := user_domain.UserModel{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	newUser.Password = hashedPassword
	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})

}

func UserHandlerUpdate(ctx *fiber.Ctx) error {

	var user user_domain.UserModel
	userId := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	userUpdate := new(user_domain.RequestUserUpdate)
	if err := ctx.BodyParser(userUpdate); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}

	if userUpdate.Address != "" {
		user.Address = userUpdate.Address
	}

	if userUpdate.Phone != "" {
		user.Phone = userUpdate.Phone
	}
	if userUpdate.Password != "" {
		// Hash password baru
		hashedPassword, err := utils.HashingPassword(userUpdate.Password)
		if err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		user.Password = hashedPassword
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	var user user_domain.UserModel

	userId := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	var isEmailUserExist user_domain.UserModel
	userUpdate := new(user_domain.RequsetUserUpdateEmail)
	errIsEmailUserExist := database.DB.First(&isEmailUserExist, "email = ?", userUpdate.Email).Error
	if errIsEmailUserExist == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "Email already used",
		})
	}

	if err := ctx.BodyParser(userUpdate); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	user.Email = userUpdate.Email

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	var user user_domain.UserModel
	userId := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	errUpdate := database.DB.Delete(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}
	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
