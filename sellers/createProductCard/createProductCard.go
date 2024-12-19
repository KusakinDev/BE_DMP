package createproductcard

import (
	dbA "back/db"
	gs "back/struct/goodsStruct"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProductCard(c *gin.Context) {
	/*id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}*/

	var newProduct gs.Goods
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}

	fmt.Println("IDS:", newProduct.IDS)
	fmt.Println("Title:", newProduct.Title)
	fmt.Println("Description:", newProduct.Description)
	fmt.Println("Price:", newProduct.Price)
	fmt.Println("Image:", newProduct.Image)

	currentDate := time.Now().Format("2006-01-02")
	query := "INSERT INTO goods (id_s, title, description, price, date_pub, is_buy, image) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	args := []interface{}{newProduct.IDS, newProduct.Title, newProduct.Description, newProduct.Price, currentDate, false, newProduct.Image}

	_, err := dbA.DB.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка добавления товара в таблицу"})
		return
	}

}
