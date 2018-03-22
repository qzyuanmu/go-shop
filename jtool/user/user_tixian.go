package user

import (
	"errors"
	"qzyuanmu/jtool"
	"time"
)

type TiXianModel struct {
	TiXian
	User
}

type TiXian struct {
	Id         int       `xorm:"pk autoincr"`
	UserName   string    `xorm:"varchar(250)"`
	RealName   string    `xorm:"varchar(250)"`
	Account    string    `xorm:"varchar(250)"`
	Bank       string    `xorm:"varchar(250)"`
	ZhiFuBao   string    `xorm:"varchar(250)"`
	CreateTime time.Time `xorm:"datetime created "`
	//提现时候的余额
	Money    int    `xorm:"default 0"`
	GetMoney int    `xorm:"default 0"`
	Memo     string `xorm:"varchar(250)"`
	//1 未审核  2 未支付 3 支付完成
	Flg      int    `xorm:"default 1"`
	FlgValue string `xorm:"varchar(250)"`
	UserId   int    `xorm:"default 0"`
	Md5      string `xorm:"varchar(250)"`
}

func (tx *TiXian) TableName() string {
	return "j_users_tixian"
}

func (tx *TiXian) Add(uId, money int, memo string) bool {

	err := errors.New("")
	session := jtool.Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	u := GetUserById(uId)

	if u.Money < money {
		return false
	}

	u.Money = u.Money - money

	_, err = session.Where("Id=?", uId).Cols("Money").Update(u)
	if err != nil {
		session.Rollback()
		jtool.Log.Error("提现出错1", err)
		return false
	}
	tx.Account = u.Account
	tx.Bank = u.Bank
	tx.Flg = 1
	tx.Money = u.Money
	tx.GetMoney = money
	tx.Memo = memo
	tx.RealName = u.RealName
	tx.ZhiFuBao = u.ZhiFuBao
	tx.FlgValue = "未审核"
	tx.UserName = u.UserName
	tx.UserId = u.Id
	_, err = session.Insert(tx)
	if err != nil {

		session.Rollback()
		jtool.Log.Error("提现出错2", err)
		return false
	}
	err = session.Commit()
	if err != nil {
		jtool.Log.Error("提现出错3", err)
		return false
	} else {
		return true
	}

}

func GetTiXian() []TiXian {
	ts := make([]TiXian, 0)
	t := time.Now()
	d := t.Format("2006-01-02")
	err := jtool.Engine.Desc("Id").
		Where("convert(varchar(10),CreateTime,120)=?", d).Find(&ts)
	if err != nil {
		jtool.Log.Error("获取会员提现失败 :", err)
	}

	return ts

}

func GetTiXianByUser(uId int) []TiXian {
	ts := make([]TiXian, 0)
	err := jtool.Engine.Desc("Id").Where("UserId=?", uId).Find(&ts)
	if err != nil {
		jtool.Log.Error("获取会员提现失败2 :", err)
	}
	return ts
}
