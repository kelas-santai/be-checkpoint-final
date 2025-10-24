package controllers

import (
	database "meeting4/databases"
	"meeting4/entity"
	"meeting4/pkg/tools"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Kita Ibaratkan Si Merchat ini dia register mandiri
func RegisterMerchant(c *fiber.Ctx) error {

	//jadi yang di butuhkan register yaitu email, password, nama
	//Jadi inputannya itu adalah berupa json
	//Saya akan menggunakan Cara dedngan Map
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal",
			"err":   err.Error(),
		})
	}
	//Mengisi data
	password := tools.GeneratePassword(data["password"])
	newData := entity.Merchant{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: password,
		NoTelpon: data["no_telpon"],
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	//Membuat Sebuah  Recorrd Data Baru
	database.DB.Create(&newData)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Membuat Data Merchant Baru",
	})
}

// Sebagai Kunci JWt
var jwtSecret = []byte("Testing123")

func LoginMerchant(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal",
			"err":   err.Error(),
		})
	}

	//Ambil Data Email Cocokan Ada Ga kalo engak ada tampil si user dengan email tersebut enggak ada
	email := data["email"]
	pass := data["password"]

	var mercant entity.Merchant
	//Cek databases
	database.DB.Where("email = ?", email).First(&mercant)
	//Baru main logic

	if mercant.Id == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "Data Tidak Di Temukan di Dalam Database",
		})
	}

	if tools.CheckPassword(pass, mercant.Password) {
		//Tembuatan JWT Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"nama":       mercant.Nama,
			"idMerchant": mercant.Id,
			"exp":        time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
			"iat":        time.Now().Unix(),
		})

		t, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal membuat token",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": t,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "PasswordSalah",
		})
	}

}

func GetUserByToken(c *fiber.Ctx) error {
	//Kita Akan Ambil Si Nama dan ID Merchan
	//nama := c.Locals("nama")
	idMerchant := c.Locals("idMerchant")

	//Kita Cari Si ID nya

	var merchant entity.Merchant

	database.DB.Where("id = ?", idMerchant).First(&merchant)

	//ubah vari able foto

	merchant.Foto = "http://localhost:3000/static/merchant/" + merchant.Foto

	//Kita Show

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengambil Data",
		"data":  merchant,
	})
}

//Meeting Rabu Kita Lanjut

//fungsi update dan juga delete

func UpdateMerchant(c *fiber.Ctx) error {

	//UpdateBerdaasrkan Token
	idMerchant := c.Locals("idMerchant")

	//Kita Cari Dulu datanya ada apa enggak
	var merchant entity.Merchant
	database.DB.Where("id =?", idMerchant).First(&merchant)

	//cek databya
	if merchant.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Data Merchant Tidak Di Temukan",
		})
	}

	//Bikin Untuk Update Datanya

	//
	var request entity.Merchant

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Pesan": "Gagal",
		})
	}

	//kita ubah data merchantnya
	merchant.Nama = request.Nama
	merchant.NoTelpon = request.NoTelpon
	merchant.Email = request.Email
	merchant.Lokasi = request.Lokasi
	merchant.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	//untuk foto itu adalah File
	fileFoto, _ := c.FormFile("foto")

	if fileFoto != nil {
		//Membuat Folder nya
		folderPath := "public/foto-merchant/" + tools.RemoveSpaces(fileFoto.Filename)
		//Simpan ke dalam foldedr si foto nya ini dan untuk di database nya itu adalah yang di smpan yaitu file name yang fotonya
		if err := c.SaveFile(fileFoto, folderPath); err != nil {
			return err
		}
		merchant.Foto = tools.RemoveSpaces(fileFoto.Filename)
	}

	//simpan perubahan
	database.DB.Save(&merchant)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Mengupdate Data Merchant",
	})
}

//
func DeletMerchant(c *fiber.Ctx) error {
	//Delete Berdasarkan Token
	idMerchant := c.Locals("idMerchant")

	//Kita Cari Dulu datanya ada apa enggak
	var merchant entity.Merchant
	database.DB.Where("id =?", idMerchant).First(&merchant)
	//cek databya
	if merchant.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Pesan": "Data Merchant Tidak Di Temukan",
		})
	}
	//Hapus Data
	database.DB.Delete(&merchant)
	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Menghapus Data Merchant",
	})

}
