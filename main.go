package main

import (
	"appApi/controller"
	"appApi/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.New()
	r.Use(middleware.Cors())
	v1 := r.Group("/v1")
	{
		///////////////////////
		//获取档口列表
		v1.GET("/ShopList", controller.GoodsCats)

		///////////////////////
		//产品分类
		v1.GET("/GoodsCats", controller.GoodsCats)

		//产品分类下的宝贝
		v1.GET("/GoodsCatsGoods", controller.GoodsCatsGoods)

		//宝贝详情
		v1.GET("/GoodsDetails", controller.GoodsDetails)
		//产品列表  ???
		//v1.POST("/GetGoodsList", controller.GetGoodsList)
		//注册用户
		v1.POST("/RegUser", controller.RegUser)
		//登录
		v1.POST("/Login", controller.Login)
		//个人中心
		v1.GET("/Member", middleware.Usercheck(), controller.Member)

		//加入购物车
		v1.POST("/AddCart", middleware.Usercheck(), controller.AddCart)
		//删除购物车
		v1.POST("/DelOnlyCart", middleware.Usercheck(), controller.DelOnlyCart)
		//用户地址
		v1.GET("/MyAddressList", middleware.Usercheck(), controller.MyAddressList)
		//添加地址
		v1.POST("/AddMyAddress", middleware.Usercheck(), controller.AddMyAddress)
		//修改默认地址
		v1.POST("/EditIsdefaultAddress", middleware.Usercheck(), controller.EditIsdefaultAddress)
		//----------------
		//我de购物车
		v1.GET("/MyCartList", middleware.Usercheck(), controller.MyCartList)
		//我的订单列表
		v1.GET("/GetUserOrders", middleware.Usercheck(), controller.GetUserOrders)
		//创建订单
		v1.POST("/Submit", middleware.Usercheck(), controller.Submit)
		//确认订单
		v1.POST("/ConfirmOrder", middleware.Usercheck(), controller.ConfirmOrder)
		//订单支付
		v1.POST("/Payments", middleware.Usercheck(), controller.Payments)
		//确认收货
		v1.POST("/Receive", middleware.Usercheck(), controller.Receive)

		//分类下- 活动中-即将活动
		v1.POST("/ActList", controller.ActList)
		/////////////////
		v1.POST("/ActionsList", controller.ActionsList)
		//指定活动编号的商品
		v1.POST("/ActionsGoodsList", controller.ActionsGoodsList)
	}

	//上传文件
	r.POST("/upload", middleware.Usercheck(), controller.Upload)
	//文件目录
	r.StaticFS("/public/upload", http.Dir("public/upload"))

	r.Run(":7000")

	//r.POST("/RegUser", controller.RegUser)
}
