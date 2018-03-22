package admin

import (
	"fmt"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/news"
	"qzyuanmu/jtool/user"
	"qzyuanmu/jtool/wx"
	"qzyuanmu/xxsc/models"

	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (this *NewsController) Prepare() {
	id := models.SessionUserId(this.Controller)
	if id > 0 {
		u := new(user.User)
		u = user.GetUserById(id)
		this.Data["user"] = u
	}
}

func (this *NewsController) Index() {
	uid := SessionUserId(this.Controller)

	ns := news.GetNewByUserId(uid)
	this.Data["ms"] = ns
	this.Data["Title"] = "信息公告-基础内容"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/news/index.html"

}

///单个公告的内容编辑
func (this *NewsController) Add() {
	uid := SessionUserId(this.Controller)
	s := wx.GetWxByUserId(uid)
	ms := wx.FindWxMenuAll(s.Id)
	id, err := this.GetInt("id")
	if err != nil {
		id = 0
	}
	ns := news.GetNewById(id)
	this.Data["ns"] = ns
	this.Data["ms"] = ms
	this.Data["Title"] = "信息公告-基础内容"
	this.Layout = "ace/_Ace_index.html"
	this.TplName = "admin/news/add.html"

}

func (this *NewsController) Edit() {
	s := &news.News{}
	msg := jtool.NewMessage()

	if err := this.ParseForm(s); err != nil {
		msg.Result = 0
		msg.Data = s
		msg.Msg = fmt.Sprint(err)
	}
	if s.Id <= 0 {
		s.UserId = SessionUserId(this.Controller)
		s.Status = 1
	}

	s = s.Add()

	mId, err := this.GetInt("wxMenu")
	if err == nil {
		ns := new(wx.WxMenuNews)
		ns = wx.GetNewsByMenuId(mId)
		ns.MenuId = mId
		ns.NewsId = s.Id
		ns.Add()
	}

	if s.Id > 0 {

		msg.Result = 1
		msg.Data = s
		msg.Msg = "修改成功！"

	}
	this.Data["json"] = msg
	this.ServeJSON()
}
