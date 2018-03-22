package city

import (
	"qzyuanmu/jt.DB"
)

type City struct {
	Id       int    `xorm:"pk autoincr"`
	ProId    int    `xorm:"default 0"`
	Name     string `xorm:"varchar(250)"`
	Sort     int    `xorm:"default 50"`
	RegFlg   int    `xorm:"default 0"`
	RegTitle string `xorm:"varchar(250)"`
}

func (s *City) TableName() string {
	return "pub_city"
}

func GetCity(proId int) []City {
	cs := make([]City, 0)
	err := jt.DB.Engine.Asc("Sort").Asc("Name").Where("ProId=?", proId).Find(&cs)
	if err != nil {
		jt.DB.Log.Error("根据省份获取地市 :", err)
	}
	return cs
}
