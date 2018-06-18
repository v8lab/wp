package state

type StateIntf interface {
	Execute()
	Stop()
	Running()
}
