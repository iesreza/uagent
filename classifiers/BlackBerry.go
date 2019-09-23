package classifiers

import (
	"regexp"
	"strings"
)

var regexBlackBerry = regexp.MustCompile(`(blackBerry\s*[a-zA-Z0-9]+)`)

type blackBerry struct{}

func (blackBerry) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "blackberry") {
		desc.Device = "BlackBerry"
		desc.OS = "BlackBerry"
		desc.Platform = Mobile
		find := regexBlackBerry.FindString(ua)
		desc.Model = find

		return true
	}
	return false
}
