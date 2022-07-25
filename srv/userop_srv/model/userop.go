package model


//LeavingMessages 留言
type LeavingMessages struct {
	BaseModel
	User        int32  `gorm:"type:int;index"`
	MessageType int32  `gorm:"type:int comment '留言类型: 1(留言),2(投诉),3(询问),4(售后),5(求购)'"`
	Subject     string `gorm:"type:varchar(100)"`

	Message string //不限长度 gorm默认为text类型
	File    string `gorm:"type:varchar(200)"` //oss 保存文件url
}

func (LeavingMessages) TableName() string {
	return "leavingmessages"
}

type Address struct {
	BaseModel

	User         int32  `gorm:"type:int;index"`
	Province     string `gorm:"type:varchar(10)"`  //省
	City         string `gorm:"type:varchar(10)"`  //市
	District     string `gorm:"type:varchar(20)"`  //区
	Address      string `gorm:"type:varchar(100)"` //详细地址
	SignerName   string `gorm:"type:varchar(20)"`
	SignerMobile string `gorm:"type:varchar(11)"`
}

//UserFav 收藏
type UserFav struct {
	BaseModel
	//组合唯一索引 unique限制一个用户只能收藏一个商品
	User  int32 `gorm:"type:int;index:idx_user_goods,unique"`
	Goods int32 `gorm:"type:int;index:idx_user_goods,unique"`
}

func (UserFav) TableName() string {
	return "userfav"
}
