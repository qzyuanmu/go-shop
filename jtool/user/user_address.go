package user

import (
	"errors"
	"qzyuanmu/jtool"
)

type UserAddress struct {
	Id           int    `xorm:"pk autoincr"`
	Area         string `xorm:"varchar(50)"`
	Address      string `xorm:"varchar(250)"`
	ProvinceId   int    `xorm:"default 0"`
	ProvinceName string `xorm:"varchar(250)"`
	CityId       int    `xorm:"default 0"`
	CityName     string `xorm:"varchar(250)"`
	DistrictId   int    `xorm:"default 0"`
	DistrictName string `xorm:"varchar(250)"`
	ZipCode      string `xorm:"varchar(50)  "`
	RealName     string `xorm:"varchar(250)"`
	Email        string `xorm:"varchar(50)"`
	Mobile       string `xorm:"varchar(50)  "`
	UserId       int    `xorm:"default 0"`
}

func (u *UserAddress) TableName() string {
	return "j_users_address"
}

func (u *UserAddress) Add() *UserAddress {
	err := errors.New("")
	if u.Id <= 0 {
		_, err = jtool.Engine.Insert(u)

	} else {
		_, err = jtool.Engine.Id(u.Id).Update(u)
	}
	if err != nil {
		u.Id = 0
		jtool.Log.Error("插入用户地址失败 :", err)
	}

	return u
}

func GetAddressByUserId(uId int) *UserAddress {
	u := &UserAddress{}
	_, err := jtool.Engine.Where("UserId=?", uId).Get(u)
	if err != nil {
		jtool.Log.Error("获取用户地址失败", err)
	}
	return u
}
