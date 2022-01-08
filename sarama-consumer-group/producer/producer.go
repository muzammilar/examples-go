package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var brokers = []string{"kafka:9094"}

func newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal // wait for local commit to succeed
	config.Producer.Return.Successes = true
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
