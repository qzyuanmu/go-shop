package template

import (
	"strings"
)

type Tpl struct {
	Url string
	//默认的css类
	Class string
	Name  string
	//选定的css类
	Current string
}

type Tpls struct {
	Navs  map[string][]Tpl
	About map[string][]Tpl
}

func NewTpls() *Tpls {
	t := new(Tpls)
	t.About = make(map[string][]Tpl)
	t.Navs = make(map[string][]Tpl)
	return t
}

var Templates *Tpls

func init() {
	Templates = NewTpls()
	jusedeNav := GetJusendeNavs()
	Templates.Navs["jusende"] = jusedeNav
}

//设置选定内容的css
func SetTpls(tpls []Tpl, name string) {
	for i, t := range tpls {
		if t.Name == name {
			tpls[i].Class = tpls[i].Class + " " + tpls[i].Current

		} else {
			tpls[i].Class = strings.Replace(tpls[i].Class, tpls[i].Current, "", -1)
		}
	}
}
