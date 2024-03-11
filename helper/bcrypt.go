package helper

import "golang.org/x/crypto/bcrypt"

func HassPassword(p string) string {
	salt := 10
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, password := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}