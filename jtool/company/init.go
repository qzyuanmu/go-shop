package company

import (
	"qzyuanmu/jtool"
)

func init() {
	jtool.Engine.Sync2(new(Company), new(CompanyMenu))
}
