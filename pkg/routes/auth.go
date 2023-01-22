package routes

import (
	"log"

	utils "github.com/amanuel15/fiber_server/pkg/Utils"
	"github.com/amanuel15/fiber_server/pkg/database"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	var login models.Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: err.Error()})
	}
	if err := utils.ValidateBody(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	var user models.User
	database.DB.DB.Find(&user, "name = ?", login.Name)

	if valid := utils.CheckPasswordHash(login.Password, user.Password); valid == true {
		token, _ := utils.GenerateJWT(user.ID)
		log.Println("Token is: ", token)
		return c.Status(fiber.StatusOK).JSON(prepareResponse(token))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Wrong name or password"})
}
