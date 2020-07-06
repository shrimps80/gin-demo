package services

import (
	"time"
	
	"gin-demo/models/Users"
	"gin-demo/modules/tools"
	"gin-demo/modules/database/redis"
	"gin-demo/defs"
)

const (
	LOGIN_EXPIRED = 86400
)

func UserLoginByEmail(email, password string) (rul *defs.ResUserLogin, err error) {
	user, err := Users.GetOneByEmail(email)
	if err != nil {
		return rul, err
	}
	//TODO 验证密码
	
	token, err := GenerateToken(user, LOGIN_EXPIRED)
	if err != nil {
		return rul, err
	}
	
	// 保存到redis
	redisKey := tools.UserTokenKey(user.Id)
	r := redis.Client
	if err := r.Set(redisKey, token, time.Duration(LOGIN_EXPIRED)*time.Second); err != nil {
		return rul, err
	}
	
	rul = &defs.ResUserLogin{
		Token:   token,
		Expired: LOGIN_EXPIRED,
		UserId:  user.Id,
	}
	return
}
