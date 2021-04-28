package middleware

import (
	"appApi/dao"
	"appApi/models"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func Usercheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		var data interface{}
		code = 1

		var res models.Users
		var tokenId = c.GetHeader("token")

		if tokenId == "" {
			code = -1
			msg = "缺少token"
		} else {
			//	res = dao.GetUserId(tokenId)
			_redisUser, err := dao.Redisdb.Get(tokenId).Result()
			if err != nil {
				log.Println("redis失败，改走数据库验证token")
				res = dao.GetUserId(tokenId)
			} else {
				json.Unmarshal([]byte(_redisUser), &res)
			}

			if res.UserId == 0 {
				code = -1
				msg = "无效token"
			} else if res.Status != 1 || res.DataFlag != 1 {
				code = -1
				msg = "无效用户"
			} else {
				//取用户id
				c.Set("userId", res.UserId)
			}
		}
		if code != 1 {
			if msg == "" {
				msg = "1212"
			}
			c.JSON(code, models.Result{
				Status: code,
				Msg:    msg,
				Data:   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
