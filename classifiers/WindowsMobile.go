package classifiers

import (
	"regexp"
	"strings"
)

type WinMobile struct {}
var regexWinMobile = regexp.MustCompile(`(Microsoft|NOKIA)+;\s*([a-zA-Z0-9\s]+)`)
func (WinMobile)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"Windows Phone") || strings.Contains(ua,"Windows Mobile"){
		desc.Device = WindowsPhone
		desc.Platform = Mobile
		desc.OS = WindowsMobile
		find := regexWinMobile.FindStringSubmatch(ua)
		if len(find) > 2{
			desc.Model = find[2]
		}

		return true
	}
	return false
}