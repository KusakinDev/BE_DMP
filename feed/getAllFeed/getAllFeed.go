package getallfeed

import (
	dbA "back/db"
	gs "back/struct/goodsStruct"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFeed(c *gin.Context) {

	query := "SELECT id, id_s, title, description, price, date_pub, date_buy, is_buy, image FROM goods"
	rows, err := dbA.DB.Query(query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	goodsList := []gs.Goods{}
	for rows.Next() {
		var goods gs.Goods
		var dateBuy sql.NullString
		if err := rows.Scan(&goods.ID, &goods.IDS, &goods.Title, &goods.Description, &goods.Price, &goods.DatePub, &dateBuy, &goods.IsBuy, &goods.Image); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if dateBuy.Valid {
			goods.DateBuy = &dateBuy.String
		} else {
			goods.DateBuy = nil
		}
		goodsList = append(goodsList, goods)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, goodsList)
}
