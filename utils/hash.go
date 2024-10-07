package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password),10)
	if err != nil {
		return "",err
	}
	return string(hashByte),nil
}

func ComparePassword(password,hashPassword string) bool {
	compare := bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))
	return compare == nil
}