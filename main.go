package main

func main() {
	db := InitDB()
	defer db.Debug()
	r := gin.Default()
	r.Post("/api/autu/register", controller.UserRegister)
	r.Run()
}