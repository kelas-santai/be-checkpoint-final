package tools

import "github.com/gofiber/fiber/v2"

func StatifFotoMerchnt(c *fiber.Ctx)error {
	//Paramter 
	namaFile := c.Params("namaFile")
	return c.SendFile("public/foto-merchant/" + namaFile)

}