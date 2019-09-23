package classifiers

import "strings"

type iPod struct{}

func (iPod) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "(ipod") {
		desc.Device = "iPod"
		desc.Platform = MusicPlayer
		desc.OS = "iOS"
		return true
	}
	return false
}
