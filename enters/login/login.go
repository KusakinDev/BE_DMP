package login

import (
	jwtconfig "back/config/jwtConfig"
	dbA "back/db"
	us "back/struct/userStruct"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user us.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	log.Println(user)
	// Проверка пользователя в базе данных
	query := "SELECT id, password FROM users WHERE name=$1"
	args := user.Name
	row := dbA.DB.QueryRow(query, args)
	var storedPassword string
	var userID int
	err := row.Scan(&userID, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Пока что пароль не захеширован, просто сравниваем строки
	if user.Password != storedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Генерация access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(30 * time.Second).Unix(), // Токен действует 30 секунд
	})
	accessTokenString, err := accessToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	// Генерация refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Токен действует 7 дней
	})
	refreshTokenString, err := refreshToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
		return
	}

	// Отправка токенов на фронт
	c.JSON(http.StatusOK, gin.H{
		"token":        accessTokenString,
		"refreshToken": refreshTokenString,
	})
}
