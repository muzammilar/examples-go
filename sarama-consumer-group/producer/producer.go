// Package producer contains the code for a sample data producer

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
var supportedPartitioners = []string{PartitionHash, PartitionRand, PartitionRoundRobin}

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
// main.go
/////////////////////////////////

// Sarama configuration options
var (
	brokersStr  = ""
	versionStr  = ""
	topicsStr   = ""
	partitioner = ""
	verbose     = false
	async       = false
	workers     = 1
)

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

func newSyncProducer(brokers string, logger *logrus.Logger) (sarama.SyncProducer, error) {
	config := producerConfig(logger)
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}

func main() { //https://github.com/tcnksm-sample/sarama/blob/master/sync-producer/main.go

	// program flags
	flag.StringVar(&brokersStr, "brokers", "kafka:9094", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.IntVar(&workers, "workers", 1, "Number of producers for the program")
	flag.StringVar(&versionStr, "version", DefaultKafkaVersion, "Kafka cluster version")
	flag.StringVar(&topicsStr, "topics", DefaultPublishTopic, "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&partitioner, "partitioner", DefaultPartitioner, fmt.Sprintf("Producer partition selection strategy. Currently only supports %+v", supportedPartitioners))
	flag.BoolVar(&async, "async", false, "Use an async producer (instead of the default sync producer)")
	flag.BoolVar(&verbose, "verbose", false, "Enable Sarama logging to console")
	flag.Parse()

	// setup logger
	logger := InitLoggerWithStdOut()

	// convert strings to lists
	brokers := strings.Split(brokersStr, ",")
	topics := strings.Split(topicsStr, ",")

	// validate brokers and topics (if needed)

	// start the workers

	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	for i := 0; ; i++ {
		msg := prepareMessage(topic, fmt.Sprintf("Sending Some secret # %d", i))
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("%s error occured.", err.Error())
		} else {
			fmt.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
		}
	}

}
