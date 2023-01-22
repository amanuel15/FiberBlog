package routes

import (
	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user", utils.VerifyToken)
	blog := v1.Group("/blog", utils.VerifyToken)
	auth := v1.Group("/auth")

	user.Post("", createUser)
	user.Get("", getUsers)
	user.Delete("", deleteAllUsers)
	user.Delete("/:id", deleteUser)

	auth.Post("/login", login)

	blog.Post("", createBlog)
	blog.Get("", getBlogs)
}

func prepareResponse(value interface{}) interface{} {
	return &interfaces.IResponse{Data: value}
}
