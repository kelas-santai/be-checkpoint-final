package main

import (
	"log"
	"meeting4/controllers"
	database "meeting4/databases"
	"meeting4/pkg/middleware"
	"meeting4/pkg/tools"

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

	//Generate Folder
	tools.GeneraFolder()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/static/merchant/:namaFile", tools.StatifFotoMerchnt)

	merchant := app.Group("/merchant")
	merchant.Post("/register", controllers.RegisterMerchant)
	merchant.Post("/login", controllers.LoginMerchant)
	merchant.Get("/getUserByToken", middleware.JWTProtected(), controllers.GetUserByToken)
	merchant.Put("/updateMerchant", middleware.JWTProtected(), controllers.UpdateMerchant)
	merchant.Delete("/deleteMerchant", middleware.JWTProtected(), controllers.DeletMerchant)

	category := app.Group("/category")
	category.Post("/createKategori", controllers.CreateKategori)
	category.Get("/getKategori", controllers.GetKategori)
	category.Get("/getKategoriByID", controllers.GetKategoriByID)
	category.Put("/updateKategori/:id", controllers.UpdateKategori)
	category.Delete("/deleteKategori/:id", controllers.DeleteKategori)

	product := app.Group("/product")
	product.Post("/createProduct", middleware.JWTProtected(), controllers.CreateProduct)
	product.Get("/getProduct", middleware.JWTProtected(), controllers.GetProduct)
	product.Get("/getProductByID", middleware.JWTProtected(), controllers.GetProductById)
	product.Put("/updateProduct", middleware.JWTProtected(), controllers.UpdateProduct)
	product.Delete("/deleteProduct", middleware.JWTProtected(), controllers.DeleteProduct)

	log.Fatal(app.Listen(":3000"))
}
