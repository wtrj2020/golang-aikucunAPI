package service

import (
	"appApi/dao"
	"appApi/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func Payments(jsonInfo models.Payments) (int, string) {
	//DB.Find(&payments, "pay_from=?", pay_from)
	payments := dao.Payments(jsonInfo.PayFrom)
	if payments.Id == 0 || payments.Enabled != 1 {
		//if payments.Id == 0 || payments.Enabled != 1 {
		return -1, "无效支付方式"
	}
	orders := dao.GetOrders(jsonInfo.UserId, jsonInfo.OrderNo)
	if len(orders) == 0 {
		return -1, "无效订单编号"
	}
	if orders[0].OrderStatus != -2 {
		return -1, "订单已经支付过了吧"
	}
	//支付金额
	var total_money float64
	for _, v := range orders {
		total_money = total_money + v.TotalMoney
	}
	switch payments.PayFrom {
	case "wallets":
		return wallets(total_money, orders, jsonInfo)
	case "wxpay":
		return wxpay()
	case "alipay":
		return alipay()
	}

	return -1, "支付失败"
}

//total_money, orders, jsonInfo
func wallets(total_money float64, orders []models.Orders, jsonInfo models.Payments) (int, string) {
	user := dao.GetUserIdX(jsonInfo.UserId)
	if user.UserMoney < total_money {
		return -1, "余额不足"
	}

	//减钱
	err := dao.DB.Model(&models.Users{UserId: jsonInfo.UserId}).UpdateColumn("user_money", gorm.Expr("user_money - ?", total_money)).Error
	if err != nil {
		panic(err)
	}

	SetOrderStatus(total_money, orders, jsonInfo)

	return 1, "余额支付成功"
}

func wxpay() (int, string) {
	return 1, "微信支付成功"
}

func alipay() (int, string) {
	return 1, "支付宝支付成功"
}

//设置支付成功订单状态,给店主发消息
func SetOrderStatus(total_money float64, orders []models.Orders, jsonInfo models.Payments) {

	for _, v := range orders {
		dao.DB.Model(&models.Orders{ShopId: v.ShopId, UserId: jsonInfo.UserId, OrderId: v.OrderId, Orderunique: v.Orderunique}).Updates(map[string]interface{}{"order_status": 0, "is_pay": 1, "pay_from": jsonInfo.PayFrom, "pay_time": time.Now().Unix()})
		fmt.Println("发消息")
		//记录订单消息
		log_orders := models.Log_orders{
			//LogId:       0,
			OrderId:     v.OrderId,
			OrderStatus: 0,
			LogContent:  "订单已支付,下单成功",
			LogUserId:   jsonInfo.UserId,
			LogType:     0,
			//LogTime:     "",
			CreateTime: time.Now().Unix(),
		}
		dao.DB.Create(&log_orders)

		log_moneys := models.Log_moneys{
			TargetType: 0,
			TargetId:   v.UserId,
			DataId:     v.OrderId,
			DataSrc:    "", //???
			Remark:     "订单编号：[" + v.OrderNo + "]支出" + fmt.Sprintln(v.GoodsMoney),
			MoneyType:  1,
			Money:      v.GoodsMoney,
			PayType:    v.PayFrom,
			DataFlag:   1,
			CreateTime: time.Now().Unix(),
		}
		dao.DB.Create(&log_moneys)
	}
}
