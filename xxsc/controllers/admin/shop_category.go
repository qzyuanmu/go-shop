package admin

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type ShopCategoryController struct {
	beego.Controller
}

func (this *ShopCategoryController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *ShopCategoryController) Category() {
	uId := SessionUserId(this.Controller)
	s := shop.GetShopByUserId(uId)
	this.Data["shop"] = s
	this.Data["ps"] = shop.FindShopCategoryParent(s.Id)
	this.Data["ms"] = shop.FindShopCategoryAll(s.Id)
	this.Data["Title"] = "店铺分类管理-后台管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/shop/category.html"

}
func (this *ShopCategoryController) GetCategory() {
	id, _ := this.GetInt("id")
	cs := shop.FindShopCategorySub(id)
	this.Data["json"] = cs
	this.ServeJSON()
}

func (this *ShopCategoryController) GetMenu() {
	id, _ := this.GetInt("id")
	c := new(shop.ShopCategory)
	m := c.Get(id)
	this.Data["json"] = m
	this.ServeJSON()
}

func (this *ShopCategoryController) Edit() {
	s := new(shop.ShopCategory)
	msg := jtool.NewMessage()

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	if s.Id <= 0 {
		s.ShopId = shop.GetShopByUserId(SessionUserId(this.Controller)).Id

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
