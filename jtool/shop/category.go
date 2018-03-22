package shop

import (
	"qzyuanmu/jtool"
	"time"
)

//公共分类
type Category struct {
	Id              int       `xorm:"pk autoincr"`
	Name            string    `xorm:"varchar(250)"`
	Description     string    `xorm:"varchar(250)"`
	MetaKeywords    string    `xorm:"varchar(250)"`
	MetaDescription string    `xorm:"varchar(250)"`
	MetaTitle       string    `xorm:"varchar(250)"`
	ParentId        int       `xorm:"default 0"`
	DisplayOrder    int       `xorm:"default 100"`
	CreateTime      time.Time `xorm:"created "`
	//1为启用，2为未启用
	Status int `xorm:"default 1"`
}

func (u *Category) TableName() string {
	return "j_category"
}

func (c *Category) Add() *Category {

	if c.Id <= 0 {
		_, err := jtool.Engine.Insert(c)
		if err != nil {
			jtool.Log.Error("插入公共分类失败 :", err)
		}
	} else {
		_, err := jtool.Engine.Id(c.Id).Update(c)
		if err != nil {
			jtool.Log.Error("更新公共分类失败 :", err)
		}
	}
	return c

}

func (c *Category) Get(id int) *Category {
	nc := new(Category)
	_, err := jtool.Engine.Where("id=?", id).Get(nc)
	if err != nil {
		jtool.Log.Error("获取商城分类信息失败 :", err)
	}
	return nc
}

func FindCategoryRoot() []Category {
	cs := make([]Category, 0)
	err := jtool.Engine.Asc("DisplayOrder").Where("ParentId=0").Find(&cs)
	if err != nil {
		jtool.Log.Error("获取公共分类的根目录失败 :", err)
	}
	return cs
}

//获取第三级公用分类
func GetThreeCategoryByRoot(rId int) []Category {
	cs := FindCategory()
	ts := make([]Category, 0)
	ths := make([]Category, 0)
	for _, c := range cs {
		if c.ParentId == rId {
			ts = append(ts, c)
		}
	}

	for _, c := range cs {

		for _, t := range ts {
			if t.Id == c.ParentId {
				ths = append(ths, c)
			}
		}
	}

	return ths
}

func GetCategoryForId(cs []Category) []int {
	ids := make([]int, 0)
	for _, c := range cs {
		ids = append(ids, c.Id)

	}
	return ids
}

func FindCategoryByParentId(pId int) []Category {
	cs := make([]Category, 0)
	err := jtool.Engine.Asc("DisplayOrder").Where("ParentId=?", pId).Find(&cs)
	if err != nil {
		jtool.Log.Error("根据父类Id获取公共分类的根目录失败 :", err)
	}
	return cs
}

///获取所有的公共分类
func FindCategory() []Category {
	cs := make([]Category, 0)
	err := jtool.Engine.Asc("DisplayOrder").Find(&cs)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}
	return cs

}
