package createproductcard

import (
	dbA "back/db"
	goodsstruct "back/struct/goodsStruct"
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
	newProduct.ID = 0
	newProduct.IDS = int(id.(float64))
	newProduct.DatePub = time.Now().Format("2006-01-02")
	newProduct.IsBuy = false
	err := dbA.DB.Create(&newProduct).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить товар", "details": err.Error()})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно добавлен"})

}
