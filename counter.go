package libratini

import "sync"

type Counter struct {
	name    string
	mutex   sync.Mutex
	count   int64
	channel chan int64
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	c.count++
	c.channel <- c.count
	c.mutex.Unlock()
}
