package models

type Category struct {
	Id   int    `json:"id" gorm:"primary_key:auto_increment; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name string `json:"name" binding:"required,name" gorm:"varchar(255);not ull; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
