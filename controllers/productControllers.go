package controllers

import (
	"meeting4/databases"
	"meeting4/entity"
	"meeting4/pkg/tools"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	// Parse form data
	var request entity.Product
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal parsing data",
			"err":   err.Error(),
		})
	}

	// Buat product baru
	newProduct := entity.Product{
		IdCategori: request.IdCategori,
		IdMerchant: request.IdMerchant,
		Nama:       request.Nama,
		Harga:      request.Harga,
		CreateAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// Handle file foto
	fileFoto, _ := c.FormFile("foto")
	if fileFoto != nil {
		// Membuat folder untuk menyimpan foto
		folderPath := "public/foto-product/" + tools.RemoveSpaces(fileFoto.Filename)
		// Simpan foto ke folder
		if err := c.SaveFile(fileFoto, folderPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Pesan": "Gagal menyimpan foto",
				"err":   err.Error(),
			})
		}
		newProduct.Foto = tools.RemoveSpaces(fileFoto.Filename)
	}

	// Simpan ke database
	database.DB.Create(&newProduct)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Pesan": "Berhasil Membuat Data Product",
		"data":  newProduct,
	})
}

func GetProduct(c *fiber.Ctx) error {
	// Ambil semua product dari database
	var products []entity.Product
	Query := database.DB
	idCategori := c.Query("idCategori")
	idMerchant := c.Query("idMerchant")
	if idMerchant != "" {
		Query = Query.Where("id_merchant = ?", idMerchant)
	}
	if idCategori != "" {
		Query = Query.Where("id_categori = ?", idCategori)
	}
	//
	Query.Find(&products)
	// Update URL foto untuk setiap product
	for i := range products {
		if products[i].Foto != "" {
			products[i].Foto = "http://localhost:3000/static/product/" + products[i].Foto
		}
	}
	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengambil Semua Data Product",
		"data":  products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	// Ambil ID dari query parameter
	id := c.Query("id")

	// Cari product berdasarkan ID
	var product entity.Product
	database.DB.Where("id = ?", id).First(&product)

	// Cek apakah product ditemukan
	if product.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Product tidak ditemukan",
		})
	}

	// Update URL foto
	if product.Foto != "" {
		product.Foto = "http://localhost:3000/static/product/" + product.Foto
	}

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengambil Data Product",
		"data":  product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	// Ambil ID dari query parameter
	id := c.Query("id")

	// Cari product berdasarkan ID
	var product entity.Product
	database.DB.Where("id = ?", id).First(&product)

	// Cek apakah product ditemukan
	if product.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Product tidak ditemukan",
		})
	}

	// Parse request body
	var request entity.Product
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal parsing data",
			"err":   err.Error(),
		})
	}

	// Update data product
	product.IdCategori = request.IdCategori
	product.IdMerchant = request.IdMerchant
	product.Nama = request.Nama
	product.Harga = request.Harga
	product.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// Handle file foto jika ada
	fileFoto, _ := c.FormFile("foto")
	if fileFoto != nil {
		// Membuat folder untuk menyimpan foto
		folderPath := "public/foto-product/" + tools.RemoveSpaces(fileFoto.Filename)
		// Simpan foto ke folder
		if err := c.SaveFile(fileFoto, folderPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Pesan": "Gagal menyimpan foto",
				"err":   err.Error(),
			})
		}
		product.Foto = tools.RemoveSpaces(fileFoto.Filename)
	}

	// Simpan perubahan ke database
	database.DB.Save(&product)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengupdate Data Product",
		"data":  product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	// Ambil ID dari query parameter
	id := c.Query("id")

	// Cari product berdasarkan ID
	var product entity.Product
	database.DB.Where("id = ?", id).First(&product)

	// Cek apakah product ditemukan
	if product.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Product tidak ditemukan",
		})
	}

	// Hapus product dari database
	database.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Menghapus Data Product",
	})
}
