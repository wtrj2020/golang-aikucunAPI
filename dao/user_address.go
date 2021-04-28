package dao

import "appApi/models"

func GetUserAddRess(userId, addressId int) (res models.User_address) {
	DB.First(&res, "user_id=? and address_id=? and data_flag=1", userId, addressId)
	return
}
