package models

type Personalidade struct {
	Id			int `gorm:"primaryKey"`
	Nome 		string `json:"nome" gorm:"not null"`
	Historia 	string `json:"historia" validate:"required"`
}	

var Personalidades []Personalidade