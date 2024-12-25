package getmygoods

import (
	"back/db"
	goodsstruct "back/struct/goodsStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyGoods(c *gin.Context) {

	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	id = int(id.(float64))

	var goods []goodsstruct.Good
	err := db.DB.Where("id_s = ?", id).Find(&goods).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товары не найдены"})
		return
	}

	c.JSON(http.StatusOK, goods)
}
