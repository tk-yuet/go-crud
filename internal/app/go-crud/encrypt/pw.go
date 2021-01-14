package encrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncrpyptPw(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func IsCorrectPassword(hashedPwd string, input string) bool {
	inputPw := []byte(input)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, inputPw)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
