package model

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;unique"`
	Telephone string `gorm:"varchar(110;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Goal      int
	Rock      int
	Paper     int
	Scissor   int
	IsPlaying bool
}
