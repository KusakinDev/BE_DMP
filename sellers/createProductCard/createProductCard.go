package createproductcard

import (
	dbA "back/db"
	goodsstruct "back/struct/goodsStruct"
	itemstruct "back/struct/itemStruct"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProductCard(c *gin.Context) {

	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}

	var newProduct goodsstruct.Good
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}
	newProduct.Id = 0
	newProduct.IdS = int(id.(float64))
	newProduct.DatePub = time.Now().Format("2006-01-02")
	newProduct.IsBuy = false
	err := dbA.DB.Create(&newProduct).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить товар", "details": err.Error()})
		return
	}

	var item itemstruct.Item
	err = dbA.DB.Model(&item).Where("id = ?", newProduct.IdI).Update("id_g", newProduct.Id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить запись в таблице items", "details": err.Error()})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно добавлен"})

}
