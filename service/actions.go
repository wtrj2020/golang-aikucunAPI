package service

import (
	"appApi/dao"
	"appApi/models"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"strings"
)

func ActList(jsonInfo models.ActPost) (int, []models.ActListResult) {
	actListResults := dao.ActList(jsonInfo)
	return 1, actListResults
}

//活动列表
func ActionsList(jsonInfo models.ActPost) (int, string, []models.Actions_template) {
	result := dao.ActionsList(jsonInfo)
	for i, v := range result {
		result[i].GoodsImgsArray = strings.Split(v.GoodsImgs, ",")
		fmt.Println()
	}
	return 1, "获取成功", result
}

func ActionsGoodsList(jsonInfo models.ActPost) (int, string, []models.Goods) {
	result := dao.ActionsGoodsList(jsonInfo)

	var goodsIds = []int{}
	for i, v := range result {
		goodsIds = append(goodsIds, v.GoodsId)
		result[i].GoodsImgsArray = strings.Split(v.GoodsImgs, ",")

	}
	goodsSpecs := dao.GetByGoodsSpecBatch(goodsIds)
	bytes, _ := json.Marshal(goodsSpecs)

	for i, v := range result {
		jsonq := gojsonq.New().JSONString(string(bytes)).Where("goods_id", "=", v.GoodsId).Get()
		lang, _ := json.Marshal(jsonq)
		var res []models.Goods_specs
		_ = json.Unmarshal([]byte(lang), &res)
		result[i].GoodsSpecs = res
	}

	return 1, "获取成功", result
}
