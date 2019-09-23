package classifiers

import (
	"regexp"
	"strings"
)

type WinMobile struct{}

var regexWinMobile = regexp.MustCompile(`(microsoft|nokia)+;\s*([a-zA-Z0-9\s]+)`)

func (WinMobile) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "windows phone") || strings.Contains(ua, "windows mobile") {
		desc.Device = WindowsPhone
		desc.Platform = Mobile
		desc.OS = WindowsMobile
		find := regexWinMobile.FindStringSubmatch(ua)
		if len(find) > 2 {
			desc.Model = find[2]
		}

		return true
	}
	return false
}
