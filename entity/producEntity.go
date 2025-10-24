package entity

//si producck ii membupunyai 1 kategori dan juga mercant dia mempunya berbagai macam procuk
type Category struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Nama     string
	CreateAt string
	UpdateAt string
}

type Product struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	IdCategori uint
	IdMerchant uint
	Nama       string
	Harga      string
	Foto       string
	CreateAt   string
	UpdateAt   string
}
