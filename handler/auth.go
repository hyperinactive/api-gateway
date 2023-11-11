package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *fiber.Ctx) error {
	var input SignInBody

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// read input
	username := input.Username
	password := input.Password

	// find user by username
	// -----------------------------------------------------------------------------------------------------------
	user, err := GetUserByUsername(username)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Credentials", "error": ""})
	}

	// compare password
	// -----------------------------------------------------------------------------------------------------------
	passValid := ComparePasswordHash(password, user.Password)

	if !passValid {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Credentials", "error": ""})
	}

	// generate token
	// -----------------------------------------------------------------------------------------------------------
	userId := user.ID.String()
	claims, err := createUserClaim(userId, "user")

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

func SignUp(c *fiber.Ctx) error {
	var input SignUpBody

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := CreateUser(input.Username, input.Email, input.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
