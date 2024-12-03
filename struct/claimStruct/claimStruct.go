package claimstruct

import "github.com/dgrijalva/jwt-go"

// Структура для обработки JWT claims
type Claims struct {
	ID   int `json:"id"`
	Type int `json:"type"`
	jwt.StandardClaims
}
