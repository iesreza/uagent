package classifiers

import (
	"strings"
)

type EBook struct{}

func (EBook) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "silk") || strings.Contains(ua, "nook bntv") || strings.Contains(ua, "kindle") {
		desc.Device = EBookReader
		desc.Platform = EBookReader
		desc.OS = "Fire OS"
		return true
	}
	return false
}
