package controller

import (
	"appApi/models"
	"appApi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","s_addressId":2,"orderSrc":"ios","payType":1,"payCode":"wallets"}
//flutter: post:请求url:http://c.judus.top:9999//app/orders/submit
//flutter: -------end------------------------------------------
//flutter: {"status":1,"domain":"http:\/\/c.judus.top:9999\/","msg":"\u63d0\u4ea4\u8ba2\u5355\u6210\u529f","data":"158362106415066281"}
//flutter:
//------post----------------------------------
//flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","payPwd":"123123","orderNo":"158362106415066281","isBatch":1}
//flutter: post:请求url:http://c.judus.top:9999//app/wallets/payByWallet
//flutter: -------end------------------------------------------
//flutter: {"status":1,"domain":"http:\/\/c.judus.top:9999\/","msg":"\u8ba2\u5355\u652f\u4ed8\u6210\u529f"}
//flutter:
//

//[ sql ] [ SQL ] INSERT INTO `wst_orders` (`userName` , `userPhone` , `areaIdPath` , `areaId` , `userAddress` , `orderNo` , `userId` , `shopId` , `payType` , `goodsMoney` , `deliverType` , `deliverMoney` , `totalMoney` , `scoreMoney` , `useScore` , `realTotalMoney` , `needPay` , `orderStatus` , `isPay` , `orderScore` , `isInvoice` , `invoiceJson` , `invoiceClient` , `orderRemarks` , `orderunique` , `orderSrc` , `dataFlag` , `payRand` , `createTime`) VALUES ('fefwfwefew' , '13818908888' , '1400004_140200_' , 140213 , '山西省大同市平城区123123121' , '33' , 302 , 28 , 1 , 22 , 0 , 0 , 22 , 0 , 0 , 22 , 22 , -2 , 0 , 22 , 0 , '' , '' , NULL , '158362085571004658' , 4 , 1 , 1 , '2020-03-08 06:40:55') [ RunTime:0.000597s ]

func Submit(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo models.Orders
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

	i, s, ordersId := service.Submit(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   ordersId,
	})
}

func ConfirmOrder(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo models.Orders
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

	confirmOrder := service.ConfirmOrder(jsonInfo.Orderunique, jsonInfo.UserId)
	c.JSON(http.StatusOK, models.Result{
		Status: 1,
		Msg:    "看吧",
		Data:   confirmOrder,
	})
}
func GetUserOrders(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	orderStatus := c.Query("type")
	fmt.Println(orderStatus)
	//c.JSON(http.StatusOK, models.Result{
	//	Status: i,
	//	Msg:    s,
	//	Data:   ordersId,
	//})

	orders := service.GetUserOrders(userId, orderStatus)
	c.JSON(http.StatusOK, models.Result{
		Status: 1,
		Msg:    "获取成功",
		Data:   orders,
	})
}

//确认收货
func Receive(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo models.Orders
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("登录使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	jsonInfo.UserId = userId
	s, i := service.Receive(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
	})
}
