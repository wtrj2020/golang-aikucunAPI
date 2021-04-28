package service

import (
	"appApi/dao"
	"appApi/models"
)

//指定店铺宝贝列表
func GetGoodsList(shopId, status int) (string, int, []models.Goods) {
	var getGoodsList []models.Goods
	if status == 0 {
		return "status为空", -1, getGoodsList
	}
	getGoodsList = dao.GetGoodsList(shopId, status)
	return "获取成功", 1, getGoodsList
}

//产品分类下的宝贝
func GoodsCatsGoods(id int) (string, int, []models.Goods) {
	var GoodsCatsGoods []models.Goods
	GoodsCatsGoods = dao.GoodsCatsGoods(id)
	return "获取成功", 1, GoodsCatsGoods
}

func GoodsDetails(jsonInfo models.Goods) (string, int, models.Goods) {

	//取商品详情
	goods := dao.GetByGoodsId(jsonInfo.GoodsId)
	if goods.GoodsId == 0 {
		return "无此产品", -1, goods
	}
	//取规格详情
	goodsSpecs := dao.GetByGoodsSpec(jsonInfo.GoodsId)
	if len(goodsSpecs) == 0 {
		return "规格异常", -1, goods
	}
	//组合起来
	goods.GoodsSpecs = goodsSpecs
	return "获取成功", 1, goods

}
