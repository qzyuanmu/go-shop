package company

import (
	"qzyuanmu/jtool"
	"time"
)

type Company struct {
	Id              int       `xorm:"pk autoincr"`
	Name            string    `xorm:"varchar(250)"`
	Address         string    `xorm:"varchar(250)"`
	Contact         string    `xorm:"varchar(250)"`
	Phone           string    `xorm:"varchar(250)"`
	Lat             float32   `xorm:"default 0"`
	Lng             float32   `xorm:"default 0"`
	Icp             string    `xorm:"varchar(250)"`
	Description     string    `xorm:"varchar(250)"`
	MetaKeywords    string    `xorm:"varchar(250)"`
	MetaDescription string    `xorm:"varchar(250)"`
	MetaTitle       string    `xorm:"varchar(250)"`
	Sort            int       `xorm:"default 100"`
	Url             string    `xorm:"varchar(250)"`
	CreateTime      time.Time `xorm:"created "`
	UserId          int       `xorm:"default 0"`
	Template        string    `xorm:"varchar(250)"`
}

func (com *Company) TableName() string {

	return "j_company"
}

func (s *Company) Add() *Company {

	if s.Id <= 0 {
		_, err := jtool.Engine.Insert(s)
		if err != nil {
			jtool.Log.Error("插入公司信息失败 :", err)
		}
	} else {

		_, err := jtool.Engine.Id(s.Id).Update(s)
		if err != nil {
			jtool.Log.Error("更新公司信息失败 :", err)
		}
	}

	return s
}

func GetCompany(id int) *Company {
	com := new(Company)
	_, err := jtool.Engine.Id(id).Get(com)
	if err != nil {
		jtool.Log.Error("根据id获取公司信息失败 :", err)
	}
	return com

}

func GetCompanyByUser(uId int) *Company {
	com := new(Company)
	_, err := jtool.Engine.Where(" UserId=?", uId).Get(com)
	if err != nil {
		jtool.Log.Error("根据用户ID获取公司信息失败", err)
	}
	return com
}
func GetCompanyByUrl(url string) *Company {
	com := new(Company)
	_, err := jtool.Engine.Where(" Url=?", url).Get(com)
	if err != nil {
		jtool.Log.Error("根据url获取公司信息失败", err)
	}
	return com
}
