package libratini

import "sync"

type Gauge struct {
	name    string
	mutex   sync.Mutex
	channel chan int64
}

func (g *Gauge) Measure(value int64) {
	g.mutex.Lock()
	g.channel <- value
	g.mutex.Unlock()
}
