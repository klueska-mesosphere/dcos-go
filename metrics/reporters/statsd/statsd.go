package statsd

import (
	"errors"
	"fmt"
	"math"
	"time"

	statsd "github.com/smira/go-statsd.git"
	tally "github.com/uber-go/tally"
)

type Reporter interface {
	tally.StatsReporter
}

type reporter struct {
	client  *statsd.Client
	options *ReporterOptions
}

func NewReporter(addr string, options ...Option) (Reporter, error) {
	reporterOptions := &ReporterOptions{
		TagFormat: TAG_FORMAT_NONE,
	}

	clientOptions := []statsd.Option{}
	for _, option := range options {
		clientOptions = append(clientOptions, option(reporterOptions))
	}

	client := statsd.NewClient(addr, clientOptions...)
	if client == nil {
		return nil, errors.New("Unable to initialize StatsD client")
	}

	reporter := &reporter{
		client:  client,
		options: reporterOptions,
	}

	return reporter, nil
}

func (r *reporter) ReportCounter(name string, tags map[string]string, value int64) {
	r.client.Incr(name, value, r.formatTags(tags)...)
}

func (r *reporter) ReportGauge(name string, tags map[string]string, value float64) {
	r.client.FGauge(name, value, r.formatTags(tags)...)
}

func (r *reporter) ReportTimer(name string, tags map[string]string, interval time.Duration) {
	r.client.PrecisionTiming(name, interval, r.formatTags(tags)...)
}

func (r *reporter) ReportHistogramValueSamples(
	name string,
	tags map[string]string,
	buckets tally.Buckets,
	bucketLowerBound,
	bucketUpperBound float64,
	samples int64,
) {
	name = fmt.Sprintf(
		"%s.%s-%s",
		name,
		r.valueBucketString(bucketLowerBound),
		r.valueBucketString(bucketUpperBound))

	r.client.Incr(name, samples, r.formatTags(tags)...)
}

func (r *reporter) ReportHistogramDurationSamples(
	name string,
	tags map[string]string,
	buckets tally.Buckets,
	bucketLowerBound,
	bucketUpperBound time.Duration,
	samples int64,
) {

	name = fmt.Sprintf(
		"%s.%s-%s",
		name,
		r.durationBucketString(bucketLowerBound),
		r.durationBucketString(bucketUpperBound))

	r.client.Incr(name, samples, r.formatTags(tags)...)
}

func (r *reporter) Capabilities() tally.Capabilities {
	return r
}

func (r *reporter) Reporting() bool {
	return true
}

func (r *reporter) Tagging() bool {
	switch r.options.TagFormat {
	case TAG_FORMAT_INFLUX:
	case TAG_FORMAT_DATADOG:
		return true
	}
	return false
}

func (r *reporter) Flush() {
	// no-op
}

func (r *reporter) formatTags(tags map[string]string) []statsd.Tag {
	ts := []statsd.Tag{}
	if r.Tagging() {
		for k, v := range tags {
			ts = append(ts, statsd.StringTag(k, v))
		}
	}
	return ts
}

func (r *reporter) valueBucketString(upperBound float64) string {
	if upperBound == math.MaxFloat64 {
		return "infinity"
	}
	if upperBound == -math.MaxFloat64 {
		return "-infinity"
	}
	return fmt.Sprintf("%.6f", upperBound)
}

func (r *reporter) durationBucketString(upperBound time.Duration) string {
	if upperBound == time.Duration(math.MaxInt64) {
		return "infinity"
	}
	if upperBound == time.Duration(math.MinInt64) {
		return "-infinity"
	}
	return upperBound.String()
}
