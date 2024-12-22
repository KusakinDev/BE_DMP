package userstruct

type User struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"type:varchar(20)"`
	Password    string  `json:"password" gorm:"type:varchar(20)"`
	Email       string  `json:"email" gorm:"type:varchar(20);unique" column:"email"`
	Rating      float32 `json:"rating" gorm:"type:real"`
	CountRating int     `json:"count_rating" gorm:"type:integer"`
	DebugWallet float64 `json:"debug_wallet" gorm:"type:numeric"`
}
