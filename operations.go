package files

import (
	"sync"

	"github.com/cornelk/hashmap"
)

type Operations struct {
	operation *hashmap.HashMap
}

type Operation struct {
	sync.RWMutex
}

func New() *Operations {
	return &Operations{
		operation: &hashmap.HashMap{},
	}
}

func (o *Operations) ReadLock(key string, call func()) {
	amount, ok := o.operation.Get(key)
	if ok {
		q := amount.(*Operation)
		q.RLock()
		call()
		q.RUnlock()
	} else {
		o.newOperation(key)
		o.ReadLock(key, call)
	}
}

func (o *Operations) WriteLock(key string, call func()) {
	amount, ok := o.operation.Get(key)
	if ok {
		q := amount.(*Operation)
		q.Lock()
		call()
		q.Unlock()
	} else {
		o.newOperation(key)
		o.WriteLock(key, call)
	}
}

func (o *Operations) newOperation(key string) {
	q := new(Operation)
	o.operation.Set(key, q)
}
