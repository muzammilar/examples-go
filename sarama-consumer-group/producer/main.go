// Package producer contains the code for a sample data producer

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

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

/////////////////////////////////
// const.go
/////////////////////////////////

const (
	// Producer Types (Async)
	ProducerSync  = iota // 0
	ProducerAsync        // 1
)

const (
	// Hashing
	PartitionHash       = "hash"       // compute the hash of the message and send the partition
	PartitionRand       = "rand"       // select a random partition
	PartitionRoundRobin = "roundrobin" // select partition using round robin i.e. (i+1)%n
)

const (
	// Defaults
	DefaultPublishTopic = "trees"       // The default topic to publish
	DefaultPartitioner  = PartitionHash // The default partitioner for kafka
	DefaultKafkaVersion = "2.8.1"       //The default kafka version
)

// Supported partitioners
var SupportedPartitioners = []string{PartitionHash, PartitionRand, PartitionRoundRobin}

// Tree Names
var Trees = []string{
	"American Beech",
	"American Chestnut",
	"American Elm",
	"American Hophornbeam",
	"American Hornbeam",
	"American Larch",
	"Arborvitae",
	"Balsam Fir",
	"Basswood",
	"Bigtooth Aspen",
	"Bitternut Hickory",
	"Black Ash",
	"Black Birch",
	"Black Cherry",
	"Black Locust",
	"Black Oak",
	"Black Walnut",
	"Black Willow",
	"Butternut",
	"Chestnut Oak",
	"Cucumber Tree",
	"Eastern Cottonwood",
	"Eastern Hemlock",
	"Eastern Redcedar",
	"Eastern White Pine",
	"Gray Birch",
	"Hawthorn",
	"Honey-Locust",
	"Northern Red Oak",
	"Paper Birch",
	"Pignut Hickory",
	"Pin Cherry",
	"Pitch Pine",
	"Quaking Aspen",
	"Red Maple",
	"Red Pine",
	"Red Spruce",
	"Sassafras",
	"Scarlet Oak",
	"Shadbush",
	"Shagbark Hickory",
	"Silver Maple",
	"Slippery Elm",
	"Sugar Maple",
	"Sycamore",
	"The Maples",
	"The Oaks",
	"Tulip Tree",
	"White Ash",
	"White Oak",
	"White Spruce",
	"Yellow Birch",
}

/////////////////////////////////
// message.go
/////////////////////////////////

// Note: JSON serialization is approximately 5-10x slower than binary formats (like protobufs) in my tests

type Message struct {
	Async  int    // Whether the data is sent as sync or async
	UserId int    // This ID is NOT unique between messages and can be repeated
	Data   string // Random tree name
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}

/////////////////////////////////
// logging.go
/////////////////////////////////

// InitLoggerWithFileOutput initializes a logger for a given configuration
func InitLoggerWithStdOut() *logrus.Logger {

	// create a new logger
	var logger *logrus.Logger = logrus.New()

	// set formatting for logger
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// add caller function name (might be somewhat expensive)
	logger.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	logger.SetOutput(os.Stdout)

	// set logging level
	logger.SetLevel(logrus.InfoLevel)

	// return the logger
	return logger
}

/////////////////////////////////
// producer.go
/////////////////////////////////

func startProducer(wg *sync.WaitGroup, ctx context.Context, id int, brokers []string, logger *logrus.Logger) {

	// wait group
	defer wg.Done()

	var asyncproducer sarama.AsyncProducer
	var syncproducer sarama.SyncProducer
	var err error
	// create the producers
	asyncproducer, err = newAsyncProducer(brokers, logger)
	if err != nil {
		logger.Errorf("Worker #%d failed to create an async producer for broker '%+v': %#v", id, brokers, err)
		return
	}
	syncproducer, err = newSyncProducer(brokers, logger)
	if err != nil {
		logger.Errorf("Worker #%d failed to create a sync producer for broker '%+v': %#v", id, brokers, err)
		return
	}

	// start the producers
	var producerWg *sync.WaitGroup = new(sync.WaitGroup)
	producerWg.Add(2) // one for the sync producer and another for the async producer
	// sync producer
	go func() {
		defer producerWg.Done()
	producerLoop:
		for {
			// check context
			select {
			case <-ctx.Done():
				if err := syncproducer.Close(); err != nil {
					logger.Errorf("Worker #%d failed to close syncproducer '%#v'", id, syncproducer)
				}
				break producerLoop
			default:
				// A `default` case is needed to make this select non-blocking
			}
			// send message
			// TODO: FIX
			msg := prepareMessage(topic, fmt.Sprintf("Sending Some secret # %d", i))
			partition, offset, err := syncproducer.SendMessage(msg)
			if err != nil {
				fmt.Printf("%s error occured.", err.Error())
			} else {
				fmt.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
			}
		}
	}()

	// async producer
	go func() {
		defer producerWg.Done()

		producerWg.Add(2) // one for the error channel and one for the succes channel
		// read from the error channel
		go func() {
			defer producerWg.Done()
			for err := range asyncproducer.Errors() {
				logger.Errorf("ERRO!!!!")
			}
		}()

		// read from the success channel
		go func() {
			defer producerWg.Done()
			for s := range asyncproducer.Successes() {
				logger.Infof("Success!")
			}
		}()

	producerLoop:
		for {
			// TODO: FIX
			msg := prepareMessage(topic, fmt.Sprintf("Sending Some secret # %d", i))
			// send message

			// check context
			select { // Do not write a default close otherwise it will become non-blocking
			case <-ctx.Done():
				asyncproducer.AsyncClose() // you can also use async close.
				break producerLoop         // see this to see how to close channels in select https://stackoverflow.com/questions/13666253/breaking-out-of-a-select-statement-when-all-channels-are-closed
			case asyncproducer.Input() <- msg:
				logger.Println("New Message produced")
			}
		}
	}()

	// close the producer when the context is done
	<-ctx.Done()
	// wait for producers to cleanly shut down
	producerWg.Wait()
}

func producerConfig(logger *logrus.Logger) sarama.Config {
	config := sarama.NewConfig()
	// add logging support in verbose mode
	if verbose {
		sarama.Logger = logger
	}
	// select partitioner
	switch partitioner {
	case PartitionHash:
		config.Producer.Partitioner = sarama.NewHashPartitioner
	case PartitionRand:
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	case PartitionRoundRobin:
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	default: // panic should generally be avoided in production, however it's a config check early in the program so it should be okay
		panic(fmt.Sprintf("Unknown Kafka partitoner: %s", partitioner))
	}
	// Kafka Acks
	config.Producer.RequiredAcks = sarama.WaitForLocal // wait for local commit to succeed
	config.Producer.Return.Errors = true               // return the errors via a channel to the user
	config.Producer.Return.Successes = false           // Used by async producer only

	return config
}

func newSyncProducer(brokers []string, logger *logrus.Logger) (sarama.SyncProducer, error) {
	config := producerConfig(logger)
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func newAsyncProducer(brokers []string, logger *logrus.Logger) (sarama.AsyncProducer, error) {
	config := producerConfig(logger)
	producer, err := sarama.NewAsyncProducer(brokers, config)

	return producer, err
}

/////////////////////////////////
// main.go
/////////////////////////////////

// Sarama configuration options
var (
	brokersStr  = ""
	versionStr  = ""
	topic       = ""
	partitioner = ""
	verbose     = false
	workers     = 1
)

func init() {

	// program flags
	flag.StringVar(&brokersStr, "brokers", "kafka:9094", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.IntVar(&workers, "workers", 1, "Number of producers for the program")
	flag.StringVar(&versionStr, "version", DefaultKafkaVersion, "Kafka cluster version")
	flag.StringVar(&topic, "topic", DefaultPublishTopic, "Kafka topic to send data to")
	flag.StringVar(&partitioner, "partitioner", DefaultPartitioner, fmt.Sprintf("Producer partition selection strategy. Currently only supports %+v", SupportedPartitioners))
	flag.BoolVar(&verbose, "verbose", false, "Enable Sarama logging to console")
	flag.Parse()

}

func main() { //https://github.com/tcnksm-sample/sarama/blob/master/sync-producer/main.go

	// setup logger
	logger := InitLoggerWithStdOut()

	// convert strings to lists
	brokers := strings.Split(brokersStr, ",")

	// validate brokers and topics and other input configs (if needed). Skipping since it's a proof-of-concept

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
	logger.Info("Waiting for workers to shutdown.")
	wg.Wait()
}
