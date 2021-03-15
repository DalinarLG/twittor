package jwt

import (
	"time"

	"github.com/DalinarLG/twittor/models"
	"github.com/dgrijalva/jwt-go"
)

//GenerateToken function
func GenerateToken(u models.User) (string, error) {

	myKey := []byte("powermetal_blind_guardian")

	payload := jwt.MapClaims{
		"email":    u.Email,
		"name":     u.Name,
		"lastname": u.Lastname,
		"birthday": u.Birthday,
		"bio":      u.Bio,
		"location": u.Location,
		"_id":      u.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
