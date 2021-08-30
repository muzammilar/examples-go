/*
 * Prometheus Stats Server Package
 */

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
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
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
	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_summary_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
	})
)

// Init function of the package
func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

/*
 * Public Functions
 */

func PromServer() {
	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(*addr, nil))

}
