package db

import "golang.org/x/crypto/bcrypt"

//EncryptPass encrypts password
func EncryptPass(p string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), cost)
	return string(bytes), err
}
