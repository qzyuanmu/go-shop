package shop

import (
	"qzyuanmu/jtool"
)

//店铺行业分类
type ShopType struct {
	Id   int    `xorm:"pk autoincr"`
	Name string `xorm:"varchar(250)"`
	Sort int    `xorm:"default 100"`
}

func (this ShopType) TableName() string {
	return "j_shop_type"
}

func GetShopType() []ShopType {
	ts := make([]ShopType, 0)
	err := jtool.Engine.Asc("Sort").Find(&ts)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}

	return ts
}
