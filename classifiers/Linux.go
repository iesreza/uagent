package classifiers

import (
	"strings"
)

type LinuxOS struct{}

func (LinuxOS) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "(x11") {
		desc.OS = Linux
		desc.Platform = Desktop
		desc.Device = Computer
		return true
	}
	return false
}
