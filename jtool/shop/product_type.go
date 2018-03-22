package shop

import (
	"qzyuanmu/jtool"
)

type ProductType struct {
	Id        int    `xorm:"pk autoincr"`
	TypeId    int    `xorm:"default 1"`
	ProductId int    `xorm:"default 0"`
	Name      string `xorm:"varchar(250)"`
	Price     int    `xorm:"default 0"`
	//主图
	MainImg string `xorm:"varchar(550)"`
	//主图的缩率图
	MinImg string `xorm:"varchar(550)"`
	//商品编码

}

func (p *ProductType) TableName() string {
	return "j_shop_product_type"

}

func (p *ProductType) Add() *ProductType {
	if p.Id <= 0 {
		_, err := jtool.Engine.Insert(p)
		if err != nil {
			jtool.Log.Error("插入产品规格信息失败1 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(p.Id).Update(p)
		if err != nil {
			jtool.Log.Error("更新产品规格信息失败1 :", err)
		}
	}

	return p
}

func (p *ProductType) Del(id int) bool {
	np := new(ProductType)
	_, err := jtool.Engine.Where("id=?", id).Delete(np)
	if err != nil {
		jtool.Log.Error("删除产品规格失败 :", err)
		return false
	}
	return true
}

func (p *ProductType) Get(id int) *ProductType {
	np := new(ProductType)
	_, err := jtool.Engine.Where("id=?", id).Get(np)
	if err != nil {
		jtool.Log.Error("获取店铺产品信息失败1 :", err)
	}
	return np
}
func GetProductTypeByProductId(pId, tId int) []ProductType {

	ps := make([]ProductType, 0)
	session := jtool.Engine.Asc("Id").Where("ProductId=? ", pId)
	if tId != 0 {
		session.And(" TypeId=?", tId)
	}
	err := session.Find(&ps)
	if err != nil {
		jtool.Log.Error("根据产品Id获取产品规格失败 :", err)
	}
	return ps
}

func GetProductTypeByTypeId(tId string) (string, string) {
	str := ""
	img := ""
	if tId == "0" {
		return str, img
	}
	ts := make([]ProductType, 0)
	err := jtool.Engine.Where("id in(" + tId + ")").Find(&ts)
	if err != nil {
		jtool.Log.Error("获取产品规格信息失败 :", err)
	}
	for _, t := range ts {
		str += t.Name
		if len(t.MinImg) > 0 {
			img = t.MinImg
		}
	}
	return str, img
}
