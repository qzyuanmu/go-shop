package admin

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type ShopController struct {
	beego.Controller
}

func (this *ShopController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *ShopController) Index() {
	uId := SessionUserId(this.Controller)

	s := shop.GetShopByUserId(uId)
	this.Data["shop"] = s
	this.Data["Title"] = "店铺信息维护-后台管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/shop/index.html"

}

func (this *ShopController) Edit() {
	s := new(shop.Shop)
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
