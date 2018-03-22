package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"qzyuanmu/jtool"
	"strconv"
)

func SyncUser() {

	//user.CreateUserTable()
}

func SessionUserId(c beego.Controller) int {

	uId := c.GetSession(jtool.S_UserId)
	id, err := strconv.Atoi(fmt.Sprint(uId))
	if err != nil {
		id = 0
	}
	return id
}

func SessionUserName(c beego.Controller) string {
	return fmt.Sprint(c.GetSession(jtool.S_UserName))
}
