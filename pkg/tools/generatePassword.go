package tools

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//Inputan Berupa String Outputnya adalah String

// sebuah Bycript
func GeneratePassword(pass string) string {
	// Generate hash dari password dengan cost (kompleksitas)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error saat generate hash:", err)
		return ""
	}

	return string(hash)
}

func CheckPassword(pass, passDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passDB), []byte(pass))
	return err == nil
}
