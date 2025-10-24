package database

//Menyimpan Sebuah Fungsi Unutk  Mengkoneksikan ke database
import (
	"fmt"
	"meeting4/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Wajib Untuk Membuat Variable Global
var DB *gorm.DB

func Connect() {
	//secara default user databasse adalaah "root" dan passwordnya kosong
	//122.7 itu adalah host kita (Hst dari IP Server ), 3306 itu adadlag port default dari Mysql
	dsn := "root@tcp(127.0.0.1:3306)/meeting5?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database gagal terkoneksi")
	} else {
		fmt.Printf("database terkoneksi")
	}

	DB = db

	db.AutoMigrate(
		&entity.Admin{},
		&entity.User{},
		&entity.Merchant{},
		&entity.Category{},
		&entity.Product{},
		&entity.UsersProduct{},
	)
}
