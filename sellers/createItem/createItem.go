package createitem

import (
	cryproconfig "back/crypto/cryproConfig"
	"back/db"
	itemstruct "back/struct/itemStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var createItem itemstruct.Item
	if err := c.ShouldBindJSON(&createItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}
	createItem.Encode(cryproconfig.KEY)
	err := db.DB.Create(&createItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
		return
	}

	createItem.Content = ""
	c.JSON(http.StatusOK, createItem)
}
