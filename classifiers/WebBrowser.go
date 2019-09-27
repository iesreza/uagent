package classifiers

import (
	"strings"
)

type WebBrowser struct{}

func (WebBrowser) Eval(ua string, desc *Describer) bool {
	if strings.Contains(ua, "opera mini") || strings.Contains(ua, "opr/") || strings.Contains(ua, "opt/") {
		desc.Browser = "Opera"
		desc.Device = Mobile
		desc.Platform = Mobile
		return true
	}
	if strings.Contains(ua, "opios") {
		desc.Browser = "Opera"
		desc.Device = Mobile
		desc.Platform = Mobile
		desc.OS = IoS
	}
	if strings.Contains(ua, "presto") && strings.Contains(ua, "opera") {
		desc.Browser = "Opera"
		return true
	}
	if strings.Contains(ua, "gecko") && strings.Contains(ua, "firefox") {
		desc.Browser = "Firefox"
		return true
	}
	if strings.Contains(ua, ".net") || strings.Contains(ua, "trident") || strings.Contains(ua, "msie") {
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
	if strings.Contains(ua, "crios") {
		desc.Browser = "Chrome"
		desc.Device = Mobile
		desc.Platform = Mobile
		desc.OS = IoS
		return true
	}
	if strings.Contains(ua, "fxios") {
		desc.Browser = "Firefox"
		desc.Device = Mobile
		desc.Platform = Mobile
		desc.OS = IoS
		return true
	}
	if strings.Contains(ua, "vivaldi") {
		desc.Browser = "Vivaldi"
		return true
	}
	if !strings.Contains(ua, "chrome") && strings.Contains(ua, "safari") {
		desc.Browser = "Safari"
		return true
	}

	return false
}
