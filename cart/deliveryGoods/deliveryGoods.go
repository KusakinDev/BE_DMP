package deliverygoods

import (
	cryproconfig "back/crypto/cryproConfig"
	"back/db"
	"back/email/email"
	goodsstruct "back/struct/goodsStruct"
	historystruct "back/struct/historyStruct"
	userstruct "back/struct/userStruct"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeliveryGoods(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	id = int(id.(float64))

	var delProduct goodsstruct.Good
	if err := c.ShouldBindJSON(&delProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}

	errGood := db.DB.Where("id = ?", delProduct.Id).First(&delProduct).Error
	if errGood != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var history historystruct.History
	errHistory := db.DB.Preload("Item").Where("id_u = ? AND id_g = ?", id, delProduct.Id).First(&history).Error
	if errHistory != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var user userstruct.User
	errUser := db.DB.Where("id = ?", id).First(&user).Error
	if errUser != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	log.Println("history.Item.Content ", history.Item.Content)
	history.Item.Decode(cryproconfig.KEY)
	log.Println("history.Item.Content ", history.Item.Content)
	errEmail := email.SendEmail(user.Email, delProduct.Title, history.Item.Content)
	if errEmail != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось доставить товар"})
		log.Print("error Не удалось доставить товар: ", errEmail)
		return
	}

}
