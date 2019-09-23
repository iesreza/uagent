package classifiers

import "strings"

type iPhone struct {}

func (iPhone)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"(iPhone"){
		desc.Device = "iPhone"
		desc.Platform = Mobile
		desc.OS = "iOS"
		return true
	}
	return false
}