package addtocart

import (
	"back/db"
	cartstruct "back/struct/cartStruct"
	goodsstruct "back/struct/goodsStruct"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddToCart(c *gin.Context) {
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

	var cart1 cartstruct.Cart
	errfind := db.DB.Where("id_u = ? AND id_p = ?", id, newProduct.ID).First(&cart1)
	if errfind.Error == nil {
		c.JSON(http.StatusAlreadyReported, gin.H{"error": "Товар уже добавлен в корзину"})
		return
	}
	if !errors.Is(errfind.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusAlreadyReported, gin.H{"error": "Товар уже добавлен в корзину"})
		return
	}

	var cart cartstruct.Cart
	cart.Id_u = int(id.(float64))
	cart.Id_p = newProduct.ID
	cart.Date = time.Now().Format("2006-01-02")

	err1 := db.DB.Create(&cart).Error
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить в корзину"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно добавлен в корзину"})
}
