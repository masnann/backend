package auth_handler

import (
	"backend/cms/auth/auth_request"
	"backend/cms/user/user_domain"
	"backend/utils"
	"backend/utils/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Login(ctx *fiber.Ctx) error {

	loginRequest := new(auth_request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}
	log.Println(loginRequest)
	//Validasi request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal",
			"error":   errValidate.Error(),
		})
	}

	//Validasi Email
	var user user_domain.UserModel
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email atau password salah",
		})
	}

	//Validasi Password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email atau password salah",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 72).Unix()

	if user.Email == "Anan@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}

	return ctx.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
		"data":    user,
	})
}
