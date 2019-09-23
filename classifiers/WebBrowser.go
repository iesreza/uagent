package classifiers

import (
	"strings"
)

type WebBrowser struct{}

func (WebBrowser) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "Opera Mini") {
		desc.Browser = "Opera Mini"
		return true
	}
	if (strings.Contains(ua, "Presto") && strings.Contains(ua, "Opera")) || strings.Contains(ua, "OPR/") {
		desc.Browser = "Opera"
		return true
	}
	if strings.Contains(ua, "Gecko") && strings.Contains(ua, "Firefox") {
		desc.Browser = "Firefox"
		return true
	}
	if strings.Contains(ua, ".NET") || strings.Contains(ua, "Trident") {
		desc.Browser = "Internet Explorer"
		return true
	}
	if strings.Contains(ua, "Edge") {
		desc.Browser = "Edge"
		return true
	}
	if strings.Contains(ua, "Kubuntu") {
		desc.Browser = "Konqueror"
		return true
	}
	if strings.Contains(ua, "Maxthon") {
		desc.Browser = "Maxthon"
		return true
	}
	if strings.Contains(ua, "Goanna") {
		desc.Browser = "PaleMoon"
		return true
	}
	if strings.Contains(ua, "Orca") {
		desc.Browser = "Orca"
		return true
	}

	if strings.Contains(ua, "Puffin") {
		desc.Browser = "Puffin"
		return true
	}
	if strings.Contains(ua, "QQBrowser") {
		desc.Browser = "QQBrowser"
		return true
	}

	if strings.Contains(ua, "UCBrowser") {
		desc.Browser = "UCBrowser"
		return true
	}
	if strings.Contains(ua, "YaBrowser") {
		desc.Browser = "Yandex"
		return true
	}
	if strings.Contains(ua, "Chimera") {
		desc.Browser = "Chimera"
		return true
	}
	if strings.Contains(ua, "Chrome") {
		desc.Browser = "Chrome"
		return true
	}
	if !strings.Contains(ua, "Chrome") && strings.Contains(ua, "Safari") {
		desc.Browser = "Safari"
		return true
	}

	return false
}
