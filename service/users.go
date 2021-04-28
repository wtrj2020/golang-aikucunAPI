package service

import (
	"appApi/dao"
	"appApi/models"
	"appApi/utils/tools"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

func Login(jsonInfo models.Users) (string, string, int) {
	userInfo := dao.UserInfo(jsonInfo)
	if userInfo.UserId == 0 {
		return "用户不存在", "", -1
	}
	if userInfo.UserType != 0 {
		return "只有普通用户才能登录", "", -1
	}

	if userInfo.UserPass != tools.Md5_salt(jsonInfo.UserPass, strconv.Itoa(userInfo.Salt)) {
		return "密码错误", "", -1
	}

	tokenId := tools.Guid(userInfo.UserId)
	dao.EditToken(userInfo, tokenId)
	//删除之前的redis token
	_, e := dao.Redisdb.Del(userInfo.Token).Result()
	if e != nil {
		log.Println(e, "删除key失败，说明redis没有开启，改走数据库")
		return "登录成功", tokenId, 1
	}
	//redis写入新的token
	userInfo.Token = tokenId
	bytes, _ := json.Marshal(userInfo)
	dao.Redisdb.Set(tokenId, string(bytes), 0)
	return "登录成功", tokenId, 1
}

func RegUser(jsonInfo models.Users) (string, int) {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	rgx := regexp.MustCompile(regular)
	if !rgx.MatchString(jsonInfo.UserPhone) {
		return "错误手机号", -1
	}
	if len(jsonInfo.UserPass) < 6 {
		return "至少六位密码", -1
	}
	userInfo := dao.UserInfo(jsonInfo)
	if userInfo.UserId > 0 {
		return "用户已存在", -1
	}
	regUser := dao.RegUser(jsonInfo)

	if regUser.UserId == 0 {
		return "注册失败", -1
	}
	return "注册成功", 1
}

func MyAddressList(jsonInfo models.Users) (status int, address []models.User_address) {
	address = dao.MyAddressList(jsonInfo)
	return 1, address
}

//个人中心
func Member(jsonInfo models.Users) (status int, member models.Member) {
	UserInfo := dao.GetUserIdX(jsonInfo.UserId)
	member.UserInfo = UserInfo
	return 1, member
}
func AddMyAddress(jsonInfo models.User_address) (status int, msg string, address models.User_address) {
	addMyAddress := dao.AddMyAddress(jsonInfo)
	if addMyAddress.AddressId == 0 {
		return -1, "添加失败", addMyAddress
	}
	return 1, "添加成功", addMyAddress
}

func EditIsdefaultAddress(jsonInfo models.User_address) (status int, address []models.User_address) {

	address = dao.EditIsdefaultAddress(jsonInfo)
	return 1, address
}
