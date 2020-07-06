package User

import (
	"github.com/gookit/validate"
)

type UserLogin struct {
	Email    string `form:"email" json:"email" validate:"required|email"`
	Password string `form:"password" json:"password" validate:"required|minLen:6"`
}

func (f UserLogin) Messages() map[string]string {
	return validate.MS{
		"required":        "参数不能为空",
		"email":           "邮箱格式不正确",
		"Password.minLen": "密码最少6位",
	}
}
