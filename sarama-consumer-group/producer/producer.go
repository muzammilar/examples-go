package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

/////////////////////////////////
// const.go
/////////////////////////////////

const (
	// Hashing
	PartitionHash       = "hash"       // compute the hash of the message and send the partition
	PartitionRand       = "rand"       // select a random partition
	PartitionRoundRobin = "roundrobin" // select partition using round robin i.e. (i+1)%n
	// Defaults
	DefaultPublishTopic = "test"        // The default topic to publish
	DefaultPartitioner  = PartitionHash // The default partitioner for kafka
	DefaultKafkaVersion = "2.8.1"       //The default kafka version

)

var supportedPartitioners = []string{PartitionHash, PartitionRand, PartitionRoundRobin}

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
	brokers     = ""
	version     = ""
	group       = ""
	topics      = ""
	partitioner = ""
	verbose     = false
)

func newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal // wait for local commit to succeed
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = false // Used by async producer only
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
	flag.StringVar(&brokers, "brokers", "kafka:9094", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&version, "version", DefaultKafkaVersion, "Kafka cluster version")
	flag.StringVar(&topics, "topics", DefaultPublishTopic, "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&partitioner, "partitioner", DefaultPartitioner, fmt.Sprintf("Producer partition selection strategy. Currently only supports %+v", supportedPartitioners))
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

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
