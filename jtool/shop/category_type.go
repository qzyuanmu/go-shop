package shop

import (
	"qzyuanmu/jtool"
	"time"
)

//分类所属规格
type CategoryType struct {
	Id           int       `xorm:"pk autoincr"`
	Name         string    `xorm:"varchar(250)"`
	Description  string    `xorm:"varchar(250)"`
	CategoryId   int       `xorm:"default 0"`
	DisplayOrder int       `xorm:"default 100"`
	CreateTime   time.Time `xorm:"created "`
}

func (u *CategoryType) TableName() string {
	return "j_category_type"
}

func (c *CategoryType) Add() *CategoryType {

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

func (c *CategoryType) Get(id int) *CategoryType {
	nc := new(CategoryType)
	_, err := jtool.Engine.Where("id=?", id).Get(nc)
	if err != nil {
		jtool.Log.Error("获取商城分类信息失败 :", err)
	}
	return nc
}

func FindCategoryType(cId int) []CategoryType {
	cs := make([]CategoryType, 0)
	err := jtool.Engine.Asc("DisplayOrder").Where("CategoryId=?", cId).Find(&cs)
	if err != nil {
		jtool.Log.Error("获取公共分类的根目录失败 :", err)
	}
	return cs
}
