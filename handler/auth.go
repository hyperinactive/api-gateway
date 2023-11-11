package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SignIn(c *fiber.Ctx) error {
	var input SignInBody

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	identity := input.Identity
	pass := input.Password

	println(identity, pass)

	// TODO: bs
	if identity != "user" || pass != "pass" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, err := createUserClaim(identity, "user")

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
