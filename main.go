package main

import (
	"github.com/empathy117/ship-of-hope/common"
	"github.com/empathy117/ship-of-hope/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	controller.InitPlayer()
	defer db.Debug()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}