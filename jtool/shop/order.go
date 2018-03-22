package shop

import (
	"qzyuanmu/jtool"
	"time"
)

type Order struct {
	Id           int    `xorm:"pk autoincr"`
	OrderId      string `xorm:"varchar(250)"`
	Address      string `xorm:"varchar(250)"`
	ProvinceId   int    `xorm:"default 0"`
	ProvinceName string `xorm:"varchar(250)"`
	CityId       int    `xorm:"default 0"`
	CityName     string `xorm:"varchar(250)"`
	DistrictId   int    `xorm:"default 0"`
	DistrictName string `xorm:"varchar(250)"`
	ZipCode      string `xorm:"varchar(50)  "`
	RealName     string `xorm:"varchar(250)"`
	Email        string `xorm:"varchar(50)"`
	Mobile       string `xorm:"varchar(50)  "`
	IsValid      int    `xorm:"default 1"`
	//送货方式
	SendType string `xorm:"varchar(50)  "`
	SumPrice int    `xorm:"default 0"`
	SumMoney int    `xorm:"default 0"`
	SumZhim  int    `xorm:"default 0"`
	SumRed   int    `xorm:"default 0"`
	BestTime string `xorm:"varchar(50)  "`
	Memo     string `xorm:"varchar(550)  "`
	Ip       string `xorm:"varchar(50)  "`

	//1 等待付款  2 已付款 3 已发货 4 已完成
	Flg        int       `xorm:"default 1"`
	FlgValue   string    `xorm:"varchar(50)  "`
	PayFlg     int       `xorm:"default 1"`
	PayValue   string    `xorm:"varchar(50)  "`
	CreateTime time.Time `xorm:"created "`
	UpdateTime time.Time `xorm:"created updated "`
	ShopId     int       `xorm:"default 0"`

	UserId int `xorm:"default 0"`
}

func (u *Order) TableName() string {
	return "j_shop_order"
}

func GetOrder(uId int) []Order {
	os := make([]Order, 0)
	err := jtool.Engine.Desc("Id").Where("UserId=?", uId).Find(&os)
	if err != nil {
		jtool.Log.Error("根据用户获取订单失败", err)
	}
	return os
}
func GetOrderByShop(shopId int) []Order {
	os := make([]Order, 0)
	err := jtool.Engine.Desc("Id").Where("ShopId=?", shopId).Find(&os)
	if err != nil {
		jtool.Log.Error("根据店铺获取订单失败", err)
	}
	return os
}
