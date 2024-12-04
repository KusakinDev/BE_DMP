package createproductcard

import (
	dbA "back/db"
	vjwt "back/jwt/verefyJWT"
	gs "back/struct/goodsStruct"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CreateProductCard(w http.ResponseWriter, r *http.Request) {
	_, err := vjwt.VerifyJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Printf("%v", err)
		return
	}
	var newProduct gs.Goods
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Неудалось декодировать JSON", http.StatusBadRequest)
		log.Printf("%v", err)
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

	_, err = dbA.DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "Ошибка добавления товара в таблицу", http.StatusBadRequest)
		log.Printf("%v", err)
		return
	}

}
