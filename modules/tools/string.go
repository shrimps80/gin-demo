package tools

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 字符串转int64
func String2Int64(val string) int64 {
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

// md5加密
func Md5Str(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
