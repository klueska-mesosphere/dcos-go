package scope

import (
	"errors"
	"time"

	tally "github.com/uber-go/tally"
)

const (
	Prefix = "dcos"
)

type Scope struct {
	tally.Scope
	CloserContext
}

func New(component string, interval time.Duration, options ...Option) (*Scope, error) {
	opts := &ScopeOptions{}

	for _, option := range options {
		option(opts)
	}

	if opts.StatsReporter == nil && opts.CachedStatsReporter == nil {
		return nil, errors.New("Must pass one of tally.StatsReporter or tally.CachedStatsReporter options")
	}

	if opts.StatsReporter != nil && opts.CachedStatsReporter != nil {
		return nil, errors.New("Must pass either a tally.StatsReporter or a tally.CachedStatsReporter option, not both")
	}

	topts := tally.ScopeOptions{
		Prefix:    Prefix,
		Tags:      opts.Tags,
		Separator: opts.Separator,
	}
	if opts.StatsReporter != nil {
		topts.Reporter = opts.StatsReporter
	} else if opts.CachedStatsReporter != nil {
		topts.CachedReporter = opts.CachedStatsReporter
	}

	rootTallyScope, closer := tally.NewRootScope(topts, interval)
	subTallyScope := rootTallyScope.SubScope(component)

	scope := &Scope{
		Scope:         subTallyScope,
		CloserContext: NewCloserContext(closer),
	}

	if opts.ReportRuntimeMetrics {
		scope.ReportRuntimeMetrics(interval)
	}

	return scope, nil
}

func (s *Scope) ReportRuntimeMetrics(interval time.Duration) {
	scope := s.SubScope("go")

	goRoutineMetrics := NewGoRoutineMetrics(scope)
	memoryMetrics := NewMemoryMetrics(scope)

	go func() {
		for {
			select {
			case <-s.Done():
				return
			case <-time.After(interval):
				goRoutineMetrics.Report()
				memoryMetrics.Report()
			}
		}
	}()
}
