package admin

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type MkCategoryController struct {
	beego.Controller
}

func (this *MkCategoryController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *MkCategoryController) Index() {
	cs := shop.FindCategoryByParentId(0)
	this.Data["cs"] = cs
	this.Data["Title"] = "商城分类管理-商城分类-商城管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/market/category.html"
}

func (this *MkCategoryController) GetCategory() {
	pid, err := this.GetInt("id")
	if err != nil {
		pid = 0
	}
	cs := shop.FindCategoryByParentId(pid)
	this.Data["json"] = cs
	this.ServeJSON()
}

func (this *MkCategoryController) GetId() {
	id, err := this.GetInt("id")
	if err != nil {
		id = 0
	}
	c := new(shop.Category)
	c = c.Get(id)
	this.Data["json"] = c
	this.ServeJSON()
}

func (this *MkCategoryController) Edit() {
	s := new(shop.Category)
	msg := jtool.NewMessage()

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}

	s = s.Add()
	if s.Id > 0 {
		//更新微信服务

		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}
