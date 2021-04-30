package controller

import (
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {

	// get data
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// verify data's legitimacy
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不得少于6位"})
		return
	}
	if len(name) < 5 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不得少于5位"})
		return
	}

	log.Println(name, telephone, password)

	// check data's existence (telephone && name)
	if !isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已被注册"})
		return
	}
	if !isNameExist(db, name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	// create user
	newUser := User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
		Goal:      0,
		Paper:     5,
		Rock:      5,
		Scissor:   5,
		IsPlaying: false,
	}
	db.Create(&newUser)

	// return result
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
})