package controller

import (
	"appApi/dao"
	"appApi/models"
	"appApi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGoodsList(c *gin.Context) {
	shopId := c.MustGet("shopId").(int)
	var jsonInfo models.Goods
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	s, i, getGoodsList := service.GetGoodsList(shopId, jsonInfo.Status)

	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   getGoodsList,
	})
}

//产品分类下的宝贝
func GoodsCatsGoods(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	s, i, getGoodsList := service.GoodsCatsGoods(Id)

	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   getGoodsList,
	})
}

//获取宝贝详情
func GoodsDetails(c *gin.Context) {
	var jsonInfo models.Goods
	GoodsId, _ := strconv.Atoi(c.Query("goods_id"))
	jsonInfo.GoodsId = GoodsId
	//if c.BindJSON(&jsonInfo) != nil {
	//	log.Println("注册使用数据格式异常！")
	//	c.JSON(http.StatusOK, models.Result{
	//		Status: -1,
	//		Msg:    "解析数据失败！",
	//		//Data:   nil,
	//	})
	//	return
	//}
	s, i, goods := service.GoodsDetails(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   goods,
	})
}

//流量重灾区，将来改成redis
func GoodsCats(c *gin.Context) {
	goodsCats := dao.GoodsCats()
	for key, value := range goodsCats {
		ss := dao.GoodsCatsGoodsShowcase(value.Id)
		for _, v := range ss {
			goodsCats[key].GoodsImg = append(goodsCats[key].GoodsImg, v.GoodsImg)
		}
	}
	c.JSON(http.StatusOK, models.Result{
		Status: 1,
		Msg:    "获取成功",
		Data:   goodsCats,
	})
}
