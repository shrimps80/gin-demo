package services

import (
	"fmt"
	"time"
	
	"github.com/dgrijalva/jwt-go"
	"gin-demo/models/Users"
)

const (
	KEY                    string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int    = 600 // default 10 minutes
)

type MyCustomClaims struct {
	*Users.User
	jwt.StandardClaims
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	newClaims := MyCustomClaims{
		claims.User,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.User.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate new fresh json web token failed !! error :", err)
		return "", err
	}
	return tokenStr, err
}

func GenerateToken(user *Users.User, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("生成令牌失败 !! error :%v", err)
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string) (*Users.User, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.User, claims.StandardClaims.ExpiresAt)
		fmt.Println("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
		return claims.User, nil
	} else {
		fmt.Println("validate tokenString failed !!!", err)
		return nil, err
	}
	return nil, nil
}
