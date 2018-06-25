package base

import (
	setting "httpsend/setting"
)

func Execute(Unit *setting.UnitStu) (Out *EntryStu, ret int) {
	Factory := GetEntryFactory()
	Entry := Factory.Create("")
	Entry.Construct(Unit)
	Out, ret = Entry.Execute()
	if ret == 0 {
		Unit.PrintSuccess()
	} else {
		Unit.PrintFailed()
	}
	return
}

var (
	Common string = ""
)

func init() {
	Factory := GetEntryFactory()
	Factory.Add(Common, func() EntryIntf {
		Entry := &CommonStu{EntryStu: NewEmptyEntry(Common)}
		return Entry
	})
}

type CommonStu struct {
	*EntryStu
}
