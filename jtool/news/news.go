package news

import (
	"fmt"
	"qzyuanmu/jtool"
	"time"
)

type News struct {
	Id         int       `xorm:"pk autoincr"`
	Title      string    `xorm:"varchar(250)" `
	Type       string    `xorm:"varchar(250)" `
	Info       string    `xorm:"ntext" `
	InfoHtml   string    `xorm:"ntext" `
	Sort       int       `xorm:"default 20"`
	UserId     int       `xorm:"default 0"`
	UpdateTime time.Time `xorm:"created "`
	//1为启用，2为未启用
	Status int `xorm:"default 1"`
}

func (s *News) TableName() string {

	return "j_news"
}

func (ns *News) Add() *News {
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

func GetNewById(id int) *News {
	n := &News{}

	if id <= 0 {
		return n
	}
	_, err := jtool.Engine.Where("Id=?", id).Get(n)
	if err != nil {
		jtool.Log.Error("获取id"+fmt.Sprintf("%d", id)+"公告失败 :", err)
	}
	return n

}

func GetNewByUserId(uid int) []News {
	n := make([]News, 0)

	if uid <= 0 {
		return n
	}
	err := jtool.Engine.Where("UserId=?", uid).Find(&n)
	if err != nil {
		jtool.Log.Error("获取id"+fmt.Sprintf("%d", uid)+"公告失败 :", err)
	}
	return n

}

func GetNewByKey(key string) *News {
	n := &News{}

	_, err := jtool.Engine.Asc("Id").Where("WxKey=?", key).Get(n)
	if err != nil {
		jtool.Log.Error("获取id"+key+"公告失败 :", err)
	}
	return n

}
