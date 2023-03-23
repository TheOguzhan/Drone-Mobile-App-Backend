package middleware

import (
	"context"
	"strings"

	"github.com/TheOguzhan/Drone-Mobile-App-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

func NewJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtHeader := c.GetReqHeaders()["Authorization"]
		jwtHeaderParts := strings.Split(jwtHeader, " ")
		var jwt_token string
		if len(jwtHeaderParts) > 1 {
			jwt_token = jwtHeaderParts[1]
			// Do something with the JWT token here
		} else {
			return c.Next()
		}

		ok, claims, err := utils.VerifyToken(jwt_token)

		if !ok || err != nil {
			return c.Next()
		}

		_ = context.WithValue(c.Context(), "user_id", claims.User_ID)

		c.Context().SetUserValue("user_id", claims.User_ID)
		return c.Next()
	}
}
