package dao

import (
	"appApi/models"
	"appApi/utils/tools"
	"strconv"
	"time"
)

//根据手机号获得用户信息
func UserInfo(jsonInfo models.Users) (user models.Users) {
	DB.Find(&user, "user_phone=?  and data_flag=1", jsonInfo.UserPhone)
	return

}

//通过token获取用户的id
func GetUserId(tokenId string) (res models.Users) {
	DB.Find(&res, "token=? and data_flag=1", tokenId)
	return
}

//通过userId获取用户的id
func GetUserIdX(userId int) (res models.Users) {

	DB.Find(&res, "user_id=? and data_flag=1", userId)
	return
}

//写入token
func EditToken(users models.Users, tokenId string) {
	DB.Model(&users).UpdateColumn("token", tokenId)
}

//注册新用户
func RegUser(jsonInfo models.Users) (user models.Users) {
	Salt := int(time.Now().Unix())
	user = models.Users{
		UserPhone:  jsonInfo.UserPhone,
		UserPass:   tools.Md5_salt(jsonInfo.UserPass, strconv.Itoa(Salt)),
		Salt:       Salt,
		Status:     1,
		OpenId:     "",
		Lastip:     jsonInfo.Lastip,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		DataFlag:   1,
	}
	DB.Create(&user)
	return
}

//地址列表
func MyAddressList(jsonInfo models.Users) (address []models.User_address) {
	DB.Find(&address, "user_id=? and data_flag=1", jsonInfo.UserId)
	return
}

//个人中心
//func Member(jsonInfo models.Users) (users models.Users) {
//	GetUserIdX.
//	//DB.Find(&users, "user_id=? and data_flag=1", jsonInfo.UserId)
//	return
//}

//修改默认你地址
func EditIsdefaultAddress(jsonInfo models.User_address) (address []models.User_address) {

	if jsonInfo.DataFlag > 0 {
		DelMyAddress(jsonInfo)
		return
	}

	DB.Model(models.User_address{}).
		Where("address_id = ? and user_id=?", jsonInfo.AddressId, jsonInfo.UserId).
		Update("is_default", 1)

	DB.Model(models.User_address{}).
		Where("address_id != ? and user_id=?", jsonInfo.AddressId, jsonInfo.UserId).
		Update("is_default", 0)

	return
}

//删除
func DelMyAddress(jsonInfo models.User_address) int {
	DB.Where("user_id=? and address_id=?", jsonInfo.UserId, jsonInfo.AddressId).Delete(jsonInfo)
	return 1
}

//添加地址
func AddMyAddress(jsonInfo models.User_address) (address models.User_address) {
	address = models.User_address{
		//AddressId:   0,
		UserId:    jsonInfo.UserId,
		UserName:  jsonInfo.UserName,
		UserPhone: jsonInfo.UserPhone,
		//AreaIdPath:  "",
		AreaId:      jsonInfo.AreaId,
		UserAddress: jsonInfo.UserAddress,
		IsDefault:   1,
		DataFlag:    1,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
	DB.Create(&address)
	if address.AddressId > 0 {
		DB.Model(models.User_address{}).
			Where("address_id != ? and user_id=?", address.AddressId, address.UserId).
			Update("is_default", 0)
	}

	//DB.Find(&address, "user_id=? and data_flag=1", jsonInfo.UserId)
	return
}
