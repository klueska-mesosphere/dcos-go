package prometheus

import (
	client "github.com/m3db/prometheus_client_golang/prometheus"
	prometheus "github.com/uber-go/tally/prometheus"
)

const Separator = prometheus.DefaultSeparator

type Reporter interface {
	prometheus.Reporter
}

func NewReporter(options ...Option) (Reporter, error) {
	opts := &ReporterOptions{}

	for _, option := range options {
		option(opts)
	}

	registry := client.NewRegistry()

	popts := prometheus.Options{
		Registerer: registry,
		Gatherer:   registry,
	}

	reporter := prometheus.NewReporter(popts)

	return reporter, nil
}
