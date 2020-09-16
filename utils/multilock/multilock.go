package multilock

import (
	"sync"
)

type mutexItem struct {
	mu      sync.RWMutex
	counter int64
}

type messageGet struct {
	key string
	c   chan *sync.RWMutex
}

type messagePut string

type Multilock interface {
	Get(key string) *sync.RWMutex
	Put(key string)
	Close()
}

// multiLocker is key based multiple locker
type multiLocker struct {
	gets  chan messageGet
	puts  chan messagePut
	inUse map[string]*mutexItem
	done  chan struct{}
	mp    sync.Pool
	cp    sync.Pool
}

// Get get key related sync.Mutex
func (l *multiLocker) Get(key string) *sync.RWMutex {
	msg := messageGet{
		key: key,
		c:   l.cp.Get().(chan *sync.RWMutex),
	}
	l.gets <- msg
	mu := <-msg.c
	l.cp.Put(msg.c)
	return mu
}

// Put release key related mutexItem
func (l *multiLocker) Put(key string) {
	l.puts <- messagePut(key)
}

// Close will stop schedule
func (l *multiLocker) Close() {
	close(l.done)
}

func (l *multiLocker) handleGet(msg *messageGet) {
	key := msg.key
	c := msg.c
	mi, ok := l.inUse[key]
	if !ok {
		mi = l.mp.Get().(*mutexItem)
		l.inUse[key] = mi
	}
	mi.counter++
	c <- &mi.mu
}

func (l *multiLocker) handlePut(msg *messagePut) {
	key := string(*msg)
	mi, ok := l.inUse[key]
	if !ok {
		panic("should call lock first")
	}
	mi.counter--
	if mi.counter == 0 {
		l.mp.Put(mi)
		delete(l.inUse, key)
	}
}

func (l *multiLocker) schedule() {
loop:
	for {
		select {
		case msg := <-l.gets:
			l.handleGet(&msg)
		case msg := <-l.puts:
			l.handlePut(&msg)
		case <-l.done:
			break loop
		}
	}
}

// NewMultilock return a new multiLocker
func NewMultilock() Multilock {
	l := &multiLocker{
		gets:  make(chan messageGet, 1000),
		puts:  make(chan messagePut, 1000),
		inUse: make(map[string]*mutexItem),
		mp: sync.Pool{
			New: func() interface{} { return &mutexItem{} },
		},
		cp: sync.Pool{
			New: func() interface{} { return make(chan *sync.RWMutex, 1) },
		},
	}
	go l.schedule()
	return l
}
