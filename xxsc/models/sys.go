package models

import (
	"fmt"
)

import (
	"qzyuanmu/jtool/red"
)

var Sys []*red.SystemPram

func init() {
	Sys = red.GetSys()
}

func GetSys(name string) red.SystemPram {

	for _, s := range Sys {
		if s.Name == name {
			return *s
		}
	}
	return *(new(red.SystemPram))
}

func SetSys(Id int, value int) bool {
	for _, s := range Sys {
		if s.Id == Id {
			v := float32(value)
			s.Value = fmt.Sprint(v / 100)
			if red.SetSys(s.Id, s.Value) {
				return true
			} else {
				return false
			}

		}
	}
	return false
}
