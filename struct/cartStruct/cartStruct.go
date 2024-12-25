package cartstruct

import (
	goodsstruct "back/struct/goodsStruct"
	userstruct "back/struct/userStruct"
)

type Cart struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"` // Первичный ключ
	Id_u int    `json:"id_u" gorm:"not null"`               // Внешний ключ на User.ID
	Id_p int    `json:"id_p" gorm:"not null"`               // Внешний ключ на Good.ID
	Date string `json:"date" gorm:"type:date"`

	// Ассоциации
	User userstruct.User  `gorm:"foreignKey:Id_u;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Good goodsstruct.Good `gorm:"foreignKey:Id_p;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
