package classifiers

import (
	"strings"
)


type Console struct {}

func (Console)Eval(ua string,desc *Describer) bool{
	if strings.Contains(ua,"PlayStation 4"){
		desc.Platform = GameConsole
		desc.Device = "PlayStation 4"
		desc.OS = "Orbis OS"
		return true
	}
	if strings.Contains(ua,"PlayStation 3"){
		desc.Platform = GameConsole
		desc.Device = "PlayStation 3"
		desc.OS = "Orbis OS"
		return true
	}
	if strings.Contains(ua,"PlayStation Vita"){
		desc.Platform = GameConsole
		desc.Device = "PlayStation Vita"
		return true
	}

	if strings.Contains(ua,"Xbox"){
		desc.Platform = GameConsole
		desc.Device = "Xbox"
		desc.OS = "Xbox OS"
		if strings.Contains(ua,"Xbox One"){
			desc.Device = "Xbox One"
		}
		return true
	}
	if strings.Contains(ua,"Nintendo Wii"){
		desc.Platform = GameConsole
		desc.Device = "Nintendo Wii"
		desc.OS = "Wii U system software"
		return true
	}
	if strings.Contains(ua,"NEXBOX"){
		desc.Platform = GameConsole
		desc.Device = "NexBox"
		desc.OS = "Android"
		return true
	}
	return false
}