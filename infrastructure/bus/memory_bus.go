package bus

import (
	"fmt"
	"modular-monolith-libraryApp/shared/port"
)

type MemoryBus struct {
	handlers map[string][]func(any)
}

func New() port.EventBus {
	return &MemoryBus{handlers: make(map[string][]func(any))}
}

func (b *MemoryBus) Publish(event any) {
	t := fmt.Sprintf("%T", event)
	for _, h := range b.handlers[t] {
		h(event)
	}
}

func (b *MemoryBus) Subscribe(t string, handler func(any)) {
	b.handlers[t] = append(b.handlers[t], handler)
}
