package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("Error uncorrect credential")
	}

	return nil
}
