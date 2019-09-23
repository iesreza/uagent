package uagent

import (
	"github.com/iesreza/uagent/classifiers"
	"strings"
)

func Parse(ua string) classifiers.Describer {
	describer := classifiers.Describer{}
	ua = strings.ToLower(ua)
	for _, cList := range classifiers.Classifiers {
		if len(cList) > 0 {
			for _, classifier := range cList {
				if classifier.Eval(ua, &describer) {
					break
				}
			}
		}
	}

	return describer
}
