package city

import (
	"qzyuanmu/jt.DB"
)

type District struct {
	Id     int    `xorm:"pk autoincr"`
	Name   string `xorm:"varchar(250)"`
	CityId int    `xorm:"default 0"`
	Sort   int    `xorm:"default 50"`
}

func (s *District) TableName() string {
	return "pub_district"
}
func GetDistrict(cityId int) []District {
	cs := make([]District, 0)
	err := jt.DB.Engine.Asc("Sort").Asc("Name").Where("CityId=?", cityId).Find(&cs)
	if err != nil {
		jt.DB.Log.Error("根据市获取区 :", err)
	}
	return cs
}
