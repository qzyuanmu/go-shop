package user

import (
	"qzyuanmu/jtool"
)

type UserRole struct {
	Id        int    `xorm:"pk autoincr"`
	RoleName  string `xorm:"varchar(50)  " `
	RoleValue int    `xorm:"default 0"`
}

func (u *UserRole) TableName() string {
	return "j_users_role"
}

func GetUserROle() []UserRole {
	us := make([]UserRole, 0)
	err := jtool.Engine.Find(&us)
	if err != nil {
		jtool.Log.Error("获取用户角色错误", err)
	}
	return us
}
