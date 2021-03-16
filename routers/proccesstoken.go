package routers

import (
	"errors"
	"strings"

	"github.com/DalinarLG/twittor/db"
	"github.com/DalinarLG/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserID string

func ProccessToken(token string) (models.Claim, bool, string, error) {
	myKey := []byte("powermetal_blind_guardian")
	claims := models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckEmail(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}

		return claims, found, UserID, nil
	}
	 
	if !tk.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
