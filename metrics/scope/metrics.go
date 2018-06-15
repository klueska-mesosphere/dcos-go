package scope

import (
	"runtime"

	tally "github.com/uber-go/tally"
)

type GoRoutineMetrics struct {
	NumGoRoutines tally.Gauge
}

func NewGoRoutineMetrics(s tally.Scope) *GoRoutineMetrics {
	scope := s.SubScope("goroutines")

	return &GoRoutineMetrics{
		NumGoRoutines: scope.Gauge("count"),
	}
}

func (g *GoRoutineMetrics) Report() {
	g.NumGoRoutines.Update(float64(runtime.NumGoroutine()))
}

type MemoryMetrics struct {
	Alloc         tally.Gauge
	TotalAlloc    tally.Gauge
	Sys           tally.Gauge
	Lookups       tally.Gauge
	Mallocs       tally.Gauge
	Frees         tally.Gauge
	HeapAlloc     tally.Gauge
	HeapSys       tally.Gauge
	HeapIdle      tally.Gauge
	HeapInuse     tally.Gauge
	HeapReleased  tally.Gauge
	HeapObjects   tally.Gauge
	StackInuse    tally.Gauge
	StackSys      tally.Gauge
	MSpanInuse    tally.Gauge
	MSpanSys      tally.Gauge
	MCacheInuse   tally.Gauge
	MCacheSys     tally.Gauge
	BuckHashSys   tally.Gauge
	GCSys         tally.Gauge
	OtherSys      tally.Gauge
	NextGC        tally.Gauge
	LastGC        tally.Gauge
	PauseTotalNs  tally.Gauge
	NumGC         tally.Gauge
	NumForcedGC   tally.Gauge
	GCCPUFraction tally.Gauge
}

func NewMemoryMetrics(s tally.Scope) *MemoryMetrics {
	scope := s.SubScope("memory")

	return &MemoryMetrics{
		Alloc:         scope.Gauge("alloc_bytes"),
		TotalAlloc:    scope.Gauge("total_alloc_bytes"),
		Sys:           scope.Gauge("sys_bytes"),
		Lookups:       scope.Gauge("lookups_count"),
		Mallocs:       scope.Gauge("mallocs_count"),
		Frees:         scope.Gauge("frees_count"),
		HeapAlloc:     scope.Gauge("heap_alloc_bytes"),
		HeapSys:       scope.Gauge("heap_sys_bytes"),
		HeapIdle:      scope.Gauge("heap_idle_bytes"),
		HeapInuse:     scope.Gauge("heap_inuse_bytes"),
		HeapReleased:  scope.Gauge("heap_released_bytes"),
		HeapObjects:   scope.Gauge("heap_objects_count"),
		StackInuse:    scope.Gauge("stack_inuse_bytes"),
		StackSys:      scope.Gauge("stack_inuse_bytes"),
		MSpanInuse:    scope.Gauge("mspan_inuse_bytes"),
		MSpanSys:      scope.Gauge("mspan_sys_bytes"),
		MCacheInuse:   scope.Gauge("mcache_inuse_bytes"),
		MCacheSys:     scope.Gauge("mcache_sys_bytes"),
		BuckHashSys:   scope.Gauge("buck_hash_sys_bytes"),
		GCSys:         scope.Gauge("gc_sys_bytes"),
		OtherSys:      scope.Gauge("other_sys_bytes"),
		NextGC:        scope.Gauge("next_gc_bytes"),
		LastGC:        scope.Gauge("last_gc_ns"),
		PauseTotalNs:  scope.Gauge("pause_total_ns"),
		NumGC:         scope.Gauge("gc_count"),
		NumForcedGC:   scope.Gauge("forced_gc_count"),
		GCCPUFraction: scope.Gauge("gc_cpu_fraction"),
	}
}

func (r *MemoryMetrics) Report() {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)

	r.Alloc.Update(float64(memStats.Alloc))
	r.TotalAlloc.Update(float64(memStats.TotalAlloc))
	r.Sys.Update(float64(memStats.Sys))
	r.Lookups.Update(float64(memStats.Lookups))
	r.Mallocs.Update(float64(memStats.Mallocs))
	r.Frees.Update(float64(memStats.Frees))
	r.HeapAlloc.Update(float64(memStats.HeapAlloc))
	r.HeapSys.Update(float64(memStats.HeapSys))
	r.HeapIdle.Update(float64(memStats.HeapIdle))
	r.HeapInuse.Update(float64(memStats.HeapInuse))
	r.HeapReleased.Update(float64(memStats.HeapReleased))
	r.HeapObjects.Update(float64(memStats.HeapObjects))
	r.StackInuse.Update(float64(memStats.StackInuse))
	r.StackSys.Update(float64(memStats.StackSys))
	r.MSpanInuse.Update(float64(memStats.MSpanInuse))
	r.MSpanSys.Update(float64(memStats.MSpanSys))
	r.MCacheInuse.Update(float64(memStats.MCacheInuse))
	r.MCacheSys.Update(float64(memStats.MCacheSys))
	r.BuckHashSys.Update(float64(memStats.BuckHashSys))
	r.GCSys.Update(float64(memStats.GCSys))
	r.OtherSys.Update(float64(memStats.OtherSys))
	r.NextGC.Update(float64(memStats.NextGC))
	r.LastGC.Update(float64(memStats.LastGC))
	r.PauseTotalNs.Update(float64(memStats.PauseTotalNs))
	r.NumGC.Update(float64(memStats.NumGC))
	r.NumForcedGC.Update(float64(memStats.NumForcedGC))
	r.GCCPUFraction.Update(float64(memStats.GCCPUFraction))
}
