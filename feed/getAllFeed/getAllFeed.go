package getallfeed

import (
	dbA "back/db"
	goodsstruct "back/struct/goodsStruct"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFeed(c *gin.Context) {

	goodsList := []goodsstruct.Good{}
	err := dbA.DB.
		Where("is_buy = ? AND is_sell = ?", false, true).
		Preload("User").
		Find(&goodsList).
		Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товары не найдены"})
		return
	}

	for i := range goodsList {
		goodsList[i].User.Id = 0
		goodsList[i].User.Password = ""
		goodsList[i].User.Email = ""
	}

	c.JSON(http.StatusOK, goodsList)
}
