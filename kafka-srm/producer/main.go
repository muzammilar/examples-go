// The module producer contains the code for a sample data producer

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Sarama configuration options
var (
	brokersStr  = ""
	versionStr  = ""
	topic       = ""
	partitioner = ""
	verbose     = false
	workers     = 1
	loglevel    = ""
)

func init() {

	// program flags
	flag.StringVar(&brokersStr, "brokers", "kafka:9094", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.IntVar(&workers, "workers", 1, "Number of producers for the program")
	flag.StringVar(&versionStr, "version", DefaultKafkaVersion, "Kafka cluster version")
	flag.StringVar(&topic, "topic", DefaultPublishTopic, "Kafka topic to send data to")
	flag.StringVar(&partitioner, "partitioner", DefaultPartitioner, fmt.Sprintf("Producer partition selection strategy. Currently only supports %+v", SupportedPartitioners))
	flag.BoolVar(&verbose, "log.sarama", false, "Enable Sarama logging to console (as Printf)")
	flag.StringVar(&loglevel, "log.level", DefaultLoggingLevel, "The logging level for the program (except sarama logs).")
	flag.Parse()

}

func main() { //https://github.com/tcnksm-sample/sarama/blob/master/sync-producer/main.go

	// setup logger
	logger := InitLoggerWithStdOut()

	// convert strings to lists
	brokers := strings.Split(brokersStr, ",")

	// validate brokers and topics and other input configs (if needed). Skipping since it's a proof-of-concept

	// start the metrics handler
	startMetricsCollector()

	// create a context
	ctx, cancel := context.WithCancel(context.Background())

	// setup the workers
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(workers)

	// start the workers
	for i := 0; i < workers; i++ {
		go startProducer(wg, ctx, i, brokers, logger)
	}

	// start the signal handler with context
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select { // Do not specify a `default` case for this select, since a `default` case makes the code non-blocking (and requires a loop)
	case <-ctx.Done():
		logger.Info("Terminating: context cancelled")
	case sig := <-sigterm:
		logger.Infof("Terminating: via signal %s", sig)
	}
	// cancel the context on signal
	cancel()

	// wait for the workers to finish
	logger.Info("Waiting for workers to shutdown")
	wg.Wait()
	logger.Info("Shutdown successful")
}
