package routes

import (
	"strconv"

	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/database"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	database.DB.DB.Preload("References", func(db *gorm.DB) *gorm.DB {
		return db.Select("link, blog_id")
	}).Find(&blogs)

	return c.Status(200).JSON(prepareResponse(blogs))
}

func updateBlog(c *fiber.Ctx) error {
	var body interfaces.IUpdateBlog
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: err.Error()})
	}
	if err = c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: err.Error()})
	}
	if err := utils.ValidateBody(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	var blog models.Blog
	database.DB.DB.Find(&blog, "id = ?", id)
	if blog.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&interfaces.IResponse{Error: "Blog not found"})
	}
	c.BodyParser(&blog)
	if result := database.DB.DB.Save(&blog); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(interfaces.IResponse{Error: result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(prepareResponse(blog))
}
