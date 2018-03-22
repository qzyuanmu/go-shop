package api

import (
	"qzyuanmu/jtool/city"

	"github.com/astaxie/beego"
)

type CityController struct {
	beego.Controller
}

func (this *CityController) Province() {
	ps := city.GetProvince()
	this.Data["json"] = ps
	this.ServeJSON()
}

func (this *CityController) City() {
	cId, _ := this.GetInt("proId")
	ps := city.GetCity(cId)
	this.Data["json"] = ps
	this.ServeJSON()
}

func (this *CityController) District() {
	cId, _ := this.GetInt("cId")
	ps := city.GetDistrict(cId)
	this.Data["json"] = ps
	this.ServeJSON()
}
