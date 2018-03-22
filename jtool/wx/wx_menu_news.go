package wx

import (
	"qzyuanmu/jtool"
	"time"
)

type WxMenuNews struct {
	Id         int       `xorm:"pk autoincr"`
	MenuId     int       `xorm:"varchar(250)" `
	NewsId     int       `xorm:"default 0"`
	UpdateTime time.Time `xorm:"created "`
}

func (this *WxMenuNews) TableName() string {

	return "j_wx_menu_news"
}

func GetNewsByMenuId(mId int) *WxMenuNews {

	n := new(WxMenuNews)
	_, err := jtool.Engine.Where("menuid=? ", mId).Get(n)
	if err != nil {
		jtool.Log.Error("获取菜单绑定信息失败 :", err)
	}

	return n
}

func (ns *WxMenuNews) Add() *WxMenuNews {
	if ns.Id <= 0 {
		_, err := jtool.Engine.Insert(ns)
		if err != nil {
			jtool.Log.Error("插入公告信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(ns.Id).Update(ns)

		if err != nil {
			jtool.Log.Error("更新公告信息失败 :", err)
		}
	}

	return ns
}
