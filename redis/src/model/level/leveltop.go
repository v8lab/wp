package level

import (
	"fmt"
	"time"
	TwLog "twlog"
)

func NewLevelTopStu(Id string) *LevelTopStu {
	var LevelTop LevelTopStu
	LevelTop.Init(Id)
	return &LevelTop
}

type LevelTopStu struct {
	Prn  *TwLog.TwLogStu
	Suba Level01Stu
	Subb Level01Stu
	Id   string
}

func (r *LevelTopStu) Init(Id string) {
	Opt := TwLog.NewTwLogOptStu()
	Opt.Format = "%v\t%v\t\t\n"
	Opt.Time = time.Now().UnixNano()
	r.Id = Id
	Opt.Tag = r.Id
	r.Prn = TwLog.NewTwLogStu(Opt)
}
func (r *LevelTopStu) Execute() {
	r.Prn = r.Prn.Enter()
	defer r.Prn.Exit()
	r.Prn.Write("helllo, i am level tops")
	r.Prn.Write("helllo, i am level tops")
	r.Suba.Execute(r.Prn)
	fmt.Println(r.Prn.GetRst())

	//	r.Subb.Execute(r.Prn)
	//	fmt.Println(r.Prn.GetRst())
}
