package user

import (
	"errors"
	"fmt"
	"qzyuanmu/jtool"
	"time"
)

type User struct {
	Id           int    `xorm:"pk autoincr"`
	UserName     string `xorm:"varchar(50) unique" `
	NickName     string `xorm:"varchar(50)  " `
	IdCard       string `xorm:"varchar(50)  "`
	PassWord     string `xorm:"varchar(50)"`
	Area         string `xorm:"varchar(50)"`
	Address      string `xorm:"varchar(250)"`
	UserImg      string `xorm:"varchar(250)"`
	ProvinceId   int    `xorm:"default 0"`
	ProvinceName string `xorm:"varchar(250)"`
	CityId       int    `xorm:"default 0"`
	CityName     string `xorm:"varchar(250)"`
	DistrictId   int    `xorm:"default 0"`
	DistrictName string `xorm:"varchar(250)"`
	RealName     string `xorm:"varchar(250)"`
	Account      string `xorm:"varchar(250)"`
	Bank         string `xorm:"varchar(250)"`
	ZhiFuBao     string `xorm:"varchar(250)"`
	Money        int    `xorm:"default 0"`
	RedBox       int    `xorm:"default 0"`
	JiFen        int    `xorm:"default 0"`
	//普通会员 业务员 商家
	Flg string `xorm:"varchar(250)"`
	//介绍人
	Introducer    string    `xorm:"varchar(50)"`
	Email         string    `xorm:"varchar(50)"`
	Mobile        string    `xorm:"varchar(50)  "`
	RegTime       time.Time `xorm:"created"`
	RegIp         string    `xorm:"varchar(50)"`
	LastLoginTime time.Time `xorm:"created "`
	LastLoginIp   string    `xorm:"varchar(50)"`
	UpdateTime    time.Time `xorm:"created "`
	Status        int       `xorm:"default 0"`
	ShopId        int       `xorm:"default 0"`
	//1 普通会员 3 店铺 5 公司 7 客服 9 财务 10系统管理员
	UserRole int `xorm:"default 1"`
}

func (u *User) TableName() string {
	return "j_users"
}

//新增和修改用户信息
func (u *User) Add() *User {
	err := errors.New("")
	if u.Id <= 0 {
		u.PassWord = jtool.Md5(u.PassWord)
		_, err = jtool.Engine.Insert(u)

	} else {
		_, err = jtool.Engine.Id(u.Id).Update(u)
	}
	if err != nil {
		u.Id = 0
		jtool.Log.Error("插入用户失败 :", err)
	}

	return u
}

func GetUserById(id int) *User {
	u := &User{}
	_, err := jtool.Engine.Id(id).Get(u)
	if err != nil {
		jtool.Log.Error("获取id"+fmt.Sprintf("%d", id)+"用户失败 :", err)
	}
	return u
}

//根据用户名查找用户
func GetUser(name string) *User {
	u := new(User)

	_, err := jtool.Engine.Where("UserName=?", name).Get(u)
	if err != nil {
		jtool.Log.Error("获取用户失败 :", err)
	}
	return u
}
func GetUserAll() []User {
	cs := make([]User, 0)
	err := jtool.Engine.Desc("Flg").Desc("Id").Find(&cs)
	if err != nil {
		jtool.Log.Error("获取商城中所有的店铺 :", err)
	}
	return cs
}
