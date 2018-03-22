package jtool

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

var Log = logs.NewLogger(10000)

const (
	Key     = "22sddfdkjskjl122``22200988.%%^^&ww2$#@@"
	RootUrl = "http://wx.test.com"
)

func init() {
 

	var err error
	Engine, err = xorm.
		NewEngine("mssql",
			"server=.;database=xxsc;user id=sa;password=admin;"+
				"port=2438;MultipleActiveResultSets=true ;Max Pool Size=150;"+
				"Connect Timeout=500;")
	Engine.ShowSQL(true)
	//Engine.TZLocation,_ =time.LoadLocation("Asia/Shanghai")
	//location, _ := time.LoadLocation("UTC")
	//Engine.TZLocation = location
	Engine.SetMapper(core.SameMapper{})
	if err != nil {
		Log.Error(" Sql错误：", err)
	}

}
