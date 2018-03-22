package shop

import (
	"qzyuanmu/jtool"
	"time"
)

type Shop struct {
	Id           int    `xorm:"pk autoincr"`
	ProvinceName string `xorm:"varchar(250)"`
	ProvinceId   int    `xorm:"default 0"`

	CityName string `xorm:"varchar(250)"`
	CityId   int    `xorm:"default 0"`

	DistrictName string `xorm:"varchar(250)"`
	DistrictId   int    `xorm:"default 0"`
	//店铺类型 商铺 企业
	ShopType string `xorm:"varchar(250)"`

	//店铺类型 服装业等
	ShopClass string `xorm:"varchar(250)"`
	Name      string `xorm:"varchar(250)"`
	Address   string `xorm:"varchar(250)"`

	Contact       string `xorm:"varchar(250)"`
	Introducer    string `xorm:"varchar(50)"`
	YYzhizhao     string `xorm:"varchar(250)"`
	YYzhizhaoInfo string `xorm:"varchar(250)"`
	YYzhizhaoImg  string `xorm:"varchar(250)"`

	Phone string `xorm:"varchar(250)"`

	//1 未审批 2  已审批
	Flg string `xorm:"varchar(250)"`

	//生产许可证
	ShengCXKZ       string    `xorm:"varchar(250)"`
	ShengCXKZImg    string    `xorm:"varchar(250)"`
	ShengFZ         string    `xorm:"varchar(250)"`
	ShengFZImg      string    `xorm:"varchar(250)"`
	ShengFZFImg     string    `xorm:"varchar(250)"`
	MapImg          string    `xorm:"varchar(250)"`
	Lat             float32   `xorm:"default 0"`
	Lng             float32   `xorm:"default 0"`
	Description     string    `xorm:"varchar(250)"`
	MetaKeywords    string    `xorm:"varchar(250)"`
	MetaDescription string    `xorm:"varchar(250)"`
	MetaTitle       string    `xorm:"varchar(250)"`
	DisplayOrder    int       `xorm:"default 100"`
	CreateTime      time.Time `xorm:"created "`
	UserId          int       `xorm:"default 0"`
	LogoImg         string    `xorm:"varchar(250)"`
	QQ              string    `xorm:"varchar(250)"`
}

func (s *Shop) TableName() string {
	return "j_shop"
}

func (s *Shop) Add() *Shop {
	if s.Id <= 0 {
		_, err := jtool.Engine.Insert(s)
		if err != nil {
			jtool.Log.Error("插入店铺信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(s.Id).Update(s)
		if err != nil {
			jtool.Log.Error("更新店铺信息失败 :", err)
		}
	}

	return s
}

func GetShop() []Shop {
	cs := make([]Shop, 0)
	err := jtool.Engine.Asc("Flg").Desc("Id").Find(&cs)
	if err != nil {
		jtool.Log.Error("获取商城中所有的店铺 :", err)
	}
	return cs
}

//更改店铺标识
func SetShopFlg(flg string, sId int) *Shop {
	s := &Shop{}
	s.Flg = flg
	s.Id = sId
	_, err := jtool.Engine.Id(s.Id).Update(s)
	if err != nil {
		jtool.Log.Error("修改店铺标识失败", err)

	}

	return s
}

func GetShopInHome(row int) []Shop {
	cs := make([]Shop, 0)
	err := jtool.Engine.Asc("Id").Limit(row, 0).Find(&cs)
	if err != nil {
		jtool.Log.Error("获取首页中显示的店铺 :", err)
	}
	return cs
}

func GetShopByUserId(uId int) *Shop {
	s := new(Shop)
	_, err := jtool.Engine.Desc("Id").Where("userId=?", uId).Get(s)
	if err != nil {
		jtool.Log.Error("获取店铺信息失败 :", err)
	}
	return s
}

func GetShopByShopId(Id int) *Shop {
	s := new(Shop)
	_, err := jtool.Engine.Where("Id=?", Id).Get(s)
	if err != nil {
		jtool.Log.Error("获取店铺信息失败 :", err)
	}
	return s
}

func GetNewProductsByShopId(shopId, row int) []Product {
	cs := make([]Product, 0)
	err := jtool.Engine.Where("ShopId=?", shopId).Desc("Id").Limit(row, 0).Find(&cs)
	if err != nil {
		jtool.Log.Error("根据店铺ID获取最新上传的产品 :", err)
	}
	return cs
}

func GetShopByUserIdFormCart(uId int) []Shop {
	sql := "select  * from xxsc.dbo.j_shop s  " +
		"where s.id in (SELECT  b.ShopId  FROM    xxsc.dbo.j_shop_product as b  " +
		"where b.Id in (   select a.ProductId from   " +
		"[xxsc].[dbo].[j_shop_cart] as a where a.UserId=?))"
	s := make([]Shop, 0)
	jtool.Engine.Sql(sql, uId).Find(&s)
	//jtool.Log.Error(s[0].Name )
	return s
}
