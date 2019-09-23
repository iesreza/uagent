package classifiers

import (
	"regexp"
	"strings"
)

type Android struct {}
var regexAndroid = regexp.MustCompile(`Linux;.+(Android\s*[\d\.]+);\s*((\s*[a-zA-Z0-9\s-_]+))`)

func (Android)Eval(ua string,desc *Describer) bool{
	find := regexAndroid.FindStringSubmatch(ua)

	if len(find) > 3{
		parts := strings.Split(regexDataParts.FindString(ua),";")
		findModel := false
		for _,part := range parts{
			part = strings.TrimSpace(part)
			desc.Platform = Mobile
			desc.Device = Mobile
			if !findModel && strings.HasPrefix(part,"Android"){
				desc.OS = part
				findModel = true
				continue
			}
			if findModel{
				if len(part) == 5 && part[2] == '-'{
					continue
				}
				part = strings.Split(part,"Build")[0]

				desc.Model = regexModel.FindString(part)
			}
		}


		return true
	}
	return false
}