package getcart

import (
	"back/db"
	verefyjwt "back/jwt/verefyJWT"
	claimstruct "back/struct/claimStruct"
	goodsstruct "back/struct/goodsStruct"
	"encoding/json"
	"log"
	"net/http"
)

type Cart struct {
	Id   int
	Id_u int
	Id_p int
	Date string
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	claim := &claimstruct.Claims{}
	var err error
	claim, err = verefyjwt.VerifyJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Printf("%v", err)
		return
	}

	queryCart := "SELECT * FROM cart WHERE id_u = $1"
	argCart := claim.ID
	rowsCart, err := db.DB.Query(queryCart, argCart)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rowsCart.Close()

	var goodsList []goodsstruct.GoodsToFront
	for rowsCart.Next() {
		var cart Cart
		if err := rowsCart.Scan(&cart.Id, &cart.Id_u, &cart.Id_p, &cart.Date); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		queryGoods := "SELECT id, id_s, title, description, price, is_buy, image FROM goods WHERE id = $1"
		argGoods := cart.Id_p
		rowsGoods, err := db.DB.Query(queryGoods, argGoods)
		if err != nil {
			log.Printf("Error querying database: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rowsGoods.Close()

		var goods goodsstruct.GoodsToFront
		if rowsGoods.Next() {

			var id int
			if err := rowsGoods.Scan(
				&goods.ID,
				&id,
				&goods.Title,
				&goods.Description,
				&goods.Price,
				&goods.IsBuy,
				&goods.Image); err != nil {
				log.Printf("Error scanning rowsGoods: %v", err)
				http.Error(w, "Error scanning rowsGoods", http.StatusInternalServerError)
				return
			}

			goods.DateBuy = &cart.Date
			log.Println("goods ", goods)

			querySellers := "SELECT name FROM users WHERE id = $1"
			argSellers := id
			rowsSellers, err := db.DB.Query(querySellers, argSellers)
			if err != nil {
				log.Printf("Error querying database: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if rowsSellers.Next() {
				if err := rowsSellers.Scan(&goods.Seller); err != nil {
					log.Printf("Error scanning row: %v", err)
					http.Error(w, "Error scanning row", http.StatusInternalServerError)
					return
				}
			}

		}
		if !goods.IsBuy {
			goodsList = append(goodsList, goods)
		}

	}

	log.Println(goodsList)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goodsList); err != nil {
		log.Printf("Error encoding response to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
