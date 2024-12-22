package getprofile

import (
	"back/db"
	userstruct "back/struct/userStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	id = int(id.(float64))

	var user userstruct.User
	err := db.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
