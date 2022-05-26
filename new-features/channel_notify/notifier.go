package main

import (
	"sync"
	"sync/atomic"
)

// EventType event type
type EventType uint8

const (
	EventDataChanged EventType = 1
)

// Notifier 数据变动通知
type Notifier struct {
	mu     sync.Mutex   // protects following fields
	notify atomic.Value // of chan EventType, created lazily
}

func newNotifier() *Notifier {
	return &Notifier{}
}

func (d *Notifier) Send(event EventType) {
	if ch, ok := d.notify.Load().(chan EventType); ok {
		ch <- event
	}
}

func (d *Notifier) sendNoneBlock(event EventType) {
	if ch, ok := d.notify.Load().(chan EventType); ok {
		select {
		case ch <- event:
		default:
		}
	}
}

// Watch return a received channel
func (d *Notifier) Watch() <-chan EventType {
	ch := d.notify.Load()
	if ch != nil {
		return ch.(chan EventType)
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	ch = d.notify.Load()
	if ch == nil {
		ch = make(chan EventType, 1)
		d.notify.Store(ch)
	}
	return ch.(chan EventType)
}
