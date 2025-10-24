package tools

//GenerateFolder Untuk Membuat Sebuah Folder
import (
	"os"
)

func GeneraFolder() {

	//Kumpulan Kumpulan dari Folder Yang Akan Di Buat
	folderNames := []string{
		"public/foto-merchant",
		"public/foto-product",
	}
	for _, folderName := range folderNames {
		err := GenerateFolder(folderName)
		if err != nil {
			panic(err)
		}
	}
}
func GenerateFolder(folderName string) error {
	//Membuat Sebuah Folder
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
