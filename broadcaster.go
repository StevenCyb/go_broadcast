package broadcast

import "sync"

// New creates a new generic `Broadcaster`.
func New[T any]() Broadcaster[T] {
	return Broadcaster[T]{
		subscribers: make(map[string]map[chan T]bool),
	}
}

// Broadcaster is a struct where different routines can subscribe and publish to topics.
type Broadcaster[T any] struct {
	subscribers map[string]map[chan T]bool
	rwLock      sync.RWMutex
}

// Subscribe adds a new subscriber for the given topic.
// It returns a new channel for the subscriber to receive data on.
func (b *Broadcaster[T]) Subscribe(topic string) chan T {
	b.rwLock.Lock()
	defer b.rwLock.Unlock()

	if _, ok := b.subscribers[topic]; !ok {
		b.subscribers[topic] = make(map[chan T]bool)
	}

	sub := make(chan T)
	b.subscribers[topic][sub] = true

	return sub
}

// Unsubscribe removes a subscriber from the given topic.
func (b *Broadcaster[T]) Unsubscribe(topic string, sub chan T) {
	b.rwLock.Lock()
	defer b.rwLock.Unlock()

	if _, ok := b.subscribers[topic]; ok {
		delete(b.subscribers[topic], sub)
	}
}

// Publish sends a data to all subscribers of the given topic.
func (b *Broadcaster[T]) Publish(topic string, data T) {
	b.rwLock.RLock()
	defer b.rwLock.RUnlock()

	if subs, ok := b.subscribers[topic]; ok {
		for sub := range subs {
			sub <- data
		}
	}
}
