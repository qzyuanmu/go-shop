package wx

import (
	"qzyuanmu/jtool"
	"time"
)

type WxPublic struct {
	Id   int    `xorm:"pk autoincr"`
	Name string `xorm:"varchar(250)" `
	//微信原始ID
	OriId      string    `xorm:"varchar(250) unique" `
	AppId      string    `xorm:"varchar(250)" `
	Token      string    `xorm:"varchar(250)" `
	AesKey     string    `xorm:"varchar(250)" `
	UserId     int       `xorm:"default 0"`
	AppSecret  string    `xorm:"varchar(250)" `
	UpdateTime time.Time `xorm:"created "`

	//1为有效 2为无效
	Status int `xorm:"default 1"`
}

func (w *WxPublic) TableName() string {
	return "j_wx_public"
}

func (s *WxPublic) Add() *WxPublic {
	if s.Id <= 0 {
		_, err := jtool.Engine.Insert(s)
		if err != nil {
			jtool.Log.Error("插入微信号信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(s.Id).Update(s)
		if err != nil {
			jtool.Log.Error("更新微信号信息失败 :", err)
		}
	}

	return s
}

func FindWxAll() []WxPublic {
	ws := make([]WxPublic, 0)
	err := jtool.Engine.Find(&ws)
	if err != nil {
		jtool.Log.Error("获取微信号信息失败 :", err)
	}
	return ws
}

func GetWxByUserId(uId int) *WxPublic {
	wx := new(WxPublic)
	_, err := jtool.Engine.Where("userId=?", uId).Get(wx)
	if err != nil {
		jtool.Log.Error("获取微信号信息失败 :", err)
	}
	return wx

}
