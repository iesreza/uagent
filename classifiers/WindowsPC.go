package classifiers

import (
	"regexp"
	"strings"
)

type WinPC struct {}
var regexWinPC= regexp.MustCompile(`Windows\s*NT\s*([\d\.]+)`)
func (WinPC)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"Windows NT"){
		desc.Device = Computer
		desc.Platform = Desktop
		desc.OS = Windows
		find := regexWinPC.FindStringSubmatch(ua)

		if len(find) > 1{
			if strings.HasPrefix(find[1],"10.") {
				desc.Model = "Windows 10"
			}else {
				switch find[1] {
				case "5.0":
					desc.Model = "Windows 2000"
				case "5.01":
					desc.Model = "Windows 2000, Service Pack 1 (SP1)"
				case "5.1":
					desc.Model = "Windows XP"
				case "5.2":
					desc.Model = "Windows XP x64 Edition"
				case "6.0":
					desc.Model = "Windows Vista"
				case "6.1":
					desc.Model = "Windows 7"
				case "6.2":
					desc.Model = "Windows 8"
				case "6.3":
					desc.Model = "Windows 8.1"
				}
			}
		}

		if strings.Contains(ua,"Touch"){
			desc.Platform = Tablet
		}

		return true
	}
	return false
}