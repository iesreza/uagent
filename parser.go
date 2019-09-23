package uagent

import "github.com/iesreza/uagent/classifiers"

func Parse(ua string) classifiers.Describer {
	describer := classifiers.Describer{}
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
