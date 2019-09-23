package classifiers

import (
	"strings"
)

type SmartTVClassifier struct{}

var SmartTVUserAgents = map[string]string{
	"philipstv":      Linux,
	"hbbtv":          Linux,
	"smarttv":        Linux,
	"smart-tv":       Linux,
	"netcast":        Linux,
	"webos":          WebOS,
	"samsungbrowser": Linux,
	"webtv":          Linux,
	"inettvbrowser":  Linux,
	"googletv":       "Android",
}

func (SmartTVClassifier) Eval(ua string, desc *Describer) bool {
	for ag, os := range SmartTVUserAgents {
		if strings.Contains(ua, ag) {
			desc.OS = os
			desc.Platform = SmartTV
			desc.Device = SmartTV
			return true
		}
	}
	return false
}
