package controller

import (
	"github.com/empathy117/ship-of-hope/common"
	"github.com/empathy117/ship-of-hope/dto"
	"github.com/empathy117/ship-of-hope/model"
	"github.com/empathy117/ship-of-hope/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)
var DB = common.InitDB()

func Register(ctx *gin.Context) {

	// get data
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// verify data's legitimacy
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不得少于6位")
		return
	}
	if len(name) < 5 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不得少于5位")
		return
	}

	//check data's existence (telephone && name)
	if !isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已被注册")
		return
	}
	if !isNameExist(DB, name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已被注册")
		return
	}

	// create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
		Goal:      0,
		Paper:     5,
		Rock:      5,
		Scissor:   5,
		IsPlaying: false,
	}
	DB.Create(&newUser)

	response.Success(ctx, nil, "注册成功")
}

// Login for user to login
func Login(ctx *gin.Context) {

	// get data
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	// verify data
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不得少于6位")
		return
	}
	if len(name) < 5 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不得少于5位")
		return
	}

	// checkout name exist
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// checkout password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) ; err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// get token
	token, err := common.GetToken(user)
	if err != nil {
		msgReleaseTokenError := "系统错误：1"
		response.Response(ctx, http.StatusInternalServerError, 500, nil, msgReleaseTokenError)
		log.Printf("token generate error: %v", err)
		return
	}

	// return result

	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}

func isNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}

func Info(ctx *gin.Context)  {
	user,_ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}