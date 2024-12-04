package createproductcard

import (
	dbA "back/db"
	gs "back/struct/goodsStruct"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateProductCard(w http.ResponseWriter, r *http.Request) {
	var newProduct gs.Goods
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Неудалось декодировать JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("IDS:", newProduct.IDS)
	fmt.Println("Title:", newProduct.Title)
	fmt.Println("Description:", newProduct.Description)
	fmt.Println("Price:", newProduct.Price)
	fmt.Println("Image:", newProduct.Image)

	query := "INSERT INTO goods (id_s, title, description, price, image) VALUES ($1, $2, $3, $4, $5)"
	args := []interface{}{newProduct.IDS, newProduct.Title, newProduct.Description, newProduct.Price, newProduct.Image}

	_, err := dbA.DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "Ошибка добавления товара в таблицу", http.StatusBadRequest)
		log.Printf("%v", err)
		return
	}

}
