package scope

import (
	tally "github.com/uber-go/tally"
)

type Option func(o *ScopeOptions)

type ScopeOptions struct {
	StatsReporter        tally.StatsReporter
	CachedStatsReporter  tally.CachedStatsReporter
	Tags                 map[string]string
	Separator            string
	ReportRuntimeMetrics bool
}

func StatsReporter(reporter tally.StatsReporter) Option {
	return func(o *ScopeOptions) {
		o.StatsReporter = reporter
	}
}

func CachedStatsReporter(reporter tally.CachedStatsReporter) Option {
	return func(o *ScopeOptions) {
		o.CachedStatsReporter = reporter
	}
}

func Tags(tags map[string]string) Option {
	return func(o *ScopeOptions) {
		o.Tags = tags
	}
}

func Separator(separator string) Option {
	return func(o *ScopeOptions) {
		o.Separator = separator
	}
}

func ReportRuntimeMetrics() Option {
	return func(o *ScopeOptions) {
		o.ReportRuntimeMetrics = true
	}
}
