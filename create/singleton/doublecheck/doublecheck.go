package doublecheck

import "sync"

type Singleton struct {
	data string
}

func NewSingleton(data string) *Singleton {
	return &Singleton{data: data}
}

func (s Singleton) String() string {
	return "data" + s.data
}

var s *Singleton

var lock = sync.Mutex{}

func Get() *Singleton {
	if s == nil {
		lock.Lock()
		defer lock.Unlock()
		if s == nil {
			s = NewSingleton("data")
		}
	}
	return s
}
