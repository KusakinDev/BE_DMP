package generatejwt

import (
	cl "back/struct/claimStruct"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID int, userType int) (string, error) {

	var jwtKey = []byte("A6F8e9BD3s72DFb8c9E5zQdP")

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &cl.Claims{
		ID:   userID,
		Type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
