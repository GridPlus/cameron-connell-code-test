package util

import "fmt"

type PrereqUpdate Target

func (prereq *PrereqUpdate) appCode() string {
	return prereq.AppCode
}

func (prereq *PrereqUpdate) currentVersion() string {
	return prereq.TargetVersion
}

func (prereq *PrereqUpdate) processPrereq() {
	fmt.Printf("%v has prerquisete updates, updating to %v\n", prereq.appCode(), prereq.currentVersion())
}
