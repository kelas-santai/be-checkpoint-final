package controllers

import (
	 "meeting4/databases"
	"meeting4/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

//CRUD Categori
func CreateKategori(c *fiber.Ctx) error {
	// Parse request body
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal parsing data",
			"err":   err.Error(),
		})
	}

	// Buat kategori baru
	newCategory := entity.Category{
		Nama:     data["nama"],
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simpan ke database
	database.DB.Create(&newCategory)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Pesan": "Create Kategori Berhasil",
		"data":  newCategory,
	})
}

func GetKategori(c *fiber.Ctx) error {
	// Ambil semua kategori dari database
	var categories []entity.Category
	database.DB.Find(&categories)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil mengambil semua data kategori",
		"data":  categories,
	})
}

func GetKategoriByID(c *fiber.Ctx) error {
	// Ambil ID dari parameter
	id := c.Query("id")

	// Cari kategori berdasarkan ID
	var category entity.Category
	database.DB.Where("id = ?", id).First(&category)

	// Cek apakah kategori ditemukan
	if category.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Kategori tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil mengambil data kategori",
		"data":  category,
	})
}

func UpdateKategori(c *fiber.Ctx) error {
	// Ambil ID dari parameter
	id := c.Params("id")

	// Cari kategori berdasarkan ID
	var category entity.Category
	database.DB.Where("id = ?", id).First(&category)

	// Cek apakah kategori ditemukan
	if category.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Kategori tidak ditemukan",
		})
	}

	// Parse request body
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal parsing data",
			"err":   err.Error(),
		})
	}

	// Update data kategori
	category.Nama = data["nama"]
	category.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// Simpan perubahan ke database
	database.DB.Save(&category)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil mengupdate kategori",
		"data":  category,
	})
}

func DeleteKategori(c *fiber.Ctx) error {
	// Ambil ID dari parameter
	id := c.Params("id")

	// Cari kategori berdasarkan ID
	var category entity.Category
	database.DB.Where("id = ?", id).First(&category)

	// Cek apakah kategori ditemukan
	if category.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Kategori tidak ditemukan",
		})
	}

	// Hapus kategori dari database
	database.DB.Delete(&category)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil menghapus kategori",
	})
}
