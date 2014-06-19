package libratini

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

func Middleware(config Config) martini.Handler {
	dashboard := NewDashboard(config)
	timeGauge := dashboard.GetGauge(config.Prefix + "time")

	return func(response http.ResponseWriter, context martini.Context) {
		start := time.Now()

		context.Next()
		status := response.(martini.ResponseWriter).Status()

		go (func() {
			counterName := fmt.Sprintf("%s%d.count", config.Prefix, status)
			counter := dashboard.GetCounter(counterName)
			counter.Increment()

			responseTime := int64(time.Since(start) / time.Millisecond)
			timeGauge.Measure(responseTime)
		})()
	}
}
