package models

import "time"

type Transaction struct {
	Id        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"startDate" `
	EndDate   time.Time `json:"endDate" `
	User      User      `json:"user"`
	UserId    int       `json:"userId"`
	Price     int       `json:"price"`
	Status    string    `json:"status" gorm:"type: varchar(255)"`
}
