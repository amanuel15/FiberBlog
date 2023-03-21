package routes

import (
	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/database"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func createComment(c *fiber.Ctx) error {
	decodedUser, ok := c.Locals("user").(*models.DecodedUser)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Unauthorized: User not found"})
	}
	var body interfaces.ICreateComment
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: err.Error()})
	}
	if err := utils.ValidateBody(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var blog models.Blog
	database.DB.DB.Find(&blog, "id = ?", body.BlogID)
	if blog.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&interfaces.IResponse{Error: "Blog not found"})
	}
	var comment models.Comment
	c.BodyParser(&comment)
	comment.WriterID = decodedUser.UserId

	if result := database.DB.DB.Create(&comment); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(interfaces.IResponse{Error: result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}
