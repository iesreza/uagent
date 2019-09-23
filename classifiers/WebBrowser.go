package classifiers

import (
	"strings"
)

type WebBrowser struct{}

func (WebBrowser) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "opera mini") {
		desc.Browser = "Opera Mini"
		return true
	}
	if (strings.Contains(ua, "presto") && strings.Contains(ua, "opera")) || strings.Contains(ua, "opr/") {
		desc.Browser = "Opera"
		return true
	}
	if strings.Contains(ua, "gecko") && strings.Contains(ua, "firefox") {
		desc.Browser = "Firefox"
		return true
	}
	if strings.Contains(ua, ".net") || strings.Contains(ua, "trident") {
		desc.Browser = "Internet Explorer"
		return true
	}
	if strings.Contains(ua, "edge") {
		desc.Browser = "Edge"
		return true
	}
	if strings.Contains(ua, "kubuntu") {
		desc.Browser = "Konqueror"
		return true
	}
	if strings.Contains(ua, "maxthon") {
		desc.Browser = "Maxthon"
		return true
	}
	if strings.Contains(ua, "goanna") {
		desc.Browser = "PaleMoon"
		return true
	}
	if strings.Contains(ua, "orca") {
		desc.Browser = "Orca"
		return true
	}

	if strings.Contains(ua, "puffin") {
		desc.Browser = "Puffin"
		return true
	}
	if strings.Contains(ua, "qqbrowser") {
		desc.Browser = "QQBrowser"
		return true
	}

	if strings.Contains(ua, "ucbrowser") {
		desc.Browser = "UCBrowser"
		return true
	}
	if strings.Contains(ua, "yabrowser") {
		desc.Browser = "Yandex"
		return true
	}
	if strings.Contains(ua, "chimera") {
		desc.Browser = "Chimera"
		return true
	}
	if strings.Contains(ua, "chrome") {
		desc.Browser = "Chrome"
		return true
	}
	if !strings.Contains(ua, "chrome") && strings.Contains(ua, "safari") {
		desc.Browser = "Safari"
		return true
	}

	return false
}
