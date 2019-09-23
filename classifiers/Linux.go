package classifiers

import (
	"regexp"
	"strings"
)

type LinuxOS struct {}
var regexLinuxOS= regexp.MustCompile(`(Mac\s*OS\s*X*\s*[\d.]*)`)
func (LinuxOS)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"(X11"){
		desc.OS = Linux
		desc.Platform = Desktop
		desc.Device = Computer
		return true
	}
	return false
}