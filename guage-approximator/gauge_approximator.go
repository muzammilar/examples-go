package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Metrics Generator dumps a new metric every interval. It may not be very accurate for high frequencies
func MetricGenerator(ch chan<- *Metric, sleepIntvl time.Duration) {
	//
	var metricStats MetricStats
	var metric *Metric

	// seed the number
	rand.Seed(time.Now().UnixNano())

	for {
		// create a metric
		start := time.Now()
		metric = &Metric{
			TS:    time.Now(),
			Value: rand.ExpFloat64(), // expfloat is always positive
		}

		// non blocking write to the channel
		metricStats.TS = metric.TS
		metricStats.Total += 1
		select {
		case ch <- metric:
			metricStats.Sent += 1
		default:
			metricStats.Dropped += 1
		}

		// Print Drop/Sent Metrics every 500 metrics
		if metricStats.Total%500 == 0 {
			fmt.Printf("Metrics Counter: %+v\n", metricStats)
		}

		// offset the channel send interval. This can result in negative duration
		execDuration := time.Since(start)
		time.Sleep(sleepIntvl - execDuration)

	}
}

// get current time and the absolute time bucket
func getCurrentTimeAndBucket(bucketIntvl uint64) (time.Time, uint64) {
	curTime := time.Now()
	curTimeBucket := uint64(curTime.Unix()) / bucketIntvl
	return curTime, curTimeBucket

}

//MetricsAggregator aggregates the metrics and computes the summary and average
func MetricsAggregator(ch <-chan *Metric, buckets uint64, bucketIntvl uint64) {

	// current time bucket
	curTime, curTimeBucket := getCurrentTimeAndBucket(bucketIntvl)
	curArrIdx := curTimeBucket % buckets // aka current array bucket index

	// setup buckets and initialize array
	var gauges []*GaugeSummary = make([]*GaugeSummary, buckets)
	for i := 0; i < int(buckets); i++ {
		gauges[i] = new(GaugeSummary)
		gauges[i].Reset()
	}

	// reads are blocking to decrease busy wait
	for metric := range ch {

		curTime, curTimeBucket = getCurrentTimeAndBucket(bucketIntvl)
		curArrIdx = curTimeBucket % buckets

		// check if we need to create a new metric (i.e. reset the summary) or use the old one
		if curTimeBucket != gauges[curArrIdx].TimeBucket {
			// print the previous bucket only
			fmt.Printf("Sending Gauge Summary: %+v\n",
				gauges[(curArrIdx+buckets-1)%buckets])
			// Reset the metric to clean up any previous state
			gauges[curArrIdx].Reset()
		}
		gauges[curArrIdx].AddMetricDataPoint(curTime, curTimeBucket, metric)

		// print live update to metric
		//fmt.Printf("%+v\n", gauges[curArrIdx])
	}
}

func main() {
	// number of buckets to store historical data
	numBucketsPtr := flag.Int("buckets", 3, "Number of buckets to store historical data")
	intervalSecPtr := flag.Int("interval", 15, "The size of the time interval (in seconds) for which the average is computed, i.e. a single bucket is used.")
	flag.Parse()

	// constants
	var channelSize int = 1000
	generatorSleep, _ := time.ParseDuration("50ms") // skipping error checking for now

	// set up shared metrics channel (buffered)
	var ch chan *Metric
	ch = make(chan *Metric, channelSize)

	// start the metrics aggregator routine
	go MetricsAggregator(ch, uint64(*numBucketsPtr), uint64(*intervalSecPtr)) // skipping error checking for now

	// start metrics generator routine
	go MetricGenerator(ch, generatorSleep)

	// start the signal handler
	exitChan := make(chan int)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			switch <-sigs {
			case syscall.SIGHUP:
				fmt.Println("Signal SIGHUP: Ignoring.")
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Signal Received: Shutting Down.")
				exitChan <- 0
			default:
				fmt.Println("Unknown Signal Caught!")
				exitChan <- 1
			}
		}
	}()

	// wait here for exit code
	exitCode := <-exitChan
	os.Exit(exitCode)
}
