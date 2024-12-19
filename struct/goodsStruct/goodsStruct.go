package goodsstruct

type Goods struct {
	ID          int     `json:"id"`
	IDS         int     `json:"id_s"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DatePub     string  `json:"date_pub"`
	DateBuy     *string `json:"date_buy"`
	IsBuy       bool    `json:"is_buy"`
	Image       string  `json:"image"`
}

type GoodsToFront struct {
	ID          int     `json:"id"`
	Seller      string  `json:"seller"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DatePub     string  `json:"date_pub"`
	DateBuy     *string `json:"date_buy"`
	IsBuy       bool    `json:"is_buy"`
	Image       string  `json:"image"`
}
