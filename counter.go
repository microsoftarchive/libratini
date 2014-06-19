package libratini

import "sync"

type Counter struct {
	mutex   sync.Mutex
	count   int64
	channel chan int64
}

func (counter *Counter) Increment() {
	counter.mutex.Lock()
	counter.count++
	counter.channel <- counter.count
	counter.mutex.Unlock()
}
