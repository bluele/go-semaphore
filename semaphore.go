package semaphore

import (
	"time"
)

const (
	resource byte = 0
)

type Semaphore struct {
	channel chan byte
	permit  int
}

type TimeLimitedSemaphore struct {
	channel chan byte
	buffer  chan byte
	per     time.Duration
	permit  int
}

func NewSemaphore(permit int) *Semaphore {
	sm := &Semaphore{
		permit:  permit,
		channel: make(chan byte, permit),
	}
	for i := 0; i < permit; i++ {
		sm.channel <- resource
	}
	return sm
}

func NewTimeLimitedSemaphore(permit int, per time.Duration) *TimeLimitedSemaphore {
	sm := &TimeLimitedSemaphore{
		permit:  permit,
		per:     per,
		channel: make(chan byte, permit),
		buffer:  make(chan byte, permit),
	}
	for i := 0; i < permit; i++ {
		sm.channel <- resource
	}
	go sm.gc()
	return sm
}

func (sm *Semaphore) Aquire(n int) {
	for i := 0; i < n; i++ {
		<-sm.channel
	}
}

func (sm *Semaphore) Release() {
	select {
	case sm.channel <- resource:
	default:
		panic("too many release")
	}
}

func (sm *Semaphore) Available() int {
	return len(sm.channel)
}

func (sm *Semaphore) Wait() {
	sm.Aquire(sm.permit)
}

func (sm *TimeLimitedSemaphore) Aquire(n int) {
	for i := 0; i < n; i++ {
		<-sm.channel
	}
}

func (sm *TimeLimitedSemaphore) Release() {
	select {
	case sm.buffer <- resource:
	default:
		panic("too many release")
	}
}

func (sm *TimeLimitedSemaphore) Available() int {
	return len(sm.channel)
}

func (sm *TimeLimitedSemaphore) Wait() {
	sm.Aquire(sm.permit)
}

func (sm *TimeLimitedSemaphore) gc() {
	tick := time.NewTicker(sm.per)
	for {
	Main:
		select {
		case <-tick.C:
			for {
				select {
				case b := <-sm.buffer:
					sm.channel <- b
				default:
					break Main
				}
			}

		}
	}
}
