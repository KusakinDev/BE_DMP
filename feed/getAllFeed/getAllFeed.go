package getallfeed

import (
	dbA "back/db"
	vjwt "back/jwt/verefyJWT"
	gs "back/struct/goodsStruct"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllFeed(w http.ResponseWriter, r *http.Request) {
	_, err := vjwt.VerifyJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Запрос к базе данных для получения всех товаров, исключая поле item
	query := "SELECT id, id_s, title, description, price, date_pub, date_buy, is_buy FROM goods"
	rows, err := dbA.DB.Query(query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	goodsList := []gs.Goods{}
	for rows.Next() {
		var goods gs.Goods
		var dateBuy sql.NullString
		if err := rows.Scan(&goods.ID, &goods.IDS, &goods.Title, &goods.Description, &goods.Price, &goods.DatePub, &dateBuy, &goods.IsBuy); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goodsList); err != nil {
		log.Printf("Error encoding response to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
