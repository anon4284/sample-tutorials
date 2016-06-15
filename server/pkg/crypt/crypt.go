package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pbyte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pbyte, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashedPassword)
}
func ComparePasswords(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
