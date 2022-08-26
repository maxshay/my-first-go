package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login() fiber.Handler {
	return func (c *fiber.Ctx) error  {
		username, password := c.FormValue("username"), c.FormValue("password")

		if username == "" || password == "" {
			fmt.Println("Empty username or password")
			return c.Redirect("/login")
		}

		return c.Redirect("/dashboard")
	}
}