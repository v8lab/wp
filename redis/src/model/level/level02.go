package level

import (
	TwLog "twlog"
)

func NewLevel02Stu() *Level02Stu {
	var Level02 Level02Stu
	return &Level02
}

type Level02Stu struct {
	Prn *TwLog.TwLogStu
	Sub Level03Stu
}

func (r *Level02Stu) Execute(Prn *TwLog.TwLogStu) {
	r.Prn = Prn.Enter()
	defer Prn.Exit()

	r.Prn.Write("helllo, i am level 02 ")
	r.Sub.Execute(r.Prn)
	r.Prn.Write("helllo, i am level 02 ")
}
