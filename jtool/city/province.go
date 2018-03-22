package city

import (
	"qzyuanmu/jtool"
)

type Province struct {
	Id   int    `xorm:"pk autoincr"`
	Name string `xorm:"varchar(250)"`
	Sort int    `xorm:"default 50"`
}

func (s *Province) TableName() string {
	return "pub_province"
}

func GetProvince() []Province {
	ps := make([]Province, 0)
	err := jtool.Engine.Asc("Sort").Asc("Name").Find(&ps)
	if err != nil {
		jtool.Log.Error("获取全国省份 :", err)
	}
	return ps
}
