package template

func GetJusendeNavs() []Tpl {
	tpls := make([]Tpl, 0)
	tpls = append(tpls, Tpl{Url: "/com/index", Class: "menu_style_homepage",
		Name: "首页", Current: "menu_current"})
	tpls = append(tpls, Tpl{Url: "/com/about", Class: "menu_style",
		Name: "公司简介", Current: "menu_current"})
	tpls = append(tpls, Tpl{Url: "/com/news", Class: "menu_style",
		Name: "新闻资讯", Current: "menu_current"})
	tpls = append(tpls, Tpl{Url: "/com/product", Class: "menu_style",
		Name: "产品中心", Current: "menu_current"})
	tpls = append(tpls, Tpl{Url: "/com/service", Class: "menu_style",
		Name: "客户留言", Current: "menu_current"})
	tpls = append(tpls, Tpl{Url: "/com/help", Class: "menu_style",
		Name: "联系我们", Current: "menu_current"})

	return tpls
}
