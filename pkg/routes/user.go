package routes

import (
	"log"

	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/database"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func createUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(interfaces.IResponse{Error: err.Error()})
	}
	if err := utils.ValidateBody(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user.Password = utils.HashPassword(user.Password)

	if result := database.DB.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(interfaces.IResponse{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(prepareResponse(user))
}

func getUsers(c *fiber.Ctx) error {
	decodedUser, ok := c.Locals("user").(*models.DecodedUser)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Unauthorized: User not found"})
	}
	log.Println("Decoded User: ", decodedUser)

	users := []models.User{}
	database.DB.DB.Find(&users)

	return c.Status(200).JSON(prepareResponse(users))
}

func deleteAllUsers(c *fiber.Ctx) error {
	database.DB.DB.Unscoped().Where("").Delete(&models.User{})
	return c.Status(200).JSON("Success")
}

func deleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(409).JSON(err.Error())
	}
	database.DB.DB.Delete(&models.User{}, id)
	return c.Status(200).JSON("Success")
}
