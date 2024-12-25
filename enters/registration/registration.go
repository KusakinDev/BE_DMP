package registration

import (
	"back/db"
	userstruct "back/struct/userStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var user userstruct.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пустые поля"})
		return
	}

	errUserCheck := db.DB.Where("email = ?", user.Email).First(&user).Error
	if errUserCheck == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким email уже существует"})
		return
	}

	err := user.HashPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось захешировать пароль"})
		return
	}

	errUser := db.DB.Create(&user).Error
	if errUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно создан"})
}
