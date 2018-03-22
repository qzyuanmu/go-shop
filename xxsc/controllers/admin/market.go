package admin

import (
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type MarketController struct {
	beego.Controller
}

func (this *MarketController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *MarketController) Sys() {

	sys := models.Sys
	this.Data["sys"] = sys
	this.Data["Title"] = "基础数据-商城管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/market/sys.html"

}

func (this *MarketController) TiXian() {

	tx := user.GetTiXian()
	this.Data["tx"] = tx
	this.Data["Title"] = "会员提现-商城管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/market/tixian.html"

}

func (this *MarketController) User() {
	u := user.GetUserAll()
	this.Data["users"] = u
	this.Data["Title"] = "会员管理-商城管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/market/muser.html"
}

func (this *MarketController) Shop() {
	s := shop.GetShop()
	this.Data["shops"] = s
	this.Data["Title"] = "商城店铺管理-商城管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/market/mshop.html"
}
