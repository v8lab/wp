package level

import (
	TwLog "twlog"
)

func NewLevel03Stu() *Level03Stu {
	var Level03 Level03Stu
	return &Level03
}

type Level03Stu struct {
	Prn *TwLog.TwLogStu
}

func (r *Level03Stu) Execute(Prn *TwLog.TwLogStu) {
	r.Prn = Prn.Enter()
	defer Prn.Exit()

	r.Prn.Write("helllo, i am level 03 ")

}
