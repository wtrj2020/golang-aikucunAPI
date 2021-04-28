package dao

import "appApi/models"

//根据userid获取店铺信息
func GetByShops(userId int) (shops models.Shops) {
	DB.First(&shops, "user_id=? and data_flag=1", userId)
	return
}

func GetAllShops() (shops []models.Shops) {
	DB.Find(&shops, "data_flag=1")
	return
}
