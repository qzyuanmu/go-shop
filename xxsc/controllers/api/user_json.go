package api

import (
	"errors"
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"
	"qzyuanmu/xxsc/models/market"

	"github.com/astaxie/beego"
)

type JsonUserController struct {
	beego.Controller
}

func (this *JsonUserController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["u"] = u
	} else {

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
func (this *JsonUserController) Apply() {

	s := new(shop.Shop)
	msg := jtool.NewMessage()
	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	s.UserId = models.SessionUserId(this.Controller)
	s.Flg = "未审批"

	s = s.Add()
	if s.Id > 0 {
		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *JsonUserController) AddOrder() {
	err := errors.New("")
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	cs := market.GetCart(uId)
	msg := jtool.NewMessage()
	var z = 1
	z, err = this.GetInt("z")
	if len(cs.Carts) >= 1 || err != nil {
		o := new(shop.Order)
		if err := this.ParseForm(o); err != nil {
			msg.Result = 0
			msg.Data = o
			msg.Msg = "数据提交出错，请联系客服"
			this.Data["json"] = msg
			this.ServeJSON()
		}
		o.Ip = this.Ctx.Request.RemoteAddr

		if z == 2 && u.JiFen < cs.SumZhim {
			msg.Result = 0
			msg.Msg = "芝麻币不足！"
			this.Data["json"] = msg
			this.ServeJSON()
		}

		b := market.AddOrder(*o, u.Id, z)
		if b {
			msg.Result = 1
		}
	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *JsonUserController) TiXian() {
	msg := jtool.NewMessage()
	money, err := this.GetInt("m")
	if err != nil || money <= 0 {
		msg.Result = 0
		msg.Msg = "输入的金额不规范，请重新输入"
	}
	memo := this.GetString("memo")
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	if u.Money < money {
		msg.Result = 0
		msg.Msg = "余额不足，请重新输入"
	}

	tx := new(user.TiXian)
	b := tx.Add(u.Id, money, memo)
	if b {
		msg.Result = 1
		msg.Msg = "申请成功，请等待财务审核"
	}
	this.Data["json"] = msg
	this.ServeJSON()

}

func (this *JsonUserController) EditAddress() {

	s := new(user.UserAddress)
	msg := jtool.NewMessage()
	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	s.UserId = models.SessionUserId(this.Controller)
	s = s.Add()
	if s.Id > 0 {
		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *JsonUserController) EditUser() {

	s := new(user.User)
	msg := jtool.NewMessage()
	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	if true {
		s.Money = 0
		s.UserName = ""
		s.Email = ""
		s.JiFen = 0
		s.RedBox = 0
	}

	s.Id = models.SessionUserId(this.Controller)
	s = s.Add()
	if s.Id > 0 {
		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *JsonUserController) ChangePwd() {
	old := this.GetString("o")
	n := this.GetString("n")
	msg := jtool.NewMessage()
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	if u.PassWord != jtool.Md5(old) {
		msg.Result = 0
		msg.Msg = "旧密码输入错误"
	} else {
		u.PassWord = jtool.Md5(n)
		u = u.Add()
		if u.Id > 0 {
			msg.Result = 1
			msg.Msg = "修改密码成功"
		}

	}
	this.Data["json"] = msg
	this.ServeJSON()
}
