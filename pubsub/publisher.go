package pubsub

import "sync"

type Publisher struct {
	subscribers map[string][]*Subscriber
	mutex       sync.Mutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string][]*Subscriber),
	}
}

func (p *Publisher) Subscribe(event string, subscriber *Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.subscribers[event]; !ok {
		p.subscribers[event] = make([]*Subscriber, 0)
	}
	p.subscribers[event] = append(p.subscribers[event], subscriber)
}

func (p *Publisher) Unsubscribe(event string, subscriber *Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if subscribers, ok := p.subscribers[event]; ok {
		for i, sub := range subscribers {
			if sub == subscriber {
				p.subscribers[event] = append(subscribers[:i], subscribers[i+1:]...)
				return
			}
		}
	}
}

func (p *Publisher) Publish(event string, eventData interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if subscribers, ok := p.subscribers[event]; ok {
		eventPtr := &Event{
			Name: event,
			Data: eventData,
		}
		for _, subscriber := range subscribers {
			go func(sub *Subscriber) {
				sub.Ch <- eventPtr
			}(subscriber)
		}
	}
}

func (p *Publisher) Listen() {
	for _, subscribers := range p.subscribers {
		for _, subscriber := range subscribers {
			go (*subscriber).StartListening()
		}
	}
}
