package verefyjwt

import (
	cl "back/struct/claimStruct"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func VerifyJWT(r *http.Request) (*cl.Claims, error) {

	var jwtKey = []byte("A6F8e9BD3s72DFb8c9E5zQdP")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	tokenStr := splitToken[1]
	claims := &cl.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("unauthorized")
		}
		return nil, fmt.Errorf("bad request")
	}

	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
