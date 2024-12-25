package goodsstruct

import userstruct "back/struct/userStruct"

type Good struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement"`
	IDS         int             `json:"id_s" gorm:"not null"`
	Title       string          `json:"title" gorm:"type:varchar(100)"`
	Description string          `json:"description" gorm:"type:text"`
	Price       float64         `json:"price" gorm:"type:numeric"`
	DatePub     string          `json:"date_pub" gorm:"type:date"`
	DateBuy     *string         `json:"date_buy" gorm:"type:date"`
	IsBuy       bool            `json:"is_buy" gorm:"type:boolean"`
	IsSell      bool            `json:"is_sell" gorm:"type:boolean"`
	Image       string          `json:"image" gorm:"type:varchar(255)"`
	User        userstruct.User `gorm:"foreignKey:IDS;references:ID"`
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
