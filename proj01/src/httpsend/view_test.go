package httpsend

import (
	"fmt"
	"testing"

	base "httpsend/base"
	setting "httpsend/setting"
)

func Test_Setting(t *testing.T) {

	CaseId := "caseid01"
	Queue := "11111222223333333"

	Setting := setting.NewSetstingFile("./xmlfile/setting.xml")

	Case, eCase := Setting.Find(CaseId)
	if eCase {
		for _, queue := range Queue {
			UnitId := string(queue)
			Unit, eUnit := Case.Find(UnitId)
			if eUnit {
				EntryRst, ret := base.Execute(Unit)
				if ret != 0 {
					fmt.Println("\tEntryRst    --> ", EntryRst.GetBody())
				}
			}
		}
	}
}
