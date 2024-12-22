package removefromcart

import (
	"back/db"
	cartstruct "back/struct/cartStruct"
	goodsstruct "back/struct/goodsStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveFromCart(c *gin.Context) {
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

	var cart cartstruct.Cart
	id = int(id.(float64))

	err := db.DB.Where("id_u = ? AND id_p = ?", id, newProduct.ID).Delete(&cart).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить из корзины"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален из корзины"})
}
