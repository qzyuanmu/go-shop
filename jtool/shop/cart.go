package shop

import (
	"errors"
	"qzyuanmu/jtool"
)

type Cart struct {
	Id              int    `xorm:"pk autoincr"`
	ProductId       int    `xorm:"default 0"`
	ProductTypeId   string `xorm:"varchar(250)"`
	ProductTypeName string `xorm:"varchar(250)"`
	ProductImg      string `xorm:"varchar(550)"`
	Quantity        int    `xorm:"default 0"`
	UserId          int    `xorm:"default 0"`
}

func (p *Cart) TableName() string {
	return "j_shop_cart"

}

func AddCart(uId int, pId int, tId string, q int) *Cart {
	c := GetCartByProduct(uId, pId, tId)
	err := errors.New("")
	if c.Id <= 0 {
		c.UserId = uId
		c.ProductId = pId
		c.ProductTypeId = tId
		c.Quantity = q
		if tId != "0" {
			str, img := GetProductTypeByTypeId(tId)
			c.ProductTypeName = str
			if len(img) > 0 {
				c.ProductImg = img
			}

		}
		_, err = jtool.Engine.Insert(c)
	} else {

		c.Quantity += q
		_, err = jtool.Engine.Id(c.Id).Update(c)
	}
	if err != nil {
		jtool.Log.Error("插入产品信息失败1 :", err)
	}
	return c
}

func GetCartByProduct(uId int, pId int, tId string) *Cart {
	c := new(Cart)
	err := errors.New("")
	if tId == "0" {
		_, err = jtool.Engine.Where("ProductId=? and UserId=?", pId, uId).Get(c)
	} else {
		_, err = jtool.Engine.
			Where("ProductId=? and UserId=? and ProductTypeId=?", pId, uId, tId).
			Get(c)
	}

	if err != nil {
		jtool.Log.Error("获取购物车中的产品信息失败 :", err)
	}
	return c
}

func GetCartByUser(uId int) []Cart {
	cs := make([]Cart, 0)
	err := jtool.Engine.Desc("Id").Where("UserId=?", uId).Find(&cs)
	if err != nil {
		jtool.Log.Error("根据用户获取购物车 :", err)
	}
	return cs
}
func GetCartById(Id int) *Cart {
	c := &(Cart{})
	_, err := jtool.Engine.Where("id=?", Id).Get(c)
	if err != nil {
		jtool.Log.Error("获取购物车失败", err)
	}
	return c
}

func EditCart(uId, cId, q int) bool {
	c := GetCartById(cId)
	if c.Id == 0 {
		return true
	}
	if c.UserId != uId {
		return false
	}
	if q <= 0 {
		_, err := jtool.Engine.Id(c.Id).Delete(c)
		if err != nil {
			return false
			jtool.Log.Error("移除购物车失败", err)
		}
	} else {
		c.Quantity = q
		_, err := jtool.Engine.Id(c.Id).Update(c)
		if err != nil {
			return false
			jtool.Log.Error("修改购物车失败", err)
		}
	}

	return true

}
