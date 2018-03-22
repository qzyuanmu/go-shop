package user

import "time"

//用户资金变化记录表
type UserMoneyLog struct {
	Id int `xorm:"pk autoincr"`
	//金额
	Money int `xorm:"default 0"`
	//来源类型 红包 订单
	Type     string `xorm:"varchar(50)"`
	Quantity int    `xorm:"default 1"`
	//信息
	Info       string    `xorm:"varchar(500)"`
	CreateTime time.Time `xorm:"created updated "`
	UserId     int       `xorm:"default 0"`
}

func (rlog *UserMoneyLog) TableName() string {
	return "j_user_money_log"

}
