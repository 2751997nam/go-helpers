package event

import "sync"

var lock = &sync.Mutex{}

type Instance struct {
	reg *EventRegister
}

var instance *Instance

func GetRegister() *EventRegister {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Instance{
				reg: &EventRegister{
					registers: []RegisterFunc{},
				},
			}
		}
	}

	return instance.reg
}
