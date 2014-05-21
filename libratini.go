package libratini

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/rcrowley/go-librato"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Collate int
	Prefix  string
	Source  string
	Token   string
	User    string
}

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

func Middleware(config Config) martini.Handler {
	api := librato.NewCollatedMetrics(config.User, config.Token, config.Source, config.Collate)
	requestTime := api.GetGauge(config.Prefix + "2xx.time")
	counters := make(map[string]*Counter)

	incrementCounter := func(name string) {
		counter, exists := counters[name]
		if exists == false {
			counters[name] = &Counter{channel: api.NewCounter(name)}
			counter = counters[name]
		}
		counter.Increment()
	}

	return func(response http.ResponseWriter, request *http.Request, context martini.Context) {
		start := time.Now()
		rw := response.(martini.ResponseWriter)
		context.Next()

		status := rw.Status()
		counterName := fmt.Sprintf("%s%d.count", config.Prefix, status)
		incrementCounter(counterName)

		if status == 200 || status == 201 {
			requestTime <- int64(time.Since(start) / time.Millisecond)
		}
	}
}
