package dao

import (
	"appApi/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/thedevsaddam/gojsonq"
	"math/rand"
	"strconv"
	"time"
)

//创建订单
func CreOrder(jsonInfo models.Orders, address models.User_address) (count int) {
	//整合购物车信息，isCheck1 代表以勾选
	allShopCartList := AllShopCartList(jsonInfo.UserId, 1)

	//var res models.Orders
	for _, v := range allShopCartList {
		//子订单编号=时间戳+用户id+店铺id+用户地址 确保不会重复
		rand.Seed(time.Now().UnixNano())
		OrderNo := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(jsonInfo.UserId) + strconv.Itoa(v.ShopId) + strconv.Itoa(address.AddressId) + strconv.Itoa(rand.Intn(9999))
		OrderGoods := models.Orders{
			//OrderId:         0,
			OrderNo:      OrderNo,
			ShopId:       v.ShopId,
			UserId:       jsonInfo.UserId,
			OrderStatus:  -2,               //待付款
			PayFrom:      jsonInfo.PayFrom, //支付方式
			GoodsMoney:   v.ShopPrice,      //总额
			TotalMoney:   v.ShopPrice,      //总额(以后积分产于付款 预留这里吧
			CleanPrice:   v.ShopCleanPrice, //总额 商家结算用
			PayType:      1,
			IsPay:        0,              //初始化是否支付
			AreaId:       address.AreaId, //城市编号
			AreaIdPath:   "",             //收货完整省市区县到编号 _分割
			UserName:     address.UserName,
			UserAddress:  address.UserAddress,
			UserPhone:    address.UserPhone,
			OrderScore:   0,
			OrderRemarks: jsonInfo.OrderRemarks, //订单备注
			OrderSrc:     jsonInfo.OrderSrc,     //订单设备
			Orderunique:  jsonInfo.Orderunique,  //总订单编号
			DataFlag:     1,
			CreateTime:   time.Now().Unix(),
			//PayTime:      time.Date(2019, 10, 28, 23, 20, 31, 0, time.Local),
		}
		err := DB.Create(&OrderGoods).Error
		if err != nil {
			panic(err)
		}

		//创建订单详细列表
		if OrderGoods.OrderId > 0 {
			for _, vv := range v.List {
				orderGoodsList := models.Order_goods{
					//Id:             0,
					OrderId:        OrderGoods.OrderId,
					GoodsId:        vv.GoodsId,
					GoodsNum:       vv.CartNum,
					GoodsPrice:     vv.GoodsPrice,
					CleanPrice:     vv.CleanPrice,
					GoodsSpecId:    vv.GoodsSpecId,
					GoodsSpecNames: vv.SpecName,
					GoodsName:      vv.GoodsName,
					GoodsImg:       vv.GoodsImg,
					GoodsSn:        vv.GoodsSn,
					Commission:     vv.Commission,
				}
				DB.Create(&orderGoodsList)
				//修改规格库存和总库存
				if orderGoodsList.GoodsSpecId > 0 {
					DB.Model(&models.Goods_specs{Id: orderGoodsList.GoodsSpecId}).Update("goods_stock", gorm.Expr("goods_stock - ?", vv.CartNum))
					sumGoodsStock := SumGoodsStock(orderGoodsList.GoodsId)
					DB.Model(&models.Goods{GoodsId: orderGoodsList.GoodsId}).Update("goods_stock", sumGoodsStock)

				}
			}
		}
		//记录订单消息
		Log_orders := models.Log_orders{
			//LogId:       0,
			OrderId:     OrderGoods.OrderId,
			OrderStatus: OrderGoods.OrderStatus,
			LogContent:  "下单成功，等待用户支付",
			LogUserId:   OrderGoods.UserId,
			LogType:     0,
			//LogTime:     "",
			CreateTime: time.Now().Unix(),
		}
		DB.Create(&Log_orders)

		count++
	}
	//删除已选的购物车商品
	DB.Where("is_check=1 and user_id=? and is_check=1", jsonInfo.UserId).Delete(make([]models.Carts, 0))

	return
}

//整合购物车信息
func AllShopCartList(userId, isCheck int) []models.ShopsOrder {
	//购物车数据合并到一起

	getCartList := GetCartList(userId)
	shopList := make([]models.ShopsOrder, 0, len(getCartList))
	shopIds := make(map[int]int) //将key作为shopid起到去重作用
	var address models.User_address
	//取用户默认收货地址
	DB.First(&address, "user_id=? and is_default=1 and data_flag=1", userId)
	////算出每个店铺订单总金额
	for _, v := range getCartList {
		if shopIds[v.ShopId] == 0 {
			shopIds[v.ShopId] = v.ShopId
			shops, goodsList := cartgoodsList(v.ShopId, getCartList, isCheck)
			if len(goodsList) <= 0 {
				continue
			}
			shopList = append(shopList, models.ShopsOrder{
				ShopId:         v.ShopId,
				ShopName:       v.ShopName,
				ShopImg:        v.ShopImg,
				ShopAddress:    v.ShopAddress,
				ShopPrice:      shops.ShopPrice,
				ShopCleanPrice: shops.ShopCleanPrice, //商家结算用
				CartNum:        shops.CartNum,
				Address:        address, //暂时放在这里吧，虽然难看但是方便。。。。
				List:           goodsList,
			})
		}
	}

	return shopList
}

//取购物中当前商铺的订单
func cartgoodsList(shopId int, getCarts []models.GetCarts, isCheck int) (shops models.ShopsOrder, goodsList []models.GetCarts) {
	goodsList = make([]models.GetCarts, 0, len(getCarts))
	for _, v := range getCarts {
		if v.ShopId != shopId {
			continue
		}
		if v.ShopId == 0 {
			continue
		}

		//是否显示未勾选的 1为勾选，（用户购物车和结算做区分）
		if isCheck == 1 {
			if v.IsCheck != 1 {
				continue
			}
		}
		if v.GoodsStock <= 0 {
			continue
		}
		if v.CartNum > v.GoodsStock {
			v.CartNum = v.GoodsStock
		}
		shops.ShopPrice = shops.ShopPrice + v.GoodsPrice*float64(v.CartNum)           //实际金额
		shops.ShopCleanPrice = shops.ShopCleanPrice + v.CleanPrice*float64(v.CartNum) //商家结算记录
		shops.CartNum = shops.CartNum + v.CartNum
		goodsList = append(goodsList, models.GetCarts{
			ShopId:      v.ShopId,
			CartId:      v.CartId,
			GoodsName:   v.GoodsName,
			GoodsImg:    v.GoodsImg,
			GoodsPrice:  v.GoodsPrice,
			GoodsSpecId: v.GoodsSpecId,
			CleanPrice:  v.CleanPrice,
			CartNum:     v.CartNum,
			GoodsStock:  v.GoodsStock,
			GoodsSn:     v.GoodsSn,
			SpecName:    v.SpecName,
			SpecColour:  v.SpecColour,
			SpecOnly:    v.SpecOnly,
			GoodsId:     v.GoodsId,
			IsCheck:     v.IsCheck,
			Commission:  v.Commission,
		})

	}
	return
}

//购物车数据合并到一起
func GetCartList(userId int) (GetCarts []models.GetCarts) {
	fmt.Println(userId)
	//DB.Table("fc_carts").Where("fc_carts.user_id=? and fc_carts.is_check=1", userId).
	DB.Table("fc_carts").Where("fc_carts.user_id=?", userId).
		Select(`
				fc_shops.shop_id,
				fc_carts.cart_id,
				fc_carts.goods_spec_id,
				fc_carts.is_check,
				fc_carts.cart_num,
				fc_goods.goods_id,
				fc_goods.goods_name,
				fc_goods.goods_img,
				fc_goods_specs.spec_name,
				fc_goods_specs.goods_sn,
				fc_goods_specs.spec_colour,
				fc_goods_specs.spec_only,
				fc_goods_specs.goods_stock,
				fc_goods_specs.goods_price,
				fc_goods_specs.clean_price,
				fc_goods_specs.commission,
				fc_goods_specs.goods_img,
				fc_shops.shop_name,
				fc_shops.shop_img,
				fc_shops.shop_address
			`).
		Joins("left join fc_goods on  fc_carts.goods_id=fc_goods.goods_id").
		Joins("left join fc_shops on  fc_shops.shop_id=fc_goods.shop_id").
		Joins("left join fc_goods_specs on  fc_carts.goods_spec_id=fc_goods_specs.id").
		Scan(&GetCarts)
	return
}

//取总订单列表
func GetOrders(userId int, orderunique string) (order []models.Orders) {
	DB.Find(&order, "data_flag=1 and user_id=? and orderunique=?", userId, orderunique)
	return order
}

//统计剩余库存量
func SumGoodsStock(goodsId int) int {
	var goodsSpecs models.Goods_specs
	DB.Select("SUM(goods_stock) as goods_stock").Find(&goodsSpecs, "goods_id=? and data_flag=1", goodsId)
	return goodsSpecs.GoodsStock
}

//用户或店铺取购物车数据
func GetUserOrders(orderStatus, userId int) []models.OrdersList {

	var order []models.Orders
	DB.Order("order_id desc").Find(&order, "data_flag=1 and order_status=? and user_id=?", orderStatus, userId)
	orderJson, _ := json.Marshal(order)

	//取店铺信息
	var shops []models.Shops
	shops = GetAllShops()
	shopsJson, _ := json.Marshal(shops)
	//取支付信息
	var payments []models.Payments
	DB.Find(&payments, "enabled=?", 1)

	orderUniques := make(map[int]string) //将key作为shopid起到去重作用
	var goodsids = []int{}
	//总订单编号去重
	for i, v := range order {
		orderUniques[i] = v.Orderunique
		goodsids = append(goodsids, v.OrderId)
	}
	var orderGoods []models.Order_goods
	DB.Find(&orderGoods, "order_id IN (?)", goodsids)
	orderGoodsv, _ := json.Marshal(orderGoods)

	var OrderList []models.OrdersList
	fmt.Println(orderUniques)

	//var i int
	//var temp = []int{}
	temp := make(map[string]bool) //将key作为shopid起到去重作用

	for i := 0; i < len(orderUniques); i++ {

		if temp[orderUniques[i]] == true {
			continue
		} else {
			temp[orderUniques[i]] = true
		}
		//temp = append(temp, orderUniques[i])
		//for _, value := range orderUniques //当心不按顺序读取
		var outputResult []models.Orders
		outputResult = getOrderGoJsonq(string(orderJson), orderUniques[i])
		fmt.Println(orderUniques[i])

		var asd models.OrdersList

		shopList := make([]models.ShopsOrder, 0, len(outputResult))

		for _, v := range outputResult {
			var outputResult []models.Orders
			outputResult = getOrderGoJsonq2(string(orderJson), orderUniques[i], v.ShopId)

			OrderGoods := getOrderGoJsonq3(string(orderGoodsv), v.OrderId)
			for k, _ := range outputResult {
				outputResult[k].OrderGoods = OrderGoods
			}
			shops := getShopsGojsonq(string(shopsJson), v.ShopId)

			shopList = append(shopList, models.ShopsOrder{
				ShopId:      v.ShopId,
				ShopName:    shops.ShopName,
				ShopImg:     shops.ShopImg,
				ShopAddress: shops.ShopAddress,
				Address: models.User_address{
					UserName:    v.UserName,
					UserPhone:   v.UserPhone,
					AreaIdPath:  v.AreaIdPath,
					AreaId:      v.AreaId,
					UserAddress: v.UserAddress,
				},
				Orders: outputResult,
			})

			asd.TotalMoney = asd.TotalMoney + v.GoodsMoney
			asd.Orderunique = v.Orderunique
			asd.UserName = v.UserName
			asd.UserPhone = v.UserPhone
			asd.AreaId = v.AreaId
			asd.UserAddress = v.UserAddress
			asd.CreateTime = v.CreateTime
			asd.PayTime = v.PayTime
			asd.OrderStatus = v.OrderStatus
			asd.Payments = payments
		}

		asd.ShopsOrder = shopList

		OrderList = append(OrderList, asd)

		//	}
	}
	//OrderListJson, _ := json.Marshal(OrderList)
	//outputResult := sortOrderList(string(OrderListJson))
	return OrderList
}

//func sortOrderList(OrderListJson string) (outputResult []models.OrdersList) {
//	res := gojsonq.New().SortBy("create_time", "desc").JSONString(string(OrderListJson)).Where("create_time", ">", 0).Get()
//	resBytes, _ := json.Marshal(res)
//	json.Unmarshal(resBytes, &outputResult)
//	fmt.Println(string(resBytes))
//
//	return
//}
func getOrderGoJsonq3(orderJson string, orderId int) (outputResult []models.Order_goods) {
	res := gojsonq.New().Sort("id desc").JSONString(string(orderJson)).Where("order_id", "=", orderId).Get()
	resBytes, _ := json.Marshal(res)
	json.Unmarshal(resBytes, &outputResult)
	return
}
func getShopsGojsonq(shopsJson string, shopId int) models.Shops {
	//.SortBy("shop_id", "desc")
	res := gojsonq.New().JSONString(string(shopsJson)).Where("shop_id", "=", shopId).First()
	resBytes, _ := json.Marshal(res)
	var outputResult models.Shops
	json.Unmarshal(resBytes, &outputResult)
	return outputResult
}
func getOrderGoJsonq(orderJson, value string) (outputResult []models.Orders) {
	//fmt.Println(value)
	res := gojsonq.New().JSONString(string(orderJson)).Where("orderunique", "=", value).Get()
	//res.SortBy("name", "asc")
	resBytes, _ := json.Marshal(res)
	json.Unmarshal(resBytes, &outputResult)
	return
}
func getOrderGoJsonq2(orderJson, value string, shopId int) (outputResult []models.Orders) {

	res := gojsonq.New().JSONString(string(orderJson)).Where("orderunique", "=", value).Where("shop_id", "=", shopId).Get()
	//var outputResult []models.Orders
	resBytes, _ := json.Marshal(res)
	json.Unmarshal(resBytes, &outputResult)
	return
}

//确认收货
func Receive(jsonInfo models.Orders) models.Orders {
	var orders models.Orders
	DB.Find(&orders, "order_id=? and user_id=? and order_status=1 and data_flag=1", jsonInfo.OrderId, jsonInfo.UserId)
	return orders
}
