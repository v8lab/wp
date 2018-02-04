package a17observer

type SubjectIntf interface {
	Register(ObserverIntf)
	UnRegister(ObserverIntf)
	Notify()
}
