package User

import (
	"github.com/gookit/validate"
)

type UserForm struct {
	Name string `form:"name" json:"name" validate:"required|CustomValidator"`
}

func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"Name.required":        "昵称不能为空",
		"Name.CustomValidator": "name改成abc你试试",
	}
}

// 自定义验证
func (f UserForm) CustomValidator(val string) bool {
	return val == "abc"
}
