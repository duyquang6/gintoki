package prometheus

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	CacheEntry = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_entry",
		Help: "Cache entry",
	})

	CacheHitRate = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_hit_rate",
		Help: "Cache hit ratio",
	})

	CacheEvacuate = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_evacuate",
		Help: "Cache hit ratio",
	})

	CacheHit = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_hit",
		Help: "Number of cache hits",
	})

	CacheMiss = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_miss",
		Help: "Number of cache misses",
	})

	ExpiredCache = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "expired_count",
		Help: "Number of expired entry",
	})

	LookupCache = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "lookup_count",
		Help: "Number of lookup entry",
	})
)

// RegisterPrometheus adds the prometheus handler to the mux router
// Note you must register every metric with prometheus for it show up
// when the /metrics route is hit.
func init() {
	prometheus.MustRegister(CacheEntry)
	prometheus.MustRegister(CacheHitRate)
	prometheus.MustRegister(CacheEvacuate)
	prometheus.MustRegister(LookupCache)
	prometheus.MustRegister(ExpiredCache)
	prometheus.MustRegister(CacheMiss)
	prometheus.MustRegister(CacheHit)
}
