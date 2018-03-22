package upfiles

import (
	"qzyuanmu/jtool"
)

func init() {
	jtool.Engine.Sync2(new(ImgStore))
}
