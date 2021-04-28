package models

type Result struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

type Member struct {
	UserInfo Users `json:"userInfo"`
}
type Shops struct {
	ShopId      int     `gorm:"Column:shop_id;PRIMARY_KEY" json:"shop_id"`
	UserId      int     `gorm:"Column:user_id" json:"user_id"`
	ShopName    string  `gorm:"Column:shop_name" json:"shop_name"`
	ShopPhone   string  `gorm:"Column:shop_phone" json:"shop_phone"`
	ShopCompany string  `gorm:"Column:shop_company" json:"shop_company"`
	ShopImg     string  `gorm:"Column:shop_img" json:"shop_img"`
	ShopAddress string  `gorm:"Column:shop_address" json:"shop_address"`
	ShopStatus  int     `gorm:"Column:shop_status" json:"shop_status"`
	ActStatus   int     `gorm:"Column:act_status" json:"act_status"`   //该店铺活动状态标记
	ShopNotice  string  `gorm:"Column:shop_notice" json:"shop_notice"` // 店铺公告
	ShopMoney   float64 `gorm:"Column:shop_money" json:"shop_money"`
	LockMoney   float64 `gorm:"Column:lock_money" json:"lock_money"`
	Sort        int     `gorm:"Column:sort" json:"sort"`
	DataFlag    int     `gorm:"Column:data_flag" json:"data_flag"`
	CreateTime  int64   `gorm:"Column:createTime" json:"createTime"`
	UpdateTime  int64   `gorm:"Column:update_time" json:"update_time"`
}

type Users struct {
	UserId     int     `gorm:"Column:user_id;PRIMARY_KEY" json:"user_id"`
	NickName   string  `gorm:"Column:nick_name" json:"nick_name"`
	HeadImg    string  `gorm:"Column:head_img" json:"head_img"`
	UserName   string  `gorm:"Column:user_name" json:"user_name"`
	UserPhone  string  `gorm:"Column:user_phone" json:"user_phone"`
	UserPass   string  `gorm:"Column:user_pass" json:"user_pass"`
	Salt       int     `gorm:"Column:salt" json:"-"`                // md5盐
	UserType   int     `gorm:"Column:user_type" json:"user_type"`   // 0普通,1店主,3管理员
	UserMoney  float64 `gorm:"Column:user_money" json:"user_money"` // 余额
	LockMoney  float64 `gorm:"Column:lock_money" json:"lock_money"` // 变动时锁定
	UserScore  int     `gorm:"Column:user_score" json:"user_score"` // 积分
	Token      string  `gorm:"Column:token" json:"token"`
	Status     int     `gorm:"Column:status" json:"status"`
	OpenId     string  `gorm:"Column:open_id" json:"open_id"`
	Lastip     string  `gorm:"Column:lastip" json:"lastip"`
	CreateTime int64   `gorm:"Column:create_time" json:"create_time"`
	UpdateTime int64   `gorm:"Column:update_time" json:"update_time"`
	DataFlag   int     `gorm:"Column:data_flag" json:"data_flag"` // -1为软删除

}

type Goods struct {
	GoodsId        int           `gorm:"Column:goods_id;PRIMARY_KEY" json:"goods_id"`
	ShopId         int           `gorm:"Column:shop_id" json:"shop_id"`   // 店主id
	GoodsSn        string        `gorm:"Column:goods_sn" json:"goods_sn"` // 商品编号
	GoodsName      string        `gorm:"Column:goods_name" json:"goods_name"`
	GoodsStock     int           `gorm:"Column:goods_stock" json:"goods_stock"`       // 库存总计
	CleanPrice     float64       `gorm:"Column:clean_price" json:"clean_price"`       // 清货价
	OriginalPrice  float64       `gorm:"Column:original_price" json:"original_price"` // 吊牌价
	GoodsPrice     float64       `gorm:"Column:goods_price" json:"goods_price"`       // 销售价
	Status         int           `gorm:"Column:status" json:"status"`
	CatId          int           `gorm:"Column:cat_id" json:"cat_id"`         // 分类id
	GoodsImg       string        `gorm:"Column:goods_img" json:"goods_img"`   // 首图
	GoodsImgs      string        `gorm:"Column:goods_imgs" json:"goods_imgs"` // 图片集
	GoodsImgsArray []string      `gorm:"-" json:"goods_imgs_array"`
	TodaySales     int           `gorm:"Column:today_sales" json:"today_sales"` // -1为软删除
	TotalSales     int           `gorm:"Column:total_sales" json:"total_sales"` // -1为软删除
	CreateTime     int64         `gorm:"Column:create_time" json:"create_time"`
	UpdateTime     int64         `gorm:"Column:update_time" json:"update_time"`
	DataFlag       int           `gorm:"Column:data_flag" json:"data_flag"` // -1为软删除
	IsEdit         bool          `gorm:"-"  json:"is_edit"`                 // 是否修改
	GoodsSpecs     []Goods_specs `gorm:"Column:goods_specs" json:"goods_specs"`
}

type Goods_cats struct {
	Id         int      `gorm:"Column:id" json:"id"`
	CatName    string   `gorm:"Column:cat_name" json:"cat_name"`   // 分类标题
	ShopLogo   string   `gorm:"Column:shop_logo" json:"shop_logo"` // 分类标题
	Address    string   `gorm:"Column:address" json:"address"`     // 分类标题
	Flow       int      `gorm:"Column:flow" json:"flow"`
	Sales      int      `gorm:"Column:sales" json:"sales"`
	Sort       int      `gorm:"Column:sort" json:"sort"`
	CreateTime int64    `gorm:"Column:create_time" json:"create_time"`
	UpdateTime int64    `gorm:"Column:update_time" json:"update_time"`
	GoodsImg   []string `gorm:"Column:goods_img" json:"goods_img"` // 首图
	DataFlag   int      `gorm:"Column:data_flag" json:"data_flag"` // -1为软删除
}

type Goods_specs struct {
	Id            int     `gorm:"Column:id;PRIMARY_KEY" json:"id"`
	ShopId        int     `gorm:"Column:shop_id" json:"shop_id"`   // 店主id
	GoodsSn       string  `gorm:"Column:goods_sn" json:"goods_sn"` // 商品规格的编号
	GoodsId       int     `gorm:"Column:goods_id" json:"goods_id"`
	SpecName      string  `gorm:"Column:spec_name;" json:"spec_name"`          // 颜色；尺码
	SpecColour    string  `gorm:"Column:spec_colour" json:"spec_colour"`       // 颜色
	SpecOnly      string  `gorm:"Column:spec_only" json:"spec_only"`           // 尺码
	GoodsStock    int     `gorm:"Column:goods_stock" json:"goods_stock"`       // 目前规格库存
	CleanPrice    float64 `gorm:"Column:clean_price" json:"clean_price"`       // 清货价
	OriginalPrice float64 `gorm:"Column:original_price" json:"original_price"` // 吊牌价
	GoodsPrice    float64 `gorm:"Column:goods_price" json:"goods_price"`       // 销售价
	Status        int     `gorm:"Column:status" json:"status"`
	GoodsImg      string  `gorm:"Column:goods_img" json:"goods_img"` // 首图
	CreateTime    int64   `gorm:"Column:create_time" json:"create_time"`
	UpdateTime    int64   `gorm:"Column:update_time" json:"update_time"`
	DataFlag      int     `gorm:"Column:data_flag" json:"data_flag"` // -1为软删除
	Commission    float64 `gorm:"commission" json:"commission"`
	//下面方便客户端处理数据
	CartNum       int `gorm:"-" json:"cart_num"`      // -
	IsCheck       int `gorm:"-" json:"is_check"`      // -
	Goods_spec_id int `gorm:"-" json:"goods_spec_id"` // -
}

type Carts struct {
	CartId      int   `gorm:"cart_id;PRIMARY_KEY" json:"cart_id"`
	ShopId      int   `gorm:"shop_id" json:"shop_id"`
	UserId      int   `gorm:"user_id" json:"user_id"`
	IsCheck     int   `gorm:"is_check" json:"is_check"`
	GoodsId     int   `gorm:"goods_id" json:"goods_id"`
	GoodsSpecId int   `gorm:"goods_spec_id" json:"goods_spec_id"`
	CartNum     int   `gorm:"cart_num" json:"cart_num"`
	CreateTime  int64 `gorm:"Column:create_time" json:"create_time"`
	UpdateTime  int64 `gorm:"Column:update_time" json:"update_time"`
}

/*
flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","s_addressId":2,"orderSrc":"ios","payType":1,"payCode":"wallets"}
flutter: post:请求url:http://c.judus.top:9999//app/orders/submit
flutter: -------end------------------------------------------
flutter: {"status":1,"domain":"http:\/\/c.judus.top:9999\/","msg":"\u63d0\u4ea4\u8ba2\u5355\u6210\u529f","data":"158374024545129852"}
flutter:
------post----------------------------------
flutter: {"tokenId":"958527b1df8bd81e5d6b4bb3a1704478","payPwd":"123123","orderNo":"158374024545129852","isBatch":1}
flutter: post:请求url:http://c.judus.top:9999//app/wallets/payByWallet

*/
type OrdersList struct {
	OrderStatus int        `gorm:"order_status" json:"order_status"`
	Payments    []Payments `json:"payments"`

	TotalMoney  float64      `gorm:"total_money" json:"total_money"`
	GoodsImg    string       `gorm:"goods_img" json:"goods_img"`
	Orderunique string       `gorm:"orderunique" json:"orderunique"` // 订单唯一编号
	AreaId      int          `gorm:"area_id" json:"area_id"`
	UserName    string       `gorm:"user_name" json:"user_name"`
	UserAddress string       `gorm:"user_address" json:"user_address"`
	UserPhone   string       `gorm:"user_phone" json:"user_phone"`
	PayTime     int64        `gorm:"pay_time" json:"pay_time"`
	CreateTime  int64        `gorm:"create_time" json:"create_time"`
	ShopsOrder  []ShopsOrder `json:"shops_order"`
	Orders      []Orders     `json:"orders"`
}
type Orders struct {
	OrderId         int           `gorm:"order_id;PRIMARY_KEY" json:"order_id"`
	OrderNo         string        `gorm:"order_no" json:"order_no"`
	ShopId          int           `gorm:"shop_id" json:"shop_id"`
	UserId          int           `gorm:"user_id" json:"user_id"`
	OrderStatus     int           `gorm:"order_status" json:"order_status"`
	GoodsMoney      float64       `gorm:"goods_money" json:"goods_money"`
	TotalMoney      float64       `gorm:"total_money" json:"total_money"`
	CleanPrice      float64       `gorm:"Column:clean_price" json:"clean_price"` // 清货价
	PayType         int           `gorm:"pay_type" json:"pay_type"`
	PayFrom         string        `gorm:"pay_from" json:"pay_from"` // 支付方式
	IsPay           int           `gorm:"is_pay" json:"is_pay"`
	AreaId          int           `gorm:"area_id" json:"area_id"`
	AreaIdPath      string        `gorm:"area_id_path" json:"area_id_path"`
	UserName        string        `gorm:"user_name" json:"user_name"`
	UserAddress     string        `gorm:"user_address" json:"user_address"`
	UserPhone       string        `gorm:"user_phone" json:"user_phone"`
	OrderScore      int           `gorm:"order_score" json:"order_score"`
	OrderRemarks    string        `gorm:"order_remarks" json:"order_remarks"` // 订单备注
	OrderSrc        string        `gorm:"order_src" json:"order_src"`         // 订单来源设备
	Orderunique     string        `gorm:"orderunique" json:"orderunique"`     // 订单唯一编号
	ExpressId       int           `gorm:"express_id" json:"express_id"`
	ExpressNo       string        `gorm:"express_no" json:"express_no"`
	DataFlag        int           `gorm:"data_flag" json:"data_flag"`
	NoticeDeliver   int           `gorm:"notice_deliver" json:"notice_deliver"`     // 提醒发货 0:未提醒 1:已提醒
	TotalCommission float64       `gorm:"total_commission" json:"total_commission"` // 佣金
	AddressId       int           `gorm:"-" json:"address_id"`
	OrderGoods      []Order_goods `gorm:"-" json:"Order_goods"`
	PayTime         int64         `gorm:"pay_time" json:"pay_time"`
	DeliveryTime    int64         `gorm:"delivery_time" json:"delivery_time"`
	ReceiveTime     int64         `gorm:"delivery_time" json:"receive_time"`
	CreateTime      int64         `gorm:"create_time" json:"create_time"`
	UpdateTime      int64         `gorm:"Column:update_time" json:"update_time"`
	Payments        []Payments    `json:"payments"`
}

type Order_goods struct {
	Id             int     `gorm:"id;PRIMARY_KEY" json:"id"`
	OrderId        int     `gorm:"order_id" json:"order_id"`
	GoodsId        int     `gorm:"goods_id" json:"goods_id"`
	GoodsNum       int     `gorm:"goods_num" json:"goods_num"`
	GoodsPrice     float64 `gorm:"goods_price" json:"goods_price"`
	CleanPrice     float64 `gorm:"Column:clean_price" json:"clean_price"` // 清货价
	GoodsSpecId    int     `gorm:"goods_spec_id" json:"goods_spec_id"`
	GoodsSpecNames string  `gorm:"goods_spec_names" json:"goods_spec_names"`
	GoodsName      string  `gorm:"goods_name" json:"goods_name"`
	GoodsImg       string  `gorm:"goods_img" json:"goods_img"`
	GoodsSn        string  `gorm:"goods_sn" json:"goods_sn"`
	Commission     float64 `gorm:"commission" json:"commission"`
	DataFlag       int     `gorm:"data_flag" json:"data_flag"`
	CreateTime     int64   `gorm:"Column:create_time" json:"create_time"`
	UpdateTime     int64   `gorm:"Column:update_time" json:"update_time"`
}

type User_address struct {
	AddressId   int    `gorm:"address_id;PRIMARY_KEY" json:"address_id"`
	UserId      int    `gorm:"user_id" json:"user_id"`
	UserName    string `gorm:"user_name" json:"user_name"`
	UserPhone   string `gorm:"user_phone" json:"user_phone"`
	AreaIdPath  string `gorm:"-" json:"areaId_path"`
	AreaId      int    `gorm:"area_id" json:"area_id"`
	UserAddress string `gorm:"user_address" json:"user_address"`
	IsDefault   int    `gorm:"is_default" json:"is_default"`
	DataFlag    int    `gorm:"data_flag" json:"data_flag"`
	CreateTime  int64  `gorm:"create_time" json:"create_time"`
	UpdateTime  int64  `gorm:"Column:update_time" json:"update_time"`
}

type GetCarts struct {
	ShopId      int     `gorm:"Column:shop_id" json:"shop_id"`
	CartId      int     `gorm:"cart_id;PRIMARY_KEY" json:"cart_id"`
	GoodsSpecId int     `gorm:"goods_spec_id" json:"goods_spec_id"`
	IsCheck     int     `gorm:"is_check" json:"is_check"`
	CartNum     int     `gorm:"cart_num" json:"cart_num"`
	GoodsId     int     `gorm:"goods_id" json:"goods_id"`
	GoodsName   string  `gorm:"goods_name" json:"goods_name"`
	GoodsImg    string  `gorm:"goods_img" json:"goods_img"`
	SpecName    string  `gorm:"Column:spec_name;" json:"spec_name"`    // 颜色；尺码
	GoodsSn     string  `gorm:"Column:goods_sn" json:"goods_sn"`       // 商品规格的编号
	SpecColour  string  `gorm:"Column:spec_colour" json:"spec_colour"` // 颜色
	SpecOnly    string  `gorm:"Column:spec_only" json:"spec_only"`     // 尺码
	GoodsStock  int     `gorm:"Column:goods_stock" json:"goods_stock"` // 目前规格库存
	GoodsPrice  float64 `gorm:"Column:goods_price" json:"goods_price"` // 销售价
	CleanPrice  float64 `gorm:"Column:clean_price" json:"clean_price"` // 清货价
	ShopName    string  `gorm:"Column:shop_name" json:"shop_name"`
	ShopImg     string  `gorm:"Column:shop_img" json:"-"`
	ShopAddress string  `gorm:"Column:shop_address" json:"-"`
	Commission  float64 `gorm:"Column:commission" json:"commission"` // 佣金

}

type ShopsOrder struct {
	ShopId   int    `gorm:"Column:shop_id;PRIMARY_KEY" json:"shop_id"`
	ShopName string `gorm:"Column:shop_name" json:"shop_name"`
	//ShopPhone   string  `gorm:"Column:shop_phone" json:"shop_phone"`
	ShopCompany string `gorm:"Column:shop_company" json:"shop_company"`
	ShopImg     string `gorm:"Column:shop_img" json:"shop_img"`
	ShopAddress string `gorm:"Column:shop_address" json:"shop_address"`
	//ShopStatus  int     `gorm:"Column:shop_status" json:"shop_status"`
	ShopPrice      float64      `gorm:"Column:shop_price" json:"shop_price"`
	ShopCleanPrice float64      `gorm:"Column:shop_clean_price" json:"shop_clean_price"`
	CartNum        int          `json:"cart_num"`
	Address        User_address `json:"user_address"`
	List           []GetCarts   `json:"listx"`
	Orders         []Orders     `json:"orders"`
}

type Payments struct {
	Id        int    `gorm:"id;PRIMARY_KEY" json:"id"`
	OrderNo   string `gorm:"-" json:"order_no"`
	PayFrom   string `gorm:"pay_from" json:"pay_from"`
	PayName   string `gorm:"pay_name" json:"pay_name"`
	PayOrder  int    `gorm:"pay_order" json:"pay_order"`
	PayConfig string `gorm:"pay_config" json:"pay_config"`
	Enabled   int    `gorm:"enabled" json:"enabled"`

	UserId int `gorm:"-" json:"user_id"`

	//请求
}

type Log_orders struct {
	LogId       int    `gorm:"log_id;PRIMARY_KEY" json:"log_id"`
	OrderId     int    `gorm:"order_id" json:"order_id"`
	OrderStatus int    `gorm:"order_status" json:"order_status"`
	LogContent  string `gorm:"log_content" json:"log_content"`
	LogUserId   int    `gorm:"log_user_id" json:"log_user_id"`
	LogType     int    `gorm:"log_type" json:"log_type"`
	LogTime     string `gorm:"-" json:"log_time"`
	CreateTime  int64  `gorm:"create_time" json:"create_time"`
}

type Log_moneys struct {
	Id         int     `gorm:"id;PRIMARY_KEY" json:"id"`
	TargetType int     `gorm:"target_type" json:"target_type"`
	TargetId   int     `gorm:"target_id" json:"target_id"`
	DataId     int     `gorm:"data_id" json:"data_id"`
	DataSrc    string  `gorm:"data_src" json:"data_src"`
	Remark     string  `gorm:"remark" json:"remark"`
	MoneyType  int     `gorm:"money_type" json:"money_type"`
	Money      float64 `gorm:"money" json:"money"`
	TradeNo    string  `gorm:"trade_no" json:"trade_no"`
	PayType    string  `gorm:"pay_type" json:"pay_type"`
	DataFlag   int     `gorm:"data_flag" json:"data_flag"`
	CreateTime int64   `gorm:"create_time" json:"create_time"`
}

type ActPost struct {
	CatId     int `json:"CatId"`
	ThemeId   int `json:"ThemeId"`
	ShopId    int `json:"ShopId"`
	PageNum   int `json:"PageNum"`
	PageSize  int `json:"PageSize"`
	ActId     int `json:"ActId"`
	IsPage    int `json:"IsPage"`
	ActStatus int `json:"ActStatus"`
	CatTagId  int `json:"CatTagId"`
}
type ActListResult struct {
	Id         int     `gorm:"id" json:"id"`
	ShopId     int     `gorm:"shop_id" json:"shop_id"`
	GoodsId    int     `gorm:"goods_id" json:"goods_id"`
	GoodsSn    string  `gorm:"Column:goods_sn" json:"goods_sn"` // 商品编号
	GoodsName  string  `gorm:"Column:goods_name" json:"goods_name"`
	ShopName   string  `gorm:"Column:shop_name" json:"shop_name"`
	GoodsPrice float64 `gorm:"Column:goods_price" json:"goods_price"` // 销售价
	Commission float64 `gorm:"commission" json:"commission"`          //佣金
	GoodsStock int     `gorm:"Column:goods_stock" json:"goods_stock"` // 库存总计
	CatId      int     `gorm:"Column:cat_id" json:"cat_id"`           // 分类id
	GoodsImg   string  `gorm:"Column:goods_img" json:"goods_img"`     // 首图
	GoodsImgs  string  `gorm:"Column:goods_imgs" json:"goods_imgs"`   // 首图
}
type Actions struct {
	Id             int    `gorm:"id;PRIMARY_KEY" json:"id"`
	ActionName     string `gorm:"action_name" json:"action_name"`
	ActionMethod   string `gorm:"action_method" json:"action_method"`     // 活动玩法
	ActionDescribe string `gorm:"action_describe" json:"action_describe"` // 活动介绍
	Sort           int    `gorm:"sort" json:"sort"`
	Path           string `gorm:"-" json:"path"`
	Path2          string `gorm:"-" json:"path2"`
	DataFlag       int    `gorm:"data_flag" json:"data_flag"`
	IsEdit         bool   `gorm:"-"  json:"is_edit"` // 是否修改
}
type Actions_items struct {
	Id             int    `gorm:"id;PRIMARY_KEY" json:"id"`
	ActionsId      int    `gorm:"actions_id" json:"actions_id"`
	ActionName     string `gorm:"action_name" json:"action_name"`
	ActionMethod   string `gorm:"action_method" json:"action_method"`     // 活动玩法
	ActionDescribe string `gorm:"action_describe" json:"action_describe"` // 活动介绍
	PubStart       int    `gorm:"pub_start" json:"pub_start"`             // 报名时间
	PubEnd         int    `gorm:"pub_end" json:"pub_end"`                 // 报名结束
	ActStart       int    `gorm:"act_start" json:"act_start"`             // 活动时间
	ActEnd         int    `gorm:"act_end" json:"act_end"`                 // 活动结束
	PubStatus      int    `gorm:"Column:pub_status" json:"pub_status"`
	ActStatus      int    `gorm:"Column:act_status" json:"act_status"`
	Status         int    `gorm:"-" json:"status"`
	PubName        string `gorm:"pub_name" json:"pub_name"`
	ActName        string `gorm:"act_name" json:"act_name"`
	StatusName     string `gorm:"-" json:"status_name"`
	DataFlag       int    `gorm:"data_flag" json:"data_flag"`
	IsEdit         bool   `gorm:"-"  json:"is_edit"` // 是否修改

}

type Actions_items_goods struct {
	Id       int `gorm:"id;PRIMARY_KEY" json:"id"`
	ItemsId  int `gorm:"items_id" json:"items_id"`
	GoodsId  int `gorm:"goods_id" json:"goods_id"`
	ShopId   int `gorm:"shop_id" json:"shop_id"`
	Sort     int `gorm:"sort" json:"sort"`
	DataFlag int `gorm:"data_flag" json:"data_flag"`

	ActionName   string `gorm:"action_name" json:"action_name"`
	ActionMethod string `gorm:"action_method" json:"action_method"` // 活动玩法

	PubName   string `gorm:"pub_name" json:"pub_name"`
	PubStatus int    `gorm:"Column:pub_status" json:"pub_status"`
	PubStart  int    `gorm:"pub_start" json:"pub_start"` // 报名时间
	PubEnd    int    `gorm:"pub_end" json:"pub_end"`     // 报名结束

	ActName   string `gorm:"act_name" json:"act_name"`
	ActStatus int    `gorm:"Column:act_status" json:"act_status"`
	ActStart  int    `gorm:"act_start" json:"act_start"` // 活动时间
	ActEnd    int    `gorm:"act_end" json:"act_end"`     // 活动结束
}

type Actions_template struct {
	ThemeId        int      `gorm:"theme_id;PRIMARY_KEY" json:"theme_id"`
	ShopId         int      `gorm:"shop_id" json:"shop_id"`
	ItemsId        int      `gorm:"items_id" json:"items_id"`
	ShopName       string   `gorm:"shop_name" json:"shop_name"`
	ThemeName      string   `gorm:"theme_name" json:"theme_name"`
	ThemeNotice    string   `gorm:"theme_notice" json:"theme_notice"`
	CatId          int      `gorm:"cat_id" json:"cat_id"`
	CatName        string   `gorm:"cat_name" json:"cat_name"`
	ActionName     string   `gorm:"action_name" json:"action_name"`
	GoodsIds       string   `gorm:"goods_ids" json:"goods_ids"`
	GoodsImgs      string   `gorm:"-" json:"goods_imgs"`
	GoodsImgsArray []string `gorm:"-" json:"goods_imgs_array"`
	PubStatus      int      `gorm:"pub_status" json:"pub_status"`
	PubName        string   `gorm:"pub_name" json:"pub_name"`
	PubStart       int      `gorm:"pub_start" json:"pub_start"`
	PubEnd         int      `gorm:"pub_end" json:"pub_end"`
	ActName        string   `gorm:"act_name" json:"act_name"`
	ActStatus      int      `gorm:"act_status" json:"act_status"` // 1开始2即将3结束
	ActStart       int      `gorm:"act_start" json:"act_start"`
	ActEnd         int      `gorm:"act_end" json:"act_end"`
	Sort           int      `gorm:"sort" json:"sort"`
	DataFlag       int      `gorm:"data_flag" json:"data_flag"`
	CreateTime     int64    `gorm:"create_time" json:"create_time"` // 添加时间
	UpdateTime     int64    `gorm:"update_time" json:"update_time"` // 更新时间
	IsEdit         int      `gorm:"-" json:"is_edit"`
	Number         int      `gorm:"-" json:"number"`
}
