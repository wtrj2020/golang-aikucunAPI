package dao

import (
	"appApi/models"
	"fmt"
	"time"
)

//指定店铺宝贝列表
//商铺id，状态id
func GetGoodsList(shopId, status int) (res []models.Goods) {
	DB.Order("goods_id desc").Find(&res, "shop_id=? and status=? and data_flag=1", shopId, status)
	return
}

//产品分类下的宝贝
func GoodsCatsGoods(id int) (res []models.Goods) {
	DB.Order("goods_id desc").Find(&res, "cat_id=? and data_flag=1", id)
	return
}

//产品分类
func GoodsCats() (res []models.Goods_cats) {
	fmt.Println(time.Now().Format(TimeFormat))

	DB.Order("sort asc").Find(&res, models.Goods_cats{DataFlag: 1})

	return
}

//产品分类的橱窗宝贝（按今天销量展示）
func GoodsCatsGoodsShowcase(id int) (res []models.Goods) {
	DB.Limit(3).Order("today_sales desc").Find(&res, "cat_id=? and data_flag=1", id)
	return
}

//根据GoodsId查找产品
func GetByGoodsId(GoodsId int) (res models.Goods) {
	DB.First(&res, "goods_id=? and status=1 and data_flag=1", GoodsId)
	return
}

//根据GoodsSn查找产品
func GetByGoodsSn(GoodsSn string) (res models.Goods) {
	DB.First(&res, "goods_sn=? and status=1 and data_flag=1", GoodsSn)
	return
}

//根据GoodsId取产品规格
func GetByGoodsSpec(GoodsId int) (goodsSpecs []models.Goods_specs) {
	DB.Find(&goodsSpecs, "goods_id=? and status=1 and  data_flag=1", GoodsId)
	return
}

//批量根据GoodsId取产品规格
func GetByGoodsSpecBatch(GoodsIds []int) (goodsSpecs []models.Goods_specs) {
	DB.Find(&goodsSpecs, "goods_id IN (?) and status=1 and  data_flag=1", GoodsIds)
	return
}
