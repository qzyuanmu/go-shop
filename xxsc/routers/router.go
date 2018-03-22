package routers

import (
	"qzyuanmu/xxsc/controllers"
	"qzyuanmu/xxsc/controllers/admin"
	"qzyuanmu/xxsc/controllers/api"
	"qzyuanmu/xxsc/controllers/uhome"

	"github.com/astaxie/beego"
	"github.com/beego/compress"
)

func init() {

	//APIS
	ns := beego.NewNamespace("/admin",

		beego.NSAutoRouter(&admin.IndexController{}),
		beego.NSAutoRouter(&admin.ShopController{}),

		beego.NSAutoRouter(&admin.NewsController{}),
		beego.NSAutoRouter(&admin.CompanyController{}),
		beego.NSAutoRouter(&admin.MkCategoryController{}),
		beego.NSAutoRouter(&admin.ShopCategoryController{}),
		beego.NSAutoRouter(&admin.ProductController{}),
		beego.NSAutoRouter(&admin.MarketController{}),
	)

	u := beego.NewNamespace("/u",
		beego.NSAutoRouter(&uhome.UserCenterController{}),
	)

	//APIS
	a := beego.NewNamespace("/api",
		beego.NSAutoRouter(&api.CityController{}),
		beego.NSAutoRouter(&api.JsonUserController{}),
		beego.NSAutoRouter(&api.JsonAdminController{}),
	)

	beego.AddNamespace(ns, u, a)
	beego.Router("/", &controllers.MkController{}, "get:Index")
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.UpImgController{})
	beego.AutoRouter(&controllers.ComController{})
	beego.AutoRouter(&controllers.MkController{})
	SettingCompress()

}
func SettingCompress() {
	// load json config file
	isProductMode := true
	setting, err := compress.LoadJsonConf("conf/compress.json", isProductMode, beego.URLFor("/"))
	if err != nil {
		beego.Error(err)
		return

	}

	// after use this api, can run command from shell.
	setting.RunCommand()

	if isProductMode {
		// if in product mode, can use this api auto compress files
		setting.RunCompress(true, false, true)
	}

	// add func to FuncMap for template use
	beego.AddFuncMap("compress_js", setting.Js.CompressJs)
	beego.AddFuncMap("compress_css", setting.Css.CompressCss)
}
