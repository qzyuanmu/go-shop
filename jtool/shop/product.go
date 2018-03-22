package shop

import (
	"qzyuanmu/jtool"
	"time"
)

type Product struct {
	Id int `xorm:"pk autoincr"`

	Name string `xorm:"varchar(250)"`
	//商品编码
	Code             string `xorm:"varchar(250)"`
	CategoryId       int    `xorm:"default 0"`
	CategoryName     string `xorm:"varchar(250)"`
	ShopCategoryId   int    `xorm:"default 0"`
	ShopCategoryName string `xorm:"varchar(250)"`
	//店铺型号
	ShopType string `xorm:"varchar(250)"`
	//主图
	MainImg string `xorm:"varchar(550)"`
	//主图的缩率图
	MinImg string `xorm:"varchar(550)"`
	Info   string `xorm:"varchar(max)"`
	//红包比率
	RedLevel int `xorm:"default 0"`
	Quantity int `xorm:"default 0"`
	OriPrice int `xorm:"default 0"`
	Price    int `xorm:"default 0"`
	Zhim     int `xorm:"default 0"`
	//1 下架  2 上架
	Flg          int       `xorm:"default 1"`
	ShopId       int       `xorm:"default 0"`
	CreateTime   time.Time `xorm:"created "`
	UpdateTime   time.Time `xorm:"created updated "`
	CartQuantity int       `xorm:"default 0"`
	CartMoney    int       `xorm:"default 0"`
	CartZhim     int       `xorm:"default 0"`
	CartRed      int       `xorm:"default 0"`
	CartZRed     int       `xorm:"default 0"`
	CartTypeId   string    `xorm:"varchar(550)"`
	CartTypeName string    `xorm:"varchar(550)"`
	CartId       int
}

func (p *Product) TableName() string {
	return "j_shop_product"

}

func (p *Product) Add() *Product {
	if p.Id <= 0 {
		_, err := jtool.Engine.Insert(p)
		if err != nil {
			jtool.Log.Error("插入产品信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(p.Id).Update(p)
		if err != nil {
			jtool.Log.Error("更新产品信息失败 :", err)
		}
	}

	return p
}

func (p *Product) Get(id int) *Product {
	np := new(Product)
	_, err := jtool.Engine.Where("id=?", id).Get(np)
	if err != nil {
		jtool.Log.Error("获取店铺产品信息失败 :", err)
	}
	return np
}

func SetProductFlg(flg, pId int) *Product {
	p := new(Product)
	p.Id = pId
	p.Flg = flg
	_, err := jtool.Engine.Id(p.Id).Update(p)
	if err != nil {
		jtool.Log.Error("商品上下架失败", err)
	}
	return p

}

//根据第一大类获取产品
func GetProductByCategoryId(cId int) []Product {

	ps := make([]Product, 0)
	cs := GetThreeCategoryByRoot(cId)
	if len(cs) <= 0 {
		return ps
	}
	ids := GetCategoryForId(cs)

	err := jtool.Engine.Asc("UpdateTime").In("CategoryId", ids).Find(&ps)
	if err != nil {
		jtool.Log.Error("根据第三级分类Id获取店铺产品失败 :", err)
	}
	return ps

}

func GetProductByShopId(shopId int) []Product {

	ps := make([]Product, 0)
	err := jtool.Engine.Desc("UpdateTime").Desc("Id").Where("ShopId=?", shopId).Find(&ps)
	if err != nil {
		jtool.Log.Error("根据店铺Id获取店铺产品失败 :", err)
	}
	return ps

}

func GetProductByUserIdFormCart(uId int) []Product {
	sql := " SELECT  b.* FROM    xxsc.dbo.j_shop_product as b  " +
		"where b.Id in (   select a.ProductId from   " +
		"[xxsc].[dbo].[j_shop_cart] as a where a.UserId=?)"
	s := make([]Product, 0)
	jtool.Engine.Sql(sql, uId).Find(&s)
	return s
}
