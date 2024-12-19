package refreshjwt

import (
	jwtconfig "back/config/jwtConfig"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	var reqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil || reqBody.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	// Проверяем валидность refresh токена
	token, err := jwt.Parse(reqBody.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtconfig.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || token.Method != jwt.SigningMethodHS256 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Извлекаем ID пользователя из токена
	userID, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Генерация нового access токена
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  int(userID),
		"exp": time.Now().Add(30 * time.Second).Unix(), // Новый срок действия: 30 секунд
	})
	accessTokenString, err := newAccessToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	// Генерация нового refresh токена
	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  int(userID),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Новый срок действия: 7 дней
	})
	refreshTokenString, err := newRefreshToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
		return
	}

	log.Println("token", accessTokenString, "refresh_token", refreshTokenString)
	// Отправка новой пары токенов на фронт
	c.JSON(http.StatusOK, gin.H{
		"token":         accessTokenString,
		"refresh_token": refreshTokenString,
	})
}
