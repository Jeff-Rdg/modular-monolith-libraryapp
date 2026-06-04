package port

type EventBus interface {
	Publish(event any)
	Subscribe(eventType string, handler func(event any))
}
