package base

import (
	"sync"
)

var SingleProblemResolveMap ProblemResolveMapStu
var SingleProblemResolveMapOnce sync.Once

func GetSingleProblemResolveMap() *ProblemResolveMapStu {
	SingleProblemResolveMapOnce.Do(SingleProblemResolveMap.Init)
	return &SingleProblemResolveMap
}

type ProblemResolveMapStu struct {
	ProRso map[string]string
	Unkown string
}

func (r *ProblemResolveMapStu) Init() {
	r.ProRso = make(map[string]string)
	r.Unkown = KINDUNKOWN

}

func (r *ProblemResolveMapStu) GetSolve(KindProblem string) (KindResolve string) {
	if Kind, ok := r.ProRso[KindProblem]; ok {
		KindResolve = Kind
	} else {
		KindResolve = r.Unkown
	}
	return
}
func (r *ProblemResolveMapStu) Add(Problem, Resolve string) {
	r.ProRso[Problem] = Resolve
}
