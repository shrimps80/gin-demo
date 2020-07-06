package tools

import (
	"fmt"
)

//用户token单设备
func UserTokenKey(userId int64) string {
	return fmt.Sprintf("shop:user:token:%d", userId)
}
