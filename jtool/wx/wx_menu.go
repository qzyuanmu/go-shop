package wx

import (
	"qzyuanmu/jtool"
	"time"
)

type WxMenu struct {
	Id int `xorm:"pk autoincr"`

	Name       string    `xorm:"varchar(250)" `
	Type       string    `xorm:"varchar(250)" `
	WxKey      string    `xorm:"varchar(250)" `
	Url        string    `xorm:"varchar(250)" `
	ParentId   int       `xorm:"default 0"`
	Sort       int       `xorm:"default 0"`
	WechatId   int       `xorm:"default 0"`
	UpdateTime time.Time `xorm:"created "`
	//1为启用，2为未启用
	Status int `xorm:"default 1"`
}

func (m *WxMenu) TableName() string {

	return "j_wx_menu"
}

//微信菜单的增加
func (s *WxMenu) Add() *WxMenu {
	if s.Id <= 0 {
		_, err := jtool.Engine.Insert(s)
		if err != nil {
			jtool.Log.Error("插入微信菜单信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(s.Id).Update(s)
		if err != nil {
			jtool.Log.Error("更新微信菜单信息失败 :", err)
		}
	}

	return s
}

func (m *WxMenu) Del(menu WxMenu) bool {
	session := jtool.Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Delete(&WxMenu{Id: menu.Id})
	if err != nil {
		session.Rollback()
		return false
	}
	_, err = session.Delete(&WxMenu{ParentId: menu.ParentId})
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

///获取微信号的自定义菜单
func FindWxMenuAll(wId int) []WxMenu {
	ws := make([]WxMenu, 0)
	err := jtool.Engine.Asc("Sort").Where("WechatId=?", wId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}
	ns := make([]WxMenu, 0)

	for _, m := range ws {
		if m.ParentId == 0 {
			ns = append(ns, m)
			ns = append(ns, GetMenuSub(ws, m.Id)...)
		}

	}
	return ns
}

///如果pId为0，则为一级菜单

func GetMenuSub(ms []WxMenu, pId int) []WxMenu {
	ns := make([]WxMenu, 0)
	for _, m := range ms {
		if m.ParentId == pId {
			ns = append(ns, m)
		}

	}
	return ns
}

///获取微信号的第一级自定义菜单
func FindWxMenuParent(wId int) []WxMenu {
	ws := make([]WxMenu, 0)
	err := jtool.Engine.Where("WechatId=? and ParentId=0", wId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信菜单信息失败 :", err)
	}
	return ws
}

func _(pId int) []WxMenu {
	ws := make([]WxMenu, 0)
	err := jtool.Engine.Where("ParentId =?", pId).Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信子菜单信息失败 :", err)
	}
	return ws

}
func GetWxMenuById(Id int) *WxMenu {
	wx := new(WxMenu)
	_, err := jtool.Engine.Where("id=?", Id).Get(wx)
	if err != nil {
		jtool.Log.Error("获取微信号信息失败 :", err)
	}
	return wx

}

func GetWxMenuByKey(wId int, k string) *WxMenu {

	wx := new(WxMenu)
	_, err := jtool.Engine.Where("wechatid=? and wxkey=?", wId, k).Get(wx)
	if err != nil {
		jtool.Log.Error("根据key获取微信号菜单信息失败 :", err)
	}
	return wx
}
