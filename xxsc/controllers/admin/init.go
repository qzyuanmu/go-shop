package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"qzyuanmu/jtool"
	"strconv"
)

const (
	RootUrl = "http://test.qzcheng.com/"
)

func SessionUserId(c beego.Controller) int {

	uId := c.GetSession(jtool.S_UserId)
	id, _ := strconv.Atoi(fmt.Sprint(uId))
	return id
}

func SessionUserName(c beego.Controller) string {
	return fmt.Sprint(c.GetSession(jtool.S_UserName))
}
