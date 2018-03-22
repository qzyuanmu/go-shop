package news

import (
	"qzyuanmu/jtool"
)

func init() {
	jtool.Engine.Sync2(new(News))
}
