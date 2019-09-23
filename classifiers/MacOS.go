package classifiers

import (
	"regexp"
	"strings"
)

type Mac struct {}
var regexMac = regexp.MustCompile(`(Mac\s*OS\s*X*\s*[\d.]*)`)
func (Mac)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"(Macintosh") || strings.Contains(ua,"PPC"){
		parts := regexMac.FindStringSubmatch(ua)
		desc.Device = Computer
		desc.Platform = Desktop
		if len(parts) == 2{
			desc.OS = parts[1]
		}else{
			desc.OS = "Mac OS"
		}

		return true
	}
	return false
}