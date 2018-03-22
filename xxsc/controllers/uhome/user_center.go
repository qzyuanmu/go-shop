package uhome

import (
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/city"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"
	"qzyuanmu/xxsc/models/market"

	"github.com/astaxie/beego"
)

type UserCenterController struct {
	beego.Controller
}

func (this *UserCenterController) Prepare() {
	id := models.SessionUserId(this.Controller)
	//fmt.Println("method"+this.Ctx.Request.Method)
	//jtool.Log.Error("method"+this.Ctx.Request.Method)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		s := shop.GetShopByUserId(u.Id)
		if s.Id >= 0 && s.Flg == "已审批" {
			u.ShopId = s.Id
		}
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

func (this *UserCenterController) Index() {

	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	this.Data["user"] = u
	this.Data["Title"] = "用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/index.html"
}

func (this *UserCenterController) Info() {

	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	this.Data["user"] = u

	ps := city.GetProvince()
	this.Data["provinces"] = ps

	cs := city.GetCity(u.ProvinceId)
	this.Data["citys"] = cs

	ds := city.GetDistrict(u.CityId)
	this.Data["districts"] = ds

	this.Data["Title"] = "修改资料-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/info.html"
}

func (this *UserCenterController) Apply() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)

	s := shop.GetShopByUserId(u.Id)
	this.Data["user"] = u
	this.Data["shop"] = s
	ps := city.GetProvince()
	this.Data["provinces"] = ps

	cs := city.GetCity(u.ProvinceId)
	this.Data["citys"] = cs

	ds := city.GetDistrict(u.CityId)
	this.Data["districts"] = ds

	ts := shop.GetShopType()
	this.Data["shopClass"] = ts

	this.Data["Title"] = "入驻申请-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/apply.html"
}

func (this *UserCenterController) Money() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	this.Data["user"] = u

	this.Data["Title"] = "资金管理-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/money.html"
}

func (this *UserCenterController) ApplyLog() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	tx := user.GetTiXianByUser(uId)
	u := user.GetUserById(uId)
	this.Data["user"] = u
	this.Data["tx"] = tx
	this.Data["Title"] = "提现记录-资金管理-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/applylog.html"

}

func (this *UserCenterController) Bank() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	this.Data["user"] = u
	this.Data["Title"] = "资金帐号-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/moneybank.html"
}

func (this *UserCenterController) Address() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	a := user.GetAddressByUserId(uId)
	this.Data["user"] = u
	this.Data["address"] = a
	ps := city.GetProvince()
	this.Data["provinces"] = ps

	cs := city.GetCity(a.ProvinceId)
	this.Data["citys"] = cs

	ds := city.GetDistrict(a.CityId)
	this.Data["districts"] = ds
	this.Data["Title"] = "收货地址-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/address.html"
}

func (this *UserCenterController) ChangePwd() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	this.Data["user"] = u
	this.Data["Title"] = "修改密码-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/changepwd.html"
}

func (this *UserCenterController) Cart() {
	uId := models.SessionUserId(this.Controller)
	cs := market.GetCart(uId)
	this.Data["carts"] = cs
	this.Data["Title"] = "购物车-芝麻乐购"
	//this.Layout = "tpl/market/share/_layoutshop.html"
	this.TplName = "tpl/market/uhome/cart.html"
}

func (this *UserCenterController) AddCart() {
	msg := jtool.NewMessage()
	uId := models.SessionUserId(this.Controller)
	if uId <= 0 {
		msg.Result = 3
		msg.Msg = "账户还未登录"
		this.Data["json"] = msg
		this.ServeJSON()
		return
	}
	pId, _ := this.GetInt("pId")
	tId := this.GetString("tId")
	q, _ := this.GetInt("q")
	cart := shop.AddCart(uId, pId, tId, q)

	if cart.Id <= 0 {
		msg.Result = 0
		msg.Msg = "添加购物车失败！"
	} else {
		msg.Result = 1
		msg.Msg = "添加购物车成功！"
	}

	this.Data["json"] = msg
	this.ServeJSON()

}

func (this *UserCenterController) EditCart() {
	msg := jtool.NewMessage()
	uId := models.SessionUserId(this.Controller)
	cId, _ := this.GetInt("cId")
	q, _ := this.GetInt("q")
	b := shop.EditCart(uId, cId, q)
	if b {
		msg.Result = 1

	} else {
		msg.Msg = "修改购物车失败，请联系管理员"
	}

	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *UserCenterController) Order() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	cs := market.GetCart(uId)
	if len(cs.Carts) <= 0 {
		url := this.URLFor("MkController.Login")
		this.Redirect(url, 302)
	}
	this.Data["carts"] = cs

	u := user.GetUserById(uId)
	a := user.GetAddressByUserId(uId)
	if a.Id <= 0 {
		url := this.URLFor("UserCenterController.Address")
		this.Redirect(url, 302)
	}
	this.Data["user"] = u
	this.Data["address"] = a
	ps := city.GetProvince()
	this.Data["provinces"] = ps
	citys := city.GetCity(a.ProvinceId)
	this.Data["citys"] = citys
	ds := city.GetDistrict(a.CityId)
	this.Data["districts"] = ds

	this.Data["Title"] = "提交订单-芝麻乐购"
	//this.Layout = "tpl/market/share/_layoutshop.html"
	this.TplName = "tpl/market/uhome/order.html"
}

func (this *UserCenterController) OrderList() {
	this.Data["hover"] = "shop_product"
	uId := models.SessionUserId(this.Controller)
	u := user.GetUserById(uId)
	ms := market.GetOrderListByUser(u.Id)
	this.Data["user"] = u
	this.Data["ms"] = ms
	this.Data["Title"] = "订单中心-用户中心-芝麻乐购"
	this.Layout = "tpl/market/share/_layout.html"
	this.TplName = "tpl/market/uhome/orderlist.html"
}
