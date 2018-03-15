package sysinfo

import (
	. "base"
)

const (
	KindCliList string = "sysinfo/clilist"
)

func init() {
	Factory := GetSingleEntryFactory()
	Factory.Add(KindCliList, func() EntryIntf { return &CliListStu{NewEmptyEntry(KindCliList)} })

	ProRso := GetSingleProblemResolveMap()
	ProRso.Add(KindCliList, KindCliList)
}

type CliListStu struct {
	*EntryStu
}

func (r *CliListStu) Execute() (ret int) {

	return
}
