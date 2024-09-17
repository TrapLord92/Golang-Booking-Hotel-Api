package api

import (
	"github.com/TrapLord92/Golang-Booking-Hotel-Api.git/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrUnAuthorized()
	}

	if !user.IsAdmin {
		return ErrUnAuthorized()
	}

	return c.Next()
}
