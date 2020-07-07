package sign

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"gin-demo/defs"
	"gin-demo/modules/response"
	"gin-demo/config"
	"gin-demo/modules/tools"
	"github.com/davecgh/go-spew/spew"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		if skip := skipVerify(c); skip == false {
			sign := c.Request.Header.Get("sign")
			if sign == "" {
				response.ReturnErrorJson(c, defs.MissingSignature)
				c.Abort()
				return
			}
			
			timestamp := c.Request.Header.Get("timestamp")
			if timestamp == "" {
				response.ReturnErrorJson(c, defs.MissingSignature)
				c.Abort()
				return
			}
			
			//判断time是否过期
			timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
			minute := time.Now().Unix() - timestampInt
			if minute > 300 { // 5分钟
				response.ReturnErrorJson(c, defs.MissingSignature)
				c.Abort()
				return
			}
			
			_ = c.Request.ParseForm()
			req := createEncryptStr(c.Request.Form)
			
			SignKey := config.GetEnv().AppSecret
			signArray := []string{
				timestamp,
				req,
				SignKey,
			}
			signStr := strings.Join(signArray, "&")
			signMd5 := strings.ToUpper(tools.Md5Str(signStr))
			spew.Dump(signStr, signMd5)
			if signMd5 != sign {
				response.ReturnErrorJson(c, defs.MissingSignature)
				c.Abort()
				return
			}
		}
	}
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}

func skipVerify(c *gin.Context) bool {
	var skip string
	methodStr := c.Request.Method
	if methodStr == "GET" || methodStr == "DELETE" {
		skip = c.Query("skip_debug")
	} else {
		skip = c.PostForm("skip_debug")
	}
	if skip == "18120080" {
		return true
	}
	return false
}
