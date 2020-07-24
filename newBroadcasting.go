package nyubroadcasting

import (
	"sync"
	"time"
)

// Broadcasting is the primary means of external communication of bot
type Broadcasting struct {
	Broadcaster chan string
	listeners   map[string]chan ExternalCommunication

	sync sync.Mutex
}

// NewBroadcasting return a new broadcaster
func NewBroadcasting() *Broadcasting {
	return &Broadcasting{
		listeners: make(map[string]chan ExternalCommunication, 0),

		sync: sync.Mutex{},
	}
}

// Start the broadcast channel and run the console
// console is the function that listen the broadcast channel and send to all listeners
func (b *Broadcasting) Start() {
	b.Broadcaster = make(chan string, 256)
	go b.console()
}

// Stop all listeners channel and close the broadcast channel
func (b *Broadcasting) Stop() {
	b.StopAllListener()
	close(b.Broadcaster)
}

// CreateNewListener create a new listener for this broadcaster
// the channel create is a buffered channel (256)
// the name given will be used to identify the listener and ask any actions
// for example, you call StopListener with the same name for close the listener
func (b *Broadcasting) CreateNewListener(name string) chan ExternalCommunication {
	listener := make(chan ExternalCommunication, 256)
	b.listeners[name] = listener
	return listener
}

// StopListener close a listener if he exist and is delete from the listeners
func (b *Broadcasting) StopListener(name string) {
	if listener, ok := b.listeners[name]; ok {
		b.sync.Lock()
		close(listener)
		delete(b.listeners, name)
		b.sync.Unlock()
	}
}

// StopAllListener close all listeners and delete from the listeners
func (b *Broadcasting) StopAllListener() {
	for name, listener := range b.listeners {
		close(listener)
		delete(b.listeners, name)
	}
}

// console listen the broadcast channel and send back to everyone listeners a ExternalCommunication struct
func (b *Broadcasting) console() {
	for {
		str, alive := <-b.Broadcaster
		if !alive {
			break
		}
		communication := ExternalCommunication{
			Message:   str,
			CreatedAt: time.Time.UTC(time.Now()),
		}

		b.sync.Lock()
		for _, listener := range b.listeners {
			listener <- communication
		}
		b.sync.Unlock()
	}
}
