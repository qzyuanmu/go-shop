package red

import "time"

//红包记录
type RedLog struct {
	Id int `xorm:"pk autoincr"`
	//红包数量
	Quantity int `xorm:"default 0"`
	//来源类型
	Type string `xorm:"varchar(50)"`
	//信息
	Info string `xorm:"varchar(500)"`

	CreateTime time.Time `xorm:"created updated "`
	UserId     int       `xorm:"default 0"`
}

func (rlog *RedLog) TableName() string {
	return "j_red_log"

}
