package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/services"
)

func NewAuthMiddleware(secret string, jwtService *services.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authorization header
		authHeader := c.Get("Authorization")

		// Check if the authorization header is present
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Unauthorized",
			})
		}

		// Split the authorization header to get the token
		authParts := strings.Fields(authHeader)
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Invalid Authorization header format",
			})
		}

		// Extract the token from authorization header
		tokenString := authParts[1]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		// Check for errors during token parsing
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Invalid token",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Internal Server Error",
			})
		}

		_, err = jwtService.RetrieveJwtTokenFromRedis(claims["id"].(string))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
				Success: false,
				Message: "Token tidak ditemukan",
			})
		}

		c.Locals("user", claims)

		return c.Next()
	}
}