package model

type Player struct {
	Id				uint `gorm:"unique"`
	Playing 	uint
	Card  		string `gorm:"size:255"`
}
