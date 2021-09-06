/*
 * Prometheus Stats Server Package
 */

// This package is not very well designed but is as such to avoid making major changes to the initial example

package promstats

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
 * Constants
 */

var (
	RpcDurations          *prometheus.SummaryVec
	RpcDurationsHistogram *prometheus.Histogram
)

// Custom init function of the package that can takes in some arguments
func Init(normMean float64, normDomain float64) {
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	RpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)
	// The same as above, but now as a summary, and only for the normal
	// distribution. The buckets are targeted to the parameters of the
	// normal distribution, with 20 buckets centered on the mean, each
	// half-sigma wide.
	RpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_summary_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(normMean-5*normDomain, .5*normDomain, 20),
	})

	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(RpcDurations)
	prometheus.MustRegister(RpcDurationsHistogram)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

/*
 * Public Functions
 */

// PromServer takes an address of a server (as a string) and starts serving a prometheus server on that address
func PromServer(addr string) {
	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(addr, nil))

}
