package disableproductcard

import (
	"back/db"
	goodsstruct "back/struct/goodsStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisableProductCard(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}

	var good goodsstruct.Good
	if err := c.ShouldBindJSON(&good); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}

	err := db.DB.Model(&good).Where("id = ? AND id_s = ?", good.Id, id).Update("is_sell", false).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно обновлен"})
}
