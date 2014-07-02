package libratini

import (
	"fmt"
	"github.com/wunderlist/go-librato"
)

type Dashboard struct {
	config   Config
	gauges   map[string]*Gauge
	counters map[string]*Counter
	api      librato.Metrics
}

func NewDashboard(config Config) *Dashboard {
	d := &Dashboard{config: config}
	d.gauges = make(map[string]*Gauge)
	d.counters = make(map[string]*Counter)
	d.api = librato.NewCollatedMetrics(config.User, config.Token, config.Source, config.Collate)
	return d
}

func (d *Dashboard) GetGauge(name string) *Gauge {
	name = fmt.Sprintf("%s%s", d.config.Prefix, name)
	gauge, exists := d.gauges[name]
	if exists == false {
		channel := d.api.NewGauge(name)
		gauge = &Gauge{name: name, channel: channel}
		d.gauges[name] = gauge
	}
	return gauge
}

func (d *Dashboard) GetCounter(name string) *Counter {
	name = fmt.Sprintf("%s%s", d.config.Prefix, name)
	counter, exists := d.counters[name]
	if exists == false {
		channel := d.api.NewCounter(name)
		counter = &Counter{name: name, channel: channel}
		d.counters[name] = counter
	}
	return counter
}
