package classifiers

import "strings"

type iPad struct{}

func (iPad) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "(ipad") {
		desc.Device = "iPad"
		desc.Platform = "Tablet"
		desc.OS = "iOS"
		return true
	}
	return false
}
