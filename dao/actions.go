package dao

import (
	"appApi/models"
	"strconv"
)

func ActList(jsonInfo models.ActPost) []models.ActListResult {
	var result []models.ActListResult
	Where := "fc_actions_items_goods.data_flag=1"
	//指定店铺
	if jsonInfo.ShopId > 0 {
		Where = Where + " and fc_actions_items_goods.shop_id=" + strconv.Itoa(jsonInfo.ShopId)
	}
	//指定活动状态
	if jsonInfo.ActStatus > 0 {
		Where = Where + " and fc_actions_items_goods.act_status=" + strconv.Itoa(jsonInfo.ActStatus)
	}

	//指定分页
	if jsonInfo.IsPage > 0 {
		offset := jsonInfo.PageSize * (jsonInfo.PageNum - 1)
		DB.Table("fc_actions_items_goods").Select("*").Where(Where).
			Limit(jsonInfo.PageSize).Offset(offset).
			Joins("right join fc_goods on fc_goods.goods_id=fc_actions_items_goods.goods_id").
			Scan(&result)
	} else {
		DB.Table("fc_actions_items_goods").Select("*").Where(Where).
			Joins("right join fc_goods on fc_goods.goods_id=fc_actions_items_goods.goods_id").
			Scan(&result)
	}

	//activityTemplate
	return result
}

//活动列表
func ActionsList(jsonInfo models.ActPost) (result []models.Actions_template) {
	Where := "fc_actions_template.data_flag=1"
	//指定分类
	if jsonInfo.CatId > 0 {
		Where = Where + " and fc_actions_template.cat_id=" + strconv.Itoa(jsonInfo.CatId)
	}
	//指定店铺
	if jsonInfo.ShopId > 0 {
		Where = Where + " and fc_actions_template.shop_id=" + strconv.Itoa(jsonInfo.ShopId)
	}
	//指定活动状态
	if jsonInfo.ActStatus > 0 {
		Where = Where + " and fc_actions_template.act_status=" + strconv.Itoa(jsonInfo.ActStatus)
	}

	//指定分页
	if jsonInfo.IsPage > 0 {
		offset := jsonInfo.PageSize * (jsonInfo.PageNum - 1)
		DB.Table("fc_actions_template").Select("*").Where(Where).
			Limit(jsonInfo.PageSize).Offset(offset).
			Scan(&result)
	} else {
		DB.Table("fc_actions_template").Select("*").Where(Where).
			//Joins("right join fc_actions_items on fc_actions_items.id=fc_actions_template.items_id").
			Scan(&result)
	}

	return
}

func ActionsGoodsList(jsonInfo models.ActPost) (res []models.Goods) {
	DB.Find(&res, "theme_id=? and data_flag=1", jsonInfo.ThemeId)
	return
}
