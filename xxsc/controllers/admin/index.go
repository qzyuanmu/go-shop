package admin

import (
	"qzyuanmu/jtool/user"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *IndexController) Index() {
	this.Data["Title"] = "后台管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/index.html"
}

func (this *IndexController) UpFiles() {

	this.Data["Title"] = "图片上传-后台管理"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/upfiles.html"

}
