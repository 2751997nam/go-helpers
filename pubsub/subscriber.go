package pubsub

import (
	"fmt"
)

type Subscriber struct {
	Handle Handle
	Ch     chan *Event
}

type Handle func(data any)

func NewSubscriber(handle Handle) *Subscriber {
	return &Subscriber{
		Handle: handle,
		Ch:     make(chan *Event),
	}
}

func (s *Subscriber) StartListening() {
	for {
		select {
		case event, ok := <-s.Ch:
			if !ok {
				fmt.Printf("Subscriber %d channel closed\n")
				return
			}
			// fmt.Printf("Subscriber %d received event: %s\n", s.ID, event.Name)
			// Process the event data
			// fmt.Printf("Event Data: %#v\n", event.Data)
			s.Handle(event.Data)
		}
	}
}
