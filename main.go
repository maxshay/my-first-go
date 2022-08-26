package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/maxshay/fiber-test/controllers"
)


func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	
	// Home
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"World": "World",
		})
	})

	// Dash
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dash", fiber.Map{})
	})

	// Login
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	// api
	app.Post("/api/login", controllers.Login())
	
	log.Fatal(app.Listen(":3000"))
}