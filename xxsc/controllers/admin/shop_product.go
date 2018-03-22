package admin

import (
	"errors"
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (this *ProductController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *ProductController) Index() {
	uId := SessionUserId(this.Controller)
	s := shop.GetShopByUserId(uId)
	cs := shop.FindShopCategoryAll(s.Id)
	this.Data["ps"] = shop.GetProductByShopId(s.Id)
	this.Data["cs"] = cs
	this.Data["Title"] = "产品管理-店铺管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/shop/product.html"
}

func (this *ProductController) Add() {
	uId := SessionUserId(this.Controller)
	id, err := this.GetInt("id")
	p := new(shop.Product)
	if err == nil && id > 0 {
		p = new(shop.Product).Get(id)

	}
	s := shop.GetShopByUserId(uId)
	shopcs := shop.FindShopCategoryAll(s.Id)
	cs := shop.FindCategoryByParentId(0)
	pimg := shop.GetProductImgByProductId(id)

	types := shop.GetProductTypeByProductId(id, 1)
	type1s := shop.GetProductTypeByProductId(id, 2)

	this.Data["type1s"] = type1s
	this.Data["types"] = types
	this.Data["imgs"] = pimg
	this.Data["ms"] = shopcs
	this.Data["cs"] = cs
	this.Data["p"] = p
	this.Data["Title"] = "产品管理-店铺管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/shop/product_add.html"
}

func (this *ProductController) GetCategory() {
	pid, err := this.GetInt("id")
	if err != nil {
		pid = 0
	}
	cs := shop.FindCategoryByParentId(pid)
	this.Data["json"] = cs
	this.ServeJSON()
}

func (this *ProductController) GetId() {
	id, err := this.GetInt("id")
	if err != nil {
		id = 0
	}
	c := new(shop.Category)
	c = c.Get(id)
	this.Data["json"] = c
	this.ServeJSON()
}

func (this *ProductController) DelImg() {
	//pid,_ :=this.GetInt("pid")
	id, _ := this.GetInt("id")
	msg := jtool.NewMessage()
	pimg := new(shop.ProductImg)
	b := pimg.Del(id)

	if b {
		//更新微信服务

		msg.Result = 1

		msg.Msg = "图片删除成功"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *ProductController) DelType() {
	//pid,_ :=this.GetInt("pid")
	id, _ := this.GetInt("id")

	msg := jtool.NewMessage()
	pimg := new(shop.ProductType)
	b := pimg.Del(id)

	if b {
		//更新微信服务

		msg.Result = 1

		msg.Msg = "图片删除成功"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *ProductController) SetImg() {
	pid, _ := this.GetInt("pid")
	id, _ := this.GetInt("id")
	s := new(shop.Product)
	msg := jtool.NewMessage()

	s.Id = pid
	img := new(shop.ProductImg)
	img = img.Get(id)
	s.MainImg = img.MaxImg
	s.MinImg = img.MidImg

	s = s.Add()
	if s.Id > 0 {
		//更新微信服务

		msg.Result = 1
		msg.Data = s
		msg.Msg = "首图设置成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *ProductController) Edit() {
	s := new(shop.Product)
	msg := jtool.NewMessage()
	uId := SessionUserId(this.Controller)
	shop := shop.GetShopByUserId(uId)

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}

	if s.Id <= 0 {
		s.ShopId = shop.Id
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

func (this *ProductController) TypeEdit() {
	s := new(shop.ProductType)
	err := errors.New("")
	msg := jtool.NewMessage()
	s.MainImg = this.GetString("img2")
	s.MinImg = this.GetString("img")
	s.Name = this.GetString("t")

	s.TypeId, err = this.GetInt("tId")
	s.ProductId, err = this.GetInt("pid")
	if err != nil || s.ProductId == 0 {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	} else {

		s = s.Add()
		if s.Id > 0 {

			msg.Result = 1
			msg.Data = s
			msg.Msg = "修改成功！"

		}
	}

	this.Data["json"] = msg
	this.ServeJSON()
}
