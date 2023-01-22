package routes

import (
	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/database"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func createBlog(c *fiber.Ctx) error {
	decodedUser, ok := c.Locals("user").(*models.DecodedUser)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Unauthorized: User not found"})
	}
	var body interfaces.ICreateBlog
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: err.Error()})
	}
	if err := utils.ValidateBody(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(prepareResponse(err))
	}
	var blog models.Blog
	c.BodyParser(&blog)
	blog.AuthorID = decodedUser.UserId
	if result := database.DB.DB.Create(&blog); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(interfaces.IResponse{Error: result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(prepareResponse(blog))
}

func getBlogs(c *fiber.Ctx) error {

	blogs := []models.Blog{}
	database.DB.DB.Preload("References").Find(&blogs)

	return c.Status(200).JSON(prepareResponse(blogs))
}
