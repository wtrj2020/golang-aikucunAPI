package controller

import (
	"appApi/models"
	"appApi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//
//flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","s_addressId":2,"orderSrc":"ios","payType":1,"payCode":"wallets"}
//flutter: post:请求url:http://c.judus.top:9999//app/orders/submit
//flutter: -------end------------------------------------------
//flutter: {"status":1,"domain":"http:\/\/c.judus.top:9999\/","msg":"\u63d0\u4ea4\u8ba2\u5355\u6210\u529f","data":"158374024545129852"}
//flutter:
//------post----------------------------------
//flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","payPwd":"123123","orderNo":"158374024545129852","isBatch":1}
//flutter: post:请求url :http://c.judus.top:9999//app/wallets/payByWallet
//					http://c.judus.top:9999//app/payments/weixinPay

func Payments(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo models.Payments
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	jsonInfo.UserId = userId
	i, s := service.Payments(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   nil,
	})
	return
}
