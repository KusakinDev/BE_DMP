package buygoods

import (
	"back/db"
	goodsstruct "back/struct/goodsStruct"
	historystruct "back/struct/historyStruct"
	userstruct "back/struct/userStruct"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BuyGoods(c *gin.Context) {

	var user userstruct.User
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	idInt := int(id.(float64))

	user.Id = idInt

	var buyProduct goodsstruct.Good
	if err := c.ShouldBindJSON(&buyProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неудалось декодировать JSON"})
		return
	}

	var checkProduct goodsstruct.Good
	errfind := db.DB.Where("id = ? AND is_buy = ? AND is_sell = ?", buyProduct.Id, false, true).First(&checkProduct)
	if errfind.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден или уже куплен"})
		return
	}

	var userBuyer userstruct.User
	errUser := db.DB.Where("id = ?", user.Id).First(&userBuyer).Error
	if errUser != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	var userSeller userstruct.User
	errSeller := db.DB.Where("id = ?", checkProduct.IdS).First(&userSeller).Error
	if errSeller != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продавец не найден"})
		log.Println(http.StatusNotFound, gin.H{"error": "Продавец не найден"})
		return
	}

	if userBuyer.DebugWallet >= checkProduct.Price {
		userBuyer.DebugWallet -= checkProduct.Price
		userSeller.DebugWallet += checkProduct.Price

		errBuy := db.DB.Model(&goodsstruct.Good{}).Where("id = ?", checkProduct.Id).Updates(map[string]interface{}{"is_buy": true, "id_u": id}).Error
		if errBuy != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось купить товар"})
			return
		}

		errWaller := db.DB.Model(&userstruct.User{}).Where("id = ?", user.Id).Updates(map[string]interface{}{"debug_wallet": userBuyer.DebugWallet}).Error
		if errWaller != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить баланс"})
			return
		}

		errWallerSeller := db.DB.Model(&userstruct.User{}).Where("id = ?", userSeller.Id).Updates(map[string]interface{}{"debug_wallet": userSeller.DebugWallet}).Error
		if errWallerSeller != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить баланс продавца"})
			return
		}

		var history historystruct.History
		history.IdU = &idInt
		history.IdG = &checkProduct.Id
		history.IdI = checkProduct.IdI
		history.Date = time.Now().Format("2006-01-02")
		errHistory := db.DB.Create(&history).Error
		if errHistory != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать историю"})
			return
		}
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Недостаточно средств"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно куплен"})
}
