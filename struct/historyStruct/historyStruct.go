package historystruct

import (
	goodsstruct "back/struct/goodsStruct"
	itemstruct "back/struct/itemStruct"
	userstruct "back/struct/userStruct"
)

type History struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	IdU  *int   `json:"id_u" gorm:"type:bigint"`
	IdG  *int   `json:"id_g" gorm:"type:bigint"`
	IdI  *int   `json:"id_i" gorm:"type:bigint"`
	Date string `json:"date" gorm:"type:date"`

	User userstruct.User  `gorm:"foreignKey:IdU;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Good goodsstruct.Good `gorm:"foreignKey:IdG;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Item itemstruct.Item  `gorm:"foreignKey:IdI;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
