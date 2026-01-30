package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"os"
)

func JwtMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Cookies("access_token")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	}

	token, err := jwt.ParseWithClaims(tokenStr, &domain.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": token})
	}

	claims, ok := token.Claims.(*domain.JwtCustomClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid claims"})
	}

	c.Locals("id", claims.ID)

	return c.Next()
}
