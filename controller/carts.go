package controller

import (
	"appApi/models"
	"appApi/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MyCartList(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	i, s, orders := service.MyCartList(userId)
	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   orders,
	})
}

func AddCart(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo []models.Carts
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	s, i := service.AddCart(userId, jsonInfo)

	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   i,
	})
}

func DelOnlyCart(c *gin.Context) {
	userId := c.MustGet("userId").(int)
	var jsonInfo models.Carts
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	s, i := service.DelOnlyCart(userId, jsonInfo)

	c.JSON(http.StatusOK, models.Result{
		Status: i,
		Msg:    s,
		Data:   i,
	})
}
