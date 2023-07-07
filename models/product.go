package models

type Product struct {
	Id          int64  `gorm:"primarykey" json:"id"`
	NameProduct string `gorm:"type:varchar(300)" json:"name_product"`
	Description string `gorm:"type:text" json:"description"`
}
