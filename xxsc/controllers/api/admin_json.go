package api

import (
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type JsonAdminController struct {
	beego.Controller
}

func (this *JsonAdminController) Prepare() {
	id := models.SessionUserId(this.Controller)
	u := new(user.User)
	if id > 0 {

		u = user.GetUserById(id)
		this.Data["u"] = u
	}
	if id <= 0 && u.UserRole < 5 {

		if this.Ctx.Request.Method == "GET" {
			url := this.URLFor("MkController.Login")
			this.Redirect(url, 302)
		} else {
			msg := jtool.NewMessage()
			msg.Result = 3
			this.Data["json"] = msg
			this.ServeJSON()
		}

	}
}
func (this *JsonAdminController) SetShopFlg() {
	sId, err := this.GetInt("sId")
	flg := this.GetString("Flg")
	msg := jtool.NewMessage()
	if err == nil || sId != 0 {
		s := shop.SetShopFlg(flg, sId)
		if s.Id > 0 {
			msg.Result = 1
			msg.Data = s
			msg.Msg = "修改成功！"

		}
	}
	this.Data["json"] = msg
	this.ServeJSON()
}
func (this *JsonAdminController) SetProductFlg() {
	sId, err := this.GetInt("sId")
	flg, _ := this.GetInt("Flg")
	msg := jtool.NewMessage()
	if err == nil || sId != 0 {
		s := shop.SetProductFlg(flg, sId)
		if s.Id > 0 {
			msg.Result = 1
			msg.Data = s
			msg.Msg = "修改成功！"

		}
	}
	this.Data["json"] = msg

	this.ServeJSON()
}
func (this *JsonAdminController) SetSys() {
	Id, err := this.GetInt("Id")
	v, _ := this.GetInt("v")
	v = 100 * v
	msg := jtool.NewMessage()
	if err == nil || Id != 0 {
		b := models.SetSys(Id, v)
		if b {
			msg.Result = 1
			msg.Msg = "修改成功！"

		}
	}
	this.Data["json"] = msg
	this.ServeJSON()
}
