package pubsub

import "sync"

var lock = &sync.Mutex{}

type Instance struct {
	pub *Publisher
}

var instance *Instance

func GetPub() *Publisher {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Instance{
				pub: NewPublisher(),
			}
		}
	}

	return instance.pub
}
