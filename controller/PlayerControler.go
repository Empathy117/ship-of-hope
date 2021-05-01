package controller

import (
	"github.com/empathy117/ship-of-hope/common"
	"github.com/empathy117/ship-of-hope/model"
)

func InitPlayer() {
	DB := common.GetDB()
	DB.Create(&model.Player{
		Id: 1,
		Playing: 0,
	})
}

// this can only creat player once
// because the Id was unique