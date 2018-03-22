package red

import (
	"qzyuanmu/jtool"
	"time"
)

type SystemPram struct {
	Id         int       `xorm:"pk autoincr"`
	Title      string    `xorm:"varchar(250)"`
	Name       string    `xorm:"varchar(250)"`
	Value      string    `xorm:"varchar(250)"`
	UpdateTime time.Time `xorm:"created updated "`
	//1 无效 2 有效
	Flg    int `xorm:"default 2"`
	UserId int `xorm:"default 0"`
}

func (sys *SystemPram) TableName() string {
	return "j_systempram"

}

func GetSys() []*SystemPram {
	sys := make([]*SystemPram, 0)
	err := jtool.Engine.Find(&sys)
	if err != nil {
		jtool.Log.Error("获取系统参数失败", err)
	}
	return sys
}

func SetSys(id int, v string) bool {
	sys := new(SystemPram)
	sys.Id = id
	sys.Value = v
	_, err := jtool.Engine.Id(sys.Id).Cols("Value").Update(sys)
	if err != nil {
		jtool.Log.Error("获取系统参数失败1", err)
		return false
	}
	return true
}
