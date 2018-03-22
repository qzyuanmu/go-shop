package controllers

import (
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/city"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"
	"qzyuanmu/xxsc/models/market"

	"github.com/astaxie/beego"
)

type MkController struct {
	beego.Controller
}

func (this *MkController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["u"] = u
	}

}

func (this *MkController) Index() {
	cs := shop.FindCategory()
	this.Data["cs"] = cs
	//shops :=shop.GetShopInHome(24)
	//this.Data["shops"]=shops
	this.Data["fs"] = market.Home()
	this.Data["hover"] = "hover"
	this.Data["Title"] = "首页-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/index.html"
}
func (this *MkController) Shop() {

	this.Data["Title"] = "首页-芝麻乐购"
	this.Layout = "tpl/market/share/_layoutshop.html"
	this.TplName = "tpl/market/shop.html"
}

func (this *MkController) Product() {
	pId, err := this.GetInt("pId")
	if err != nil || pId <= 0 {
		this.Redirect(this.URLFor("MkController.Index"), 302)
		jtool.Log.Alert("用户非法登录商城产品界面：")
	}
	cs := shop.FindCategory()
	this.Data["cs"] = cs

	hp := market.NewHomeProduct(pId)
	this.Data["hp"] = hp
	s := shop.GetShopByShopId(hp.Product.ShopId)
	this.Data["shop"] = s

	np := shop.GetNewProductsByShopId(hp.Product.ShopId, 3)
	this.Data["np"] = np

	this.Data["hover"] = "shop_product"
	this.Data["Title"] = hp.Product.Name + "-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/product.html"
}

func (this *MkController) Login() {

	this.Data["Title"] = "用户登录-芝麻乐购"
	//this.Layout = "tpl/market/share/_layoutshop.html"
	this.TplName = "tpl/market/login.html"
}
func (this *MkController) Register() {
	ps := city.GetProvince()
	this.Data["ps"] = ps

	this.Data["Title"] = "用户注册-芝麻乐购"
	//this.Layout = "tpl/market/share/_layoutshop.html"
	this.TplName = "tpl/market/register.html"
}

func (this *MkController) LoginOut() {

	this.DelSession(jtool.S_UserName) //设置session用户名
	this.DelSession(jtool.S_UserId)
	this.Redirect(this.URLFor("MkController.Index"), 302)
}
