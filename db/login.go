package db

import (
	"github.com/DalinarLG/twittor/models"	
	"golang.org/x/crypto/bcrypt"
)

// Login function
func Login(email, pass string) (models.User, bool) {

	user, found, _ := CheckEmail(email)
	if !found {
		return user, false
	}
	
	res := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if res != nil {
		return user, false
	}

	return user, true

}
