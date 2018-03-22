package shop

import (
	"qzyuanmu/jtool"
	"time"
)

type ProductImg struct {
	Id int `xorm:"pk autoincr"`
	//图片名称
	FileName string `xorm:"varchar(250)"`
	//后缀名
	Suffix string `xorm:"varchar(50)"`
	//图片使用类型，如shop 等
	ImgType string `xorm:"varchar(50)"`

	//原图
	OriImg string `xorm:"varchar(350)"`

	//800*800
	MaxImg string `xorm:"varchar(350)"`
	//350*350
	MidImg string `xorm:"varchar(350)"`
	//120*120
	MinImg     string    `xorm:"varchar(350)"`
	CreateTime time.Time `xorm:"created "`
	ProductId  int       `xorm:"default 0"`
}

func (p *ProductImg) TableName() string {
	return "j_shop_product_img"

}

func (p *ProductImg) Add() *ProductImg {
	if p.Id <= 0 {
		_, err := jtool.Engine.Insert(p)
		if err != nil {
			jtool.Log.Error("插入产品图片失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(p.Id).Update(p)
		if err != nil {
			jtool.Log.Error("更新产品图片失败 :", err)
		}
	}

	return p
}

func (p *ProductImg) Del(id int) bool {
	np := new(ProductImg)
	_, err := jtool.Engine.Where("id=?", id).Delete(np)
	if err != nil {
		jtool.Log.Error("删除产品图片失败 :", err)
		return false
	}
	return true
}

func (p *ProductImg) Get(id int) *ProductImg {
	np := new(ProductImg)
	_, err := jtool.Engine.Where("id=?", id).Get(np)
	if err != nil {
		jtool.Log.Error("获取店铺产品信息失败 :", err)
	}
	return np
}

func GetProductImgByProductId(pId int) []ProductImg {

	ps := make([]ProductImg, 0)
	err := jtool.Engine.Asc("CreateTime").Where("ProductId =?", pId).Find(&ps)
	if err != nil {
		jtool.Log.Error("根据产品Id获取产品图片列表失败 :", err)
	}
	return ps

}
