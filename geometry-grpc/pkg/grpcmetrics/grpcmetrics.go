/*
	Package grpcmetrics contains the prometheus handlers for grpc related metrics
*/

package grpcmetrics

import "github.com/prometheus/client_golang/prometheus"

/*
 * Public - Module Variables storing the metrics
 */

var (
	// RPC Stats
	DataBytes *prometheus.SummaryVec
)

/*
 * Init function
 */

// init is the packages init function that is called by default (the users don't need to call it)
func init() {
	namespace := "grpc"
	DataBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  namespace,
			Name:       "data_bytes",
			Help:       "The amount of bytes recieved/sent by the node",
			Objectives: map[float64]float64{0.0: 0.1, 0.25: 0.1, 0.5: 0.1, 0.75: 0.1, 0.9: 0.01, 0.99: 0.001, 1.0: 0.0001},
		},
		[]string{"direction", "type"}, // direction can be in our out, type is generally 'compressed' or 'uncompressed'
	)
}
