package dao

import (
	"appApi/models"
	"fmt"
	"time"
)

func GetCarts(userId int) (res []models.Carts) {
	DB.Find(&res, "user_id=? and is_check=1", userId)
	return
}

func CheckGoodsSaleSpec(goodsId, goodsSpecId int) (int, models.Goods_specs) {
	fmt.Println("CheckGoodsSaleSpec，验证商品是否合法")
	var goods models.Goods
	var goodsSpecs models.Goods_specs

	DB.Find(&goods, "goods_id=? and status=1 and data_flag=1", goodsId)
	if goods.GoodsId == 0 {
		return 0, goodsSpecs
	}
	DB.Find(&goodsSpecs, "id=? and goods_id=? and status=1 and data_flag=1", goodsSpecId, goodsId)
	if goodsSpecs.Id == 0 {
		return 0, goodsSpecs
	}
	//库存为0排除
	if goodsSpecs.GoodsStock == 0 {
		return 0, goodsSpecs
	}
	return 1, goodsSpecs
}

func IsExistCart(goodsId, goodsSpecId int) models.Carts {
	var cart models.Carts
	DB.Find(&cart, "goods_id=? and goods_spec_id=?", goodsId, goodsSpecId)
	return cart
}
func AddCart(userId int, jsonInfo []models.Carts) int {

	var num = 0 //记录成功订单
	if len(jsonInfo) == 0 {
		return 0
	}
	for _, value := range jsonInfo {
		if value.CartNum <= 0 || value.GoodsSpecId <= 0 || value.GoodsId <= 0 {
			continue
		}
		//验证商品是否合法
		i, goodsSpecs := CheckGoodsSaleSpec(value.GoodsId, value.GoodsSpecId)
		if i == 0 {
			continue
		}
		//如果购买量大于库存，能买多少卖多少
		if value.CartNum > goodsSpecs.GoodsStock {
			value.CartNum = goodsSpecs.GoodsStock
		}
		//购物车里如果存在则增加购买数量,否则新增
		var cart = IsExistCart(value.GoodsId, value.GoodsSpecId)
		if cart.CartId > 0 {
			//根据之前数量加上现在购买数量 如果大于库存以库存为准
			if cart.CartNum+value.CartNum > goodsSpecs.GoodsStock {
				DB.Model(&cart).Updates(map[string]interface{}{"cart_num": goodsSpecs.GoodsStock, "update_time": time.Now().Unix()})
				//DB.Model(&cart).UpdateColumn("cart_num", goodsSpecs.GoodsStock)
				num++
			} else {
				DB.Model(&cart).Updates(map[string]interface{}{"cart_num": cart.CartNum + value.CartNum, "update_time": time.Now().Unix()})
				//DB.Model(&cart).UpdateColumn("cart_num", cart.CartNum+value.CartNum)
				num++
			}

		} else {
			res := models.Carts{
				//CartId:      0,
				UserId:      userId,
				GoodsId:     value.GoodsId,
				ShopId:      goodsSpecs.ShopId,
				GoodsSpecId: value.GoodsSpecId,
				CartNum:     value.CartNum,
				IsCheck:     value.IsCheck,
				CreateTime:  time.Now().Unix(),
				UpdateTime:  time.Now().Unix(),
			}
			DB.Create(&res)
			if res.CartId > 0 {
				num++
			}
		}

	}
	return num
}

func DelOnlyCart(userId int, jsonInfo models.Carts) int {
	DB.Where("user_id=? and cart_id=?", userId, jsonInfo.CartId).Delete(jsonInfo)
	return 1
}
