package wx

import (
	"qzyuanmu/jtool"
)

func init() {
	jtool.Engine.Sync2(new(WxPublic), new(WxMenu), new(WxMenuNews))
}
