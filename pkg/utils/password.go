package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Fungsi untuk mengenerate hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Fungsi untuk menvalidasi password yang sudah di hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
