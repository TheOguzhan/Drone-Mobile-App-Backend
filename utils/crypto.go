package utils

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent"
)

func HashPassword(password ent.Value) ent.Value {
	salt := os.Getenv("PASSWORD_HASH")
	if salt == "" {
		panic("PASSWORD_HASH environment variable not set")
	}
	saltedPassword := []byte(fmt.Sprintf("%x", password) + salt)
	hash := sha256.Sum256(saltedPassword)

	return fmt.Sprintf("%x", hash)
}

func HashToString(password string) string {
	salt := os.Getenv("PASSWORD_HASH")
	if salt == "" {
		panic("PASSWORD_HASH environment variable not set")
	}
	saltedPassword := []byte(fmt.Sprintf("%x", password) + salt)
	hash := sha256.Sum256(saltedPassword)

	return fmt.Sprintf("%x", hash)
}
