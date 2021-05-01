package main

import (
	"github.com/empathy117/ship-of-hope/common"
	"github.com/empathy117/ship-of-hope/controller"
	"github.com/empathy117/ship-of-hope/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	err := common.DB.AutoMigrate(&model.User{})
	if err != nil {
		return 
	}
	err = common.DB.AutoMigrate(&model.Player{})
	if err != nil {
		return 
	}
	controller.InitPlayer()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}