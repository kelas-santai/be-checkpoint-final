package main

import (
	"log"
	"meeting4/controllers"
	"meeting4/databases"
	"meeting4/pkg/middleware"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(
		logger.New(),
		cors.New(),
	)
	//Untuk Koneksi Database
	database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	merchant := app.Group("/merchant")
	merchant.Post("/register", controllers.RegisterMerchant)
	merchant.Post("/login", controllers.LoginMerchant)
	merchant.Get("/getUserByToken", middleware.JWTProtected(), controllers.GetUserByToken)

	log.Fatal(app.Listen(":3000"))
}
