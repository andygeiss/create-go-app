package event

// Bus provides the low-level transport layer for events.
type Bus struct {
	TopicSubscribers map[string][]func(data interface{})
}

// Publish sends data in context of a specific topic to all registered functions.
func (b *Bus) Publish(topic string, data interface{}) {
	for _, fn := range b.TopicSubscribers[topic] {
		fn(data)
	}
}

// Subscribe registers a function to a specific topic.
func (b *Bus) Subscribe(topic string, fn func(data interface{})) {
	if _, exists := b.TopicSubscribers[topic]; !exists {
		b.TopicSubscribers[topic] = make([]func(data interface{}), 0)
	}
	b.TopicSubscribers[topic] = append(b.TopicSubscribers[topic], fn)
}

// NewBus creates a Bus and returns its pointer.
func NewBus() *Bus {
	return &Bus{
		TopicSubscribers: make(map[string][]func(data interface{})),
	}
}
