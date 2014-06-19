package libratini

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/rcrowley/go-librato"
	"net/http"
	"time"
)

type Config struct {
	Collate int
	Prefix  string
	Source  string
	Token   string
	User    string
}

func Middleware(config Config) martini.Handler {
	api := librato.NewCollatedMetrics(config.User, config.Token, config.Source, config.Collate)
	requestTime := api.GetGauge(config.Prefix + "time")
	counters := make(map[string]*Counter)

	incrementCounter := func(name string) {
		counter, exists := counters[name]
		if exists == false {
			counters[name] = &Counter{channel: api.NewCounter(name)}
			counter = counters[name]
		}
		counter.Increment()
	}

	return func(response http.ResponseWriter, context martini.Context) {
		start := time.Now()
		defer (func() {
			rw := response.(martini.ResponseWriter)
			status := rw.Status()
			counterName := fmt.Sprintf("%s%d.count", config.Prefix, status)
			incrementCounter(counterName)
			requestTime <- int64(time.Since(start) / time.Millisecond)
		})()

		context.Next()

	}
}
