package service

import (
	"appApi/dao"
	"appApi/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

//var (
//	aitPay       int = -2
//	waitDelivery int = 0
//	waitReceive  int = 1
//	waitAppraise int = 2
//	finish       int = 2
//	abnormal     int = -3
//)

func Submit(jsonInfo models.Orders) (int, string, models.Orders) {
	var allShopCartList []models.ShopsOrder
	var orderRes models.Orders //空的
	//获取支付方式
	payments := dao.Payments(jsonInfo.PayFrom)
	if payments.Id == 0 {
		return -1, "下单失败，请选择有效支付方式", orderRes
	}

	//取购物车数据(后面要检查购物车数据是否合法)
	getCarts := dao.GetCarts(jsonInfo.UserId)
	if len(getCarts) == 0 {
		return -1, "下单失败，请选择有效的库存商品", orderRes
	}
	//取下单地址
	getUserAddRess := dao.GetUserAddRess(jsonInfo.UserId, jsonInfo.AddressId)
	if getUserAddRess.AddressId == 0 {
		return -1, "无效的用户地址", orderRes
	}
	//生成支付编号
	jsonInfo.Orderunique = strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(getUserAddRess.UserId)
	//把购物车产品到写到goods——order便于计算

	//整合购物车信息
	allShopCartList = dao.AllShopCartList(jsonInfo.UserId, 1)
	if len(allShopCartList) == 0 {
		return 1, "订单创建失败", orderRes
	}

	//创建订单
	dao.CreOrder(jsonInfo, getUserAddRess)
	return 1, "订单创建成功", ConfirmOrder(jsonInfo.Orderunique, jsonInfo.UserId)
}

//确认订单
func ConfirmOrder(orderUnique string, userId int) (orderRes models.Orders) {
	var order []models.Orders
	dao.DB.Find(&order, "orderunique=? and user_id=? and order_status=-2 and data_flag=1", orderUnique, userId)
	var payments []models.Payments
	dao.DB.Find(&payments, "enabled=?", 1)
	for _, v := range order {
		orderRes.GoodsMoney = orderRes.GoodsMoney + v.GoodsMoney
		orderRes.Orderunique = orderUnique
		orderRes.PayFrom = v.PayFrom
		orderRes.PayType = v.PayType
		orderRes.OrderStatus = v.OrderStatus
		orderRes.Payments = payments
	}
	return
}

func GetUserOrders(userId int, orderStatus string) (order []models.OrdersList) {
	switch orderStatus {
	case "waitPay":
		fmt.Println(orderStatus)
		return dao.GetUserOrders(-2, userId)
	case "waitDelivery":
		return dao.GetUserOrders(0, userId)
	case "waitReceive":
		return dao.GetUserOrders(1, userId)
	case "waitAppraise":
		return dao.GetUserOrders(2, userId)
	case "finish":
		return dao.GetUserOrders(2, userId)
	case "abnormal":
		return dao.GetUserOrders(-3, userId)
	default:
		fmt.Println("ss")
		return
	}
	return
}

//确认收货
func Receive(jsonInfo models.Orders) (string, int) {
	orders := dao.Receive(jsonInfo)
	if orders.OrderId == 0 {
		return "操作失败，请检查订单状态是否已改变", -1
	}

	err := dao.DB.Model(&orders).UpdateColumns(map[string]interface{}{"order_status": 2, "receive_time": time.Now().Unix()}).Error
	if err != nil {
		return "操作失败~", -1
	}
	//验证是否Update成功
	if orders.OrderStatus != 2 {
		return "确认收货失败~", -1
	}

	//给商家结算
	var shops models.Shops
	//dao.DB.First(&shops)
	shops.ShopId = orders.ShopId
	dao.DB.Model(&shops).Update("shop_money", gorm.Expr("shop_money+?", orders.CleanPrice))
	//新增订单日志
	log_orders := models.Log_orders{
		OrderId:     orders.OrderId,
		OrderStatus: 2,
		LogContent:  "订单编号：[" + orders.OrderNo + "]确认收货",
		LogUserId:   jsonInfo.UserId,
		LogType:     0,
		CreateTime:  time.Now().Unix(),
	}
	dao.DB.Create(&log_orders)

	//var shops models.Shops
	//dao.DB.First(&shops, "shop_id=?", orders.ShopId)
	log_moneys := models.Log_moneys{
		TargetType: 1,
		TargetId:   orders.ShopId,
		DataId:     orders.OrderId,
		DataSrc:    "", //???
		Remark:     "订单编号：[" + orders.OrderNo + "]确认收货结算" + fmt.Sprintln(orders.CleanPrice),
		MoneyType:  1,
		Money:      orders.CleanPrice,
		PayType:    orders.PayFrom,
		DataFlag:   1,
		CreateTime: time.Now().Unix(),
	}
	dao.DB.Create(&log_moneys)

	return "操作成功~", 1

}
