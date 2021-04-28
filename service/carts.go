package service

import (
	"appApi/dao"
	"appApi/models"
)

func MyCartList(userId int) (int, string, []models.ShopsOrder) {
	//整合购物车信息
	//ischeck0显示全部
	allShopCartList := dao.AllShopCartList(userId, 0)
	if len(allShopCartList) == 0 {
		return 1, "购物车为空", allShopCartList
	}
	return 1, "获取成功", allShopCartList
}

//加入购物车
func AddCart(userId int, jsonInfo []models.Carts) (string, int) {
	cartNum := dao.AddCart(userId, jsonInfo)
	if cartNum == 0 {
		return "加入购物车失败", 0
	}
	return "加入购物车成功", cartNum
}

//删除购物车商品
func DelOnlyCart(userId int, jsonInfo models.Carts) (string, int) {
	cartNum := dao.DelOnlyCart(userId, jsonInfo)
	if cartNum == 0 {
		return "加入购物车失败", 0
	}
	return "加入购物车成功", cartNum
}
