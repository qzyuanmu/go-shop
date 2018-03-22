package shop

import "qzyuanmu/jtool"

type OrderDet struct {
	Id              int    `xorm:"pk autoincr"`
	ProductId       int    `xorm:"default 0"`
	ProductTypeId   string `xorm:"varchar(250)"`
	ProductTypeName string `xorm:"varchar(250)"`
	ProductImg      string `xorm:"varchar(550)"`
	ProductName     string `xorm:"varchar(250)"`
	Quantity        int    `xorm:"default 0"`
	OriPrice        int    `xorm:"default 0"`
	Price           int    `xorm:"default 0"`
	Zhim            int    `xorm:"default 0"`
	//红包比率
	RedLevel int `xorm:"default 0"`
	ShopId   int `xorm:"default 0"`
	OrderId  int `xorm:"default 0"`
	UserId   int `xorm:"default 0"`
}

func (p *OrderDet) TableName() string {
	return "j_shop_order_det"

}

func GetOrderDet(oId int) []OrderDet {
	os := make([]OrderDet, 0)
	err := jtool.Engine.Desc("Id").Where("OrderId=?", oId).Find(&os)
	if err != nil {
		jtool.Log.Error("获取订单明细失败", err)
	}
	return os
}
