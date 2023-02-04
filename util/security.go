package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword, candidatePassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword)); err == nil {
		return true
	} else {
		return false
	}
}
