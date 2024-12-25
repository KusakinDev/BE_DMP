package login

import (
	jwtconfig "back/config/jwtConfig"
	dbA "back/db"
	us "back/struct/userStruct"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userFront us.User
	if err := c.ShouldBindJSON(&userFront); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	log.Println(userFront)

	var user us.User
	// Проверка пользователя в базе данных
	err := dbA.DB.Where("name = ?", userFront.Name).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Пока что пароль не захеширован, просто сравниваем строки
	if userFront.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Генерация access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(30 * time.Second).Unix(), // Токен действует 30 секунд
	})
	accessTokenString, err := accessToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	// Генерация refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
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
