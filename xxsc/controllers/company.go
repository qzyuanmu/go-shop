package controllers

import (
	"qzyuanmu/jtool/company"

	"github.com/astaxie/beego"

	"qzyuanmu/jtool"
	"qzyuanmu/jtool/template"
)

type ComController struct {
	beego.Controller
}

func (this *ComController) Index() {

	id, _ := this.Ctx.Input.Session(jtool.S_CompanyId).(int)
	com := company.GetCompany(id)
	if com.Url == "www.0571zm.com" {

		this.Redirect(this.URLFor("MkController.Index"), 302)
		return
	}
	tpls := template.Templates.Navs[com.Template]
	template.SetTpls(tpls, "首页")
	this.Data["com"] = com
	this.Data["tpl"] = tpls
	this.Data["Title"] = com.Name + "公司网站"
	this.Layout = "tpl/" + com.Template + "/share/_layout.html"
	this.TplName = "tpl/" + com.Template + "/index.html"
}

func (this *ComController) About() {
	id, _ := this.Ctx.Input.Session(jtool.S_CompanyId).(int)
	com := company.GetCompany(id)
	tpls := template.Templates.Navs[com.Template]
	template.SetTpls(tpls, "公司简介")
	this.Data["com"] = com
	this.Data["tpl"] = tpls
	this.Data["menu"] = company.GetMenuByComId(com.Id)

	id, _ = this.GetInt("id")
	if id <= 0 {
		id = 2
	}
	menu := company.GetMenu(id)
	this.Data["info"] = menu

	this.Data["Title"] = com.Name + "公司网站"
	this.Layout = "tpl/jusende/share/_layout.html"
	this.TplName = "tpl/jusende/about.html"
}

func (this *ComController) News() {

	this.Layout = "tpl/jusende/share/_layout.html"
	this.TplName = "tpl/jusende/index.html"
}

func (this *ComController) Product() {

	this.Layout = "tpl/jusende/share/_layout.html"
	this.TplName = "tpl/jusende/index.html"
}

func (this *ComController) Service() {

	this.Layout = "tpl/jusende/share/_layout.html"
	this.TplName = "tpl/jusende/index.html"
}

func (this *ComController) Help() {

	this.Layout = "tpl/jusende/share/_layout.html"
	this.TplName = "tpl/jusende/index.html"
}
