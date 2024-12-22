package getcart

import (
	"back/db"
	cartstruct "back/struct/cartStruct"
	goodsstruct "back/struct/goodsStruct"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {

	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	id = int(id.(float64))

	var carts []cartstruct.Cart
	err := db.DB.
		Preload("Good.User"). // Информация о пользователе товара
		Preload("User").      // Информация о пользователе корзины
		Where("id_u = ?", id).
		Find(&carts).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Корзина не найдена"})
		return
	}

	var goods []goodsstruct.Good
	for _, item := range carts {
		item.Good.User.ID = 0
		item.Good.User.Password = ""
		item.Good.User.Email = ""
		goods = append(goods, item.Good)
	}

	log.Println("goods ", goods)
	c.JSON(http.StatusOK, goods)
}
