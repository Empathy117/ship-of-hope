package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/empathy117/ship-of-hope/model"
	"time"
)

// string
var key = []byte("empathy117")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func GetToken(user model.User) (string, error) {
	// expire time
	expireTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			// expire time
			ExpiresAt: expireTime.Unix(),
			// release time
			IssuedAt: time.Now().Unix(),
			// issuer
			Issuer: "empathy",
			// subject
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return res, nil
}

// ParseToken parse the token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error)  {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})
	return token, claims, err
}