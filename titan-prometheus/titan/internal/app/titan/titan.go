/*
 * Titan Oscillator Metrics Generator
 */

package titan

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
 * Constants
 */

/*
 * Public Functions
 */

func StartTitan(oscillationPeriod time.Duration) {
	start := time.Now()

	oscillationFactor := getOscillationFactorFunc(start, oscillationPeriod)

	// There defer function is not needed and is only used as an example
	defer func() {
		fmt.Println("All Titan Oscillators have been started.")
	}()

	// Prometheus Example: Periodically record some sample latencies for the three services.
	go exponentialOscillator([]string{"titan", "exponential"}, oscillationFactor)
	go uniformOscillator([]string{"titan", "uniform"}, oscillationFactor)
	go normalOscillator([]string{"titan", "normal"}, oscillationFactor)

	// gokit-metrics prometheus example

	// gokit-metrics statsd example

}

/*
 * Private Functions
 */

func exponentialOscillator(labels []string, oscillationFactor func() float64) {
	for {
		v := rand.ExpFloat64() / 1e6
		rpcDurations.WithLabelValues(labels...).Observe(v)
		time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
	}
}

func uniformOscillator(labels []string, oscillationFactor func() float64) {
	v := rand.Float64() * *uniformDomain
	rpcDurations.WithLabelValues("uniform").Observe(v)
	time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)

}

func normalOscillator(labels []string, oscillationFactor func() float64) {
	for {
		v := (rand.NormFloat64() * *normDomain) + *normMean
		rpcDurations.WithLabelValues("normal").Observe(v)
		// Demonstrate exemplar support with a dummy ID. This
		// would be something like a trace ID in a real
		// application.  Note the necessary type assertion. We
		// already know that rpcDurationsHistogram implements
		// the ExemplarObserver interface and thus don't need to
		// check the outcome of the type assertion.
		rpcDurationsHistogram.(prometheus.ExemplarObserver).ObserveWithExemplar(
			v, prometheus.Labels{"dummyID": fmt.Sprint(rand.Intn(100000))},
		)
		time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
	}

}

func getOscillationFactorFunc(start time.Time, oscillationPeriod time.Duration) func() float64 {
	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(oscillationPeriod)))
	}
	return oscillationFactor
}
