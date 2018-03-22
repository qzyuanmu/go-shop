package main

import (
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/company"
	"qzyuanmu/jtool/user"

	"qzyuanmu/xxsc/models"
	_ "qzyuanmu/xxsc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterAdmin = func(ctx *context.Context) {
	v, ok := ctx.Input.Session(jtool.S_UserName).(string)

	if len(v) <= 0 && !ok {
		ctx.Redirect(302, "/login/login")
	} else if user.GetUser(v).UserRole < 1 {

		ctx.Redirect(302, "/login/login")

	}

}

var FilterUser = func(ctx *context.Context) {
	v, ok := ctx.Input.Session(jtool.S_UserName).(string)

	if len(v) <= 0 && !ok {
		ctx.Redirect(302, "/mk/login")
	}
}

var FilterWeb = func(ctx *context.Context) {
	v, ok := ctx.Input.Session(jtool.S_CompanyId).(int)
	if v <= 0 && !ok {
		url := ctx.Input.Domain()
		com := company.GetCompanyByUrl(url)

		if com.Id > 0 {
			ctx.Input.CruSession.Set(jtool.S_CompanyId, com.Id)
			//ctx.Input.CruSession.Set(jtool.S_Template,com.Template)
		}
	}

}

func main() {
	//beego.InsertFilter("/", beego.BeforeRouter, FilterWeb)
	beego.InsertFilter("/com/*", beego.BeforeRouter, FilterWeb)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterAdmin)
	//beego.InsertFilter("/u/*", beego.BeforeRouter, FilterUser)
	beego.SetStaticPath("/upFiles", "upFiles")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AddFuncMap("fp", models.FormatPrice)
	beego.AddFuncMap("fr", models.FormatRed)
	beego.AddFuncMap("row", models.FormatRow)
	beego.AddFuncMap("dt", models.FormatDateTime)
	beego.Run()
}
