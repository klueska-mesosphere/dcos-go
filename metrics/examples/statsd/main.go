package main

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	statsd "github.com/dcos/dcos-go/metrics/reporters/statsd"
	scope "github.com/dcos/dcos-go/metrics/scope"
)

func main() {
	reporter, err := statsd.NewReporter(
		"127.0.0.1:8125",
		statsd.TagStyle(statsd.TAG_FORMAT_DATADOG))
	if err != nil {
		log.Errorf("Error creating reporter: %v.", err)
		os.Exit(1)
	}

	scope, err := scope.New(
		"component",
		time.Second,
		scope.StatsReporter(reporter),
		scope.ReportRuntimeMetrics())
	if err != nil {
		log.Errorf("Error creating scope: %v.", err)
		os.Exit(1)
	}
	defer scope.Close()

	littlehand := time.NewTicker(time.Millisecond)
	bighand := time.NewTicker(60 * time.Millisecond)
	hugehand := time.NewTicker(60 * 60 * time.Millisecond)

	gauge := scope.Gauge("gauge")
	timer := scope.Timer("timer")
	counter := scope.Tagged(
		map[string]string{
			"tag1": "value1",
			"tag2": "value2",
			"tag3": "value3",
		}).Counter("counter")

	go func() {
		for {
			select {
			case <-littlehand.C:
				counter.Inc(1)
			case <-bighand.C:
				gauge.Update(42.1)
			case <-hugehand.C:
				timer.Record(60 * 60 * 10 * time.Millisecond)
			}
		}
	}()

	select {}
}
