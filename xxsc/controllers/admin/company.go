package admin

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/company"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type CompanyController struct {
	beego.Controller
}

func (this *CompanyController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *CompanyController) Add() {
	uId := SessionUserId(this.Controller)

	s := company.GetCompanyByUser(uId)
	this.Data["c"] = s
	this.Data["Title"] = "公司信息维护-后台管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/company/add.html"

}

func (this *CompanyController) Edit() {
	s := new(company.Company)
	msg := jtool.NewMessage()

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	if s.Id <= 0 {
		s.UserId = SessionUserId(this.Controller)

	}
	s = s.Add()
	if s.Id > 0 {

		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *CompanyController) Menu() {
	id := SessionUserId(this.Controller)
	//获取微信号的ID
	com := company.GetCompanyByUser(id)
	ms := company.GetMenuByComId(com.Id)
	this.Data["ms"] = ms
	this.Data["Title"] = "栏目管理-公司栏目-公司管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/company/menu.html"

}

func (this *CompanyController) MenuAdd() {
	id := SessionUserId(this.Controller)
	//获取公司的ID
	com := company.GetCompanyByUser(id)
	if com.Id <= 0 {
		this.Redirect("/admin/company/add", 302)
	}

	id, err := this.GetInt("id")
	if err != nil {
		id = 0
	}
	menu := company.GetMenu(id)
	ms := company.GetMenuByComId(com.Id)
	this.Data["menu"] = menu
	this.Data["ms"] = ms
	this.Data["Title"] = "栏目编辑-公司栏目-公司管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/company/menuadd.html"

}

func (this *CompanyController) MenuEdit() {
	id := SessionUserId(this.Controller)
	//获取公司的ID
	com := company.GetCompanyByUser(id)

	s := &company.CompanyMenu{}
	msg := jtool.NewMessage()

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	if s.Id <= 0 {
		s.CompanyId = com.Id

		s.Status = 1
	}

	s = s.Add()

	if s.Id > 0 {

		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}
