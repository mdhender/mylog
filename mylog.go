package mylog

import (
	"fmt"
	"sync"
)

type Log struct {
	sync.Mutex
	nextId int
	log    []item
}

type item struct {
	id      int
	payload interface{}
}

func New(initialCapacity int) (*Log, error) {
	if initialCapacity < 1 {
		return nil, fmt.Errorf("invalid capacity")
	}
	l := Log{
		nextId: 1,
		log:    make([]item, initialCapacity),
	}
	l.log[0] = item{id: 0}
	return &l, nil
}

// Add appends a payload to the log. Payload must be
// an immutable object for the log to be immutable.
func (l *Log) Add(p interface{}) (id int, err error) {
	l.Lock()
	id, l.nextId = l.nextId, l.nextId+1
	l.log = append(l.log, item{id, p})
	l.Unlock()
	return id, err
}

func (l *Log) Get(id int) (p interface{}, err error) {
	tmp := l.log // this gives us a snapshot to avoid races (not really)
	if id < 0 || len(tmp)-1 < id {
		return nil, fmt.Errorf("invalid id")
	}
	return tmp[id].payload, nil
}

func (l *Log) GetRange(id int) (p []item, err error) {
	tmp := l.log // this gives us a snapshot to avoid races (not really)
	if id < 0 || len(tmp)-1 < id {
		return nil, fmt.Errorf("invalid id")
	}
	return tmp[id:], nil
}

func (i item) ID() int {
	return i.id
}

func (i item) Payload() interface{} {
	return i.payload
}
