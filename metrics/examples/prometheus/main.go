package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	prometheus "github.com/dcos/dcos-go/metrics/reporters/prometheus"
	scope "github.com/dcos/dcos-go/metrics/scope"
)

func main() {
	reporter, err := prometheus.NewReporter()
	if err != nil {
		log.Errorf("Error creating reporter: %v.", err)
		os.Exit(1)
	}

	scope, err := scope.New(
		"component",
		time.Second,
		scope.CachedStatsReporter(reporter),
		scope.Separator(prometheus.Separator),
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
	counter := scope.Counter("counter")

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

	http.Handle("/metrics", reporter.HTTPHandler())
	fmt.Printf("Serving :8080/metrics\n")
	fmt.Printf("%v\n", http.ListenAndServe(":8080", nil))

	select {}
}
