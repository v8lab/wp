package level

import (
	TwLog "twlog"
)

func NewLevel01Stu() *Level01Stu {
	var Level01 Level01Stu
	return &Level01
}

type Level01Stu struct {
	Prn   *TwLog.TwLogStu
	Sub   Level02Stu
	Sub03 Level03Stu
}

func (r *Level01Stu) Execute(Prn *TwLog.TwLogStu) {
	r.Prn = Prn.Enter()
	defer Prn.Exit()

	r.Prn.Write("helllo, i am level 01 ")
	r.Prn.Write("helllo, i am level 01 ")
	r.Sub.Execute(r.Prn)
	r.Prn.Write("helllo, i am level 01 ")
}
