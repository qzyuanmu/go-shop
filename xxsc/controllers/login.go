package controllers

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/user"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	this.Data["Title"] = "后台管理登录"
	this.Data["Layout"] = "login-layout"
	this.Data["r"] = this.GetString("r")
	this.Layout = "ace/_Ace.html"
	this.TplName = "admin/login.html"

}

func (this *LoginController) ChkLogin() {
	u := new(user.User)
	msg := new(jtool.Message)

	if err := this.ParseForm(u); err != nil {
		msg.Result = 0
		msg.Msg = "数据转化失败！"
	}
	pwd := u.PassWord
	u = user.GetUser(u.UserName)

	if u.Id > 0 {

		if u.PassWord == jtool.Md5(pwd) {

			this.SetSession(jtool.S_UserName, u.UserName) //设置session用户名
			this.SetSession(jtool.S_UserId, u.Id)         //设置session用户Id

			fmt.Printf("password", jtool.Md5(pwd))
			msg.Result = 1
			//msg.Data = u
			msg.Msg = "验证通过"
		} else {
			msg.Result = 0
			//msg.Data = u
			msg.Msg = "帐号或密码错误！"
		}
	} else {
		msg.Result = 0

		msg.Msg = "帐号或密码错误！"
	}

	this.Data["json"] = msg
	this.ServeJSON()

}

func (this *LoginController) Register() {
	u := new(user.User)
	msg := new(jtool.Message)
	if err := this.ParseForm(u); err != nil {
		msg.Result = 0
		msg.Msg = "数据转化失败！"
	}

	if len(u.UserName) != 11 {
		msg.Result = 0
		msg.Msg = "请输入正确的手机号码"
		this.Data["json"] = msg
		this.ServeJSON()
		return
	}
	if len(u.Introducer) > 0 {

		intr := user.GetUser(u.Introducer)
		if intr.Id <= 0 {
			msg.Result = 0
			msg.Msg = "推荐人帐号不存在"
			this.Data["json"] = msg
			this.ServeJSON()
			return

		} else if intr.Flg != "商铺" && intr.Flg != "业务员" {
			msg.Result = 0
			msg.Msg = "推荐人帐号不能作为推荐人"
			this.Data["json"] = msg
			this.ServeJSON()
			return

		}
	}

	n := user.GetUser(u.UserName)
	if n.Id > 0 {
		msg.Result = 0
		msg.Msg = "用户已经注册！"
	} else {
		u = u.Add()
	}

	if u.Id > 0 {
		msg.Result = 1
		msg.Msg = "验证通过"

		this.SetSession(jtool.S_UserName, u.UserName) //设置session用户名
		this.SetSession(jtool.S_UserId, u.Id)         //设置session用户Id

	}

	this.Data["json"] = msg
	this.ServeJSON()

}

func (this *LoginController) ChkPhone() {
	phone := this.GetString("mobile")
	msg := true
	if len(phone) == 0 {

		msg = false

	} else {
		u := user.GetUser(phone)
		if u.Id <= 0 {
			msg = false
		}
	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *LoginController) PhoneRand() {
	phone := this.GetString("m")

	msg := new(jtool.Message)
	if len(phone) == 0 {
		msg.Result = 0
		msg.Msg = "请填写正确的手机号"
	} else {
		phone = user.Rand(4)
		msg.Result = 1
		msg.Data = phone
		msg.Msg = "验证通过"
	}
	this.Data["json"] = msg
	this.ServeJSON()
}
