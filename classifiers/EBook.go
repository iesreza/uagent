package classifiers

import (
	"strings"
)

type EBook struct {}

func (EBook)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"Silk") || strings.Contains(ua,"NOOK BNTV") || strings.Contains(ua,"Kindle"){
		desc.Device = EBookReader
		desc.Platform = EBookReader
		desc.OS = "Fire OS"
		return true
	}
	return false
}