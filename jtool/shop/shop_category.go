package shop

import (
	"qzyuanmu/jtool"
	"time"
)

//各店铺自定义分类
type ShopCategory struct {
	Id              int    `xorm:"pk autoincr"`
	Name            string `xorm:"varchar(250)"`
	Description     string `xorm:"varchar(250)"`
	MetaKeywords    string `xorm:"varchar(250)"`
	MetaDescription string `xorm:"varchar(250)"`
	MetaTitle       string `xorm:"varchar(250)"`
	ParentId        int    `xorm:"default 0"`

	DisplayOrder int       `xorm:"default 100"`
	CreateTime   time.Time `xorm:"created "`
	ShopId       int
}

func (this ShopCategory) TableName() string {
	return "j_shop_category"
}

func (c *ShopCategory) Add() *ShopCategory {

	if c.Id <= 0 {
		_, err := jtool.Engine.Insert(c)
		if err != nil {
			jtool.Log.Error("插入店铺分类失败 :", err)
		}
	} else {
		_, err := jtool.Engine.Id(c.Id).Update(c)
		if err != nil {
			jtool.Log.Error("更新店铺分类失败 :", err)
		}
	}
	return c

}

func (c *ShopCategory) Get(id int) *ShopCategory {
	nc := new(ShopCategory)
	_, err := jtool.Engine.Where("id=?", id).Get(nc)
	if err != nil {
		jtool.Log.Error("获取店铺分类信息失败 :", err)
	}
	return nc
}

func (m *ShopCategory) Del(menu ShopCategory) bool {
	session := jtool.Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Delete(&ShopCategory{Id: menu.Id})
	if err != nil {
		session.Rollback()
		return false
	}
	_, err = session.Delete(&ShopCategory{ParentId: menu.ParentId})
	if err != nil {
		session.Rollback()
		return false
	}
	err = session.Commit()
	if err != nil {
		return false
	} else {
		return true
	}

}

func FindShopCategoryAll(wId int) []ShopCategory {
	ws := make([]ShopCategory, 0)
	err := jtool.Engine.Asc("DisplayOrder").Where("ShopId=?", wId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}
	ns := make([]ShopCategory, 0)

	for _, m := range ws {
		if m.ParentId == 0 {
			ns = append(ns, m)
			ns = append(ns, GetShopCategorySub(ws, m.Id)...)
		}

	}
	return ns
}

///如果pId为0，则为一级菜单

func GetShopCategorySub(ms []ShopCategory, pId int) []ShopCategory {
	ns := make([]ShopCategory, 0)
	for _, m := range ms {
		if m.ParentId == pId {
			ns = append(ns, m)
		}

	}
	return ns
}

func FindShopCategoryParent(wId int) []ShopCategory {
	ws := make([]ShopCategory, 0)
	err := jtool.Engine.Where("ShopId=? and ParentId=0", wId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}
	return ws
}

func FindShopCategorySub(pId int) []ShopCategory {
	ws := make([]ShopCategory, 0)
	err := jtool.Engine.Where("ParentId =?", pId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信子菜单信息失败 :", err)
	}
	return ws

}
func GetShopCategoryById(Id int) *ShopCategory {
	wx := new(ShopCategory)
	_, err := jtool.Engine.Where("id=?", Id).Get(wx)
	if err != nil {
		jtool.Log.Error("获取微信号信息失败 :", err)
	}
	return wx

}
