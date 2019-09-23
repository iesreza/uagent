package classifiers

import (
	"strings"
)

type Console struct{}

func (Console) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "playstation 4") {
		desc.Platform = GameConsole
		desc.Device = "PlayStation 4"
		desc.OS = "Orbis OS"
		return true
	}
	if strings.Contains(ua, "playstation 3") {
		desc.Platform = GameConsole
		desc.Device = "PlayStation 3"
		desc.OS = "Orbis OS"
		return true
	}
	if strings.Contains(ua, "playstation vita") {
		desc.Platform = GameConsole
		desc.Device = "PlayStation Vita"
		return true
	}

	if strings.Contains(ua, "xbox") {
		desc.Platform = GameConsole
		desc.Device = "Xbox"
		desc.OS = "Xbox OS"
		if strings.Contains(ua, "xbox One") {
			desc.Device = "Xbox One"
		}
		return true
	}
	if strings.Contains(ua, "nintendo wii") {
		desc.Platform = GameConsole
		desc.Device = "Nintendo Wii"
		desc.OS = "Wii U system software"
		return true
	}
	if strings.Contains(ua, "nexbox") {
		desc.Platform = GameConsole
		desc.Device = "NexBox"
		desc.OS = "Android"
		return true
	}
	return false
}
