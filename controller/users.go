package controller

import (
	"appApi/models"
	"appApi/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var jsonInfo models.Users
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("登录使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	//jsonInfo.Lastip = c.ClientIP()
	s, s2, i := service.Login(jsonInfo)
	c.JSON(200, models.Result{
		Status: i,
		Msg:    s,
		Data:   s2,
	})
}

//注册新用户
func RegUser(c *gin.Context) {
	var jsonInfo models.Users
	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}
	jsonInfo.Lastip = c.ClientIP()
	s, i := service.RegUser(jsonInfo)
	c.JSON(200, models.Result{
		Status: i,
		Msg:    s,
		//Data:   nil,
	})
}

func MyAddressList(c *gin.Context) {
	var jsonInfo models.Users
	jsonInfo.UserId = c.MustGet("userId").(int)
	status, address := service.MyAddressList(jsonInfo)
	c.JSON(200, models.Result{
		Status: status,
		Msg:    "获取成功",
		Data:   address,
	})
}

func Member(c *gin.Context) {
	var jsonInfo models.Users
	jsonInfo.UserId = c.MustGet("userId").(int)
	status, member := service.Member(jsonInfo)
	c.JSON(200, models.Result{
		Status: status,
		Msg:    "获取成功",
		Data:   member,
	})
}

//添加地址
func AddMyAddress(c *gin.Context) {
	var jsonInfo models.User_address

	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}

	jsonInfo.UserId = c.MustGet("userId").(int)
	status, msg, address := service.AddMyAddress(jsonInfo)
	c.JSON(200, models.Result{
		Status: status,
		Msg:    msg,
		Data:   address,
	})
}

func EditIsdefaultAddress(c *gin.Context) {
	var jsonInfo models.User_address

	if c.BindJSON(&jsonInfo) != nil {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, models.Result{
			Status: -1,
			Msg:    "解析数据失败！",
			//Data:   nil,
		})
		return
	}

	jsonInfo.UserId = c.MustGet("userId").(int)
	status, address := service.EditIsdefaultAddress(jsonInfo)
	c.JSON(200, models.Result{
		Status: status,
		Msg:    "msg",
		Data:   address,
	})
}
