package company

import (
	"qzyuanmu/jtool"
	"time"
)

type CompanyMenu struct {
	Id         int       `xorm:"pk autoincr"`
	Name       string    `xorm:"varchar(250)" `
	Info       string    `xorm:"varchar(max)" `
	InfoHtml   string    `xorm:"varchar(max)" `
	ParentId   int       `xorm:"default 0"`
	Sort       int       `xorm:"default 0"`
	CompanyId  int       `xorm:"default 0"`
	UpdateTime time.Time `xorm:"created "`
	//1为启用，2为未启用
	Status int `xorm:"default 1"`
}

func (m *CompanyMenu) TableName() string {

	return "j_company_menu"
}

//公司栏目菜单的增加
func (s *CompanyMenu) Add() *CompanyMenu {
	if s.Id <= 0 {
		_, err := jtool.Engine.Insert(s)
		if err != nil {
			jtool.Log.Error("插入公司栏目菜单信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(s.Id).Update(s)
		if err != nil {
			jtool.Log.Error("更新公司栏目菜单信息失败 :", err)
		}
	}

	return s
}

func (m *CompanyMenu) Del(menu CompanyMenu) bool {
	session := jtool.Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Delete(&CompanyMenu{Id: menu.Id})
	if err != nil {
		session.Rollback()
		return false
	}
	_, err = session.Delete(&CompanyMenu{ParentId: menu.ParentId})
	if err != nil {
		session.Rollback()
		return false
	}
	err = session.Commit()
	if err != nil {
		jtool.Log.Error("删除公司栏目菜单失败", err)
		return false
	} else {
		return true
	}

}

func GetMenu(id int) *CompanyMenu {
	m := new(CompanyMenu)
	_, err := jtool.Engine.Id(id).Get(m)
	if err != nil {
		jtool.Log.Error("根据id获取栏目菜单失败", err)
	}
	return m
}

func GetMenuByComId(cId int) []CompanyMenu {
	ms := make([]CompanyMenu, 0)
	err := jtool.Engine.Where("CompanyId=?", cId).Find(&ms)
	if err != nil {
		jtool.Log.Error("根据公司Id获取公司栏目菜单失败", err)
	}

	return ms

}

func GetMenuByParentId(pId int) []CompanyMenu {
	ms := make([]CompanyMenu, 0)
	err := jtool.Engine.Where("ParentId=?", pId).Find(&ms)
	if err != nil {
		jtool.Log.Error("根据父级Id获取公司栏目菜单失败", err)
	}

	return ms
}
