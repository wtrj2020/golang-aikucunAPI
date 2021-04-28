package controller

import (
	"appApi/models"
	"appApi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//https://app.doucang.com/api/biz/biz/app/activity/
// selectActivityByTag?pageNum=1&pageSize=8&activityTagId=1688b14e-9782-11e9-a46a-00163f00aa2a&isToday=true
// selectActivityByTag?pageNum=3&pageSize=8&activityTagId=ALL&isToday=true
///selectActivityBySalesCatetoryId?pageNum=1&pageSize=8&salesCategoryId=6&isToday=true
//活动列表????
func ActList(c *gin.Context) {
	var jsonInfo models.ActPost
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("登录使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	i, results := service.ActList(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    "aaagg",
		Data:   results,
	})
}

//活动列表
func ActionsList(c *gin.Context) {
	var jsonInfo models.ActPost
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	i, s, results := service.ActionsList(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   results,
	})
	//cart_num
}

//店铺活动的宝贝 详情
func ActionsGoodsList(c *gin.Context) {
	var jsonInfo models.ActPost
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	i, s, results := service.ActionsGoodsList(jsonInfo)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   results,
	})
}
