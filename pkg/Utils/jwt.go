package utils

import (
	"strings"
	"time"

	"github.com/amanuel15/fiber_server/pkg/configs"
	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/amanuel15/fiber_server/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uint) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(configs.JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(c *fiber.Ctx) error {
	tokenString, found := c.GetReqHeaders()["Authorization"]
	if !found {
		return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Authorization header not found"})
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(strings.Split(tokenString, " ")[1], claims, func(token *jwt.Token) (interface{}, error) {
		return configs.JWT_SECRET, nil
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&interfaces.IResponse{Error: "Error parsing Token: " + err.Error()})
	}
	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(&interfaces.IResponse{Error: "Unauthorized token"})
	}
	// log.Println("Final token: ", claims)
	c.Locals("user", &models.DecodedUser{UserId: claims.UserId})
	return c.Next()
}
