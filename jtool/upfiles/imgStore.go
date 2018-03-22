package upfiles

import (
	"qzyuanmu/jtool"
	"time"
)

type ImgStore struct {
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
	UserId     int       `xorm:"default 0"`
}

func (img *ImgStore) TableName() string {
	return "j_imgstore"
}

func (img *ImgStore) Add() *ImgStore {
	_, err := jtool.Engine.Insert(img)
	if err != nil {
		jtool.Log.Error("插入图片失败 :", err)
	}
	return img
}
