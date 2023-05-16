package route

import (
	"backend/cms/auth/auth_handler"
	"backend/cms/book/book_handler"
	"backend/cms/photo/photo_handler"
	"backend/cms/user/user_handler"
	"backend/middleware"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	// Route User
	r.Get("/user", middleware.Auth, user_handler.UserHandlerGetAll)
	r.Get("/user/:id", user_handler.UserHandlerGetById)
	r.Post("user/create", user_handler.UserCreate)
	r.Put("/user/update/:id", user_handler.UserHandlerUpdate)
	r.Put("/user/update-email/:id", user_handler.UserHandlerUpdateEmail)
	r.Delete("/user/delete/:id", user_handler.UserHandlerDelete)

	// Route Login
	r.Post("/login", auth_handler.Login)

	// Route Book
	r.Post("/book/create", utils.HandleSingleFile, book_handler.BookHandlerCreate)

	// Route Galery
	r.Post("/gallery/create", utils.HandleMultipleFile, photo_handler.PhotoHandlerCreate)
}
