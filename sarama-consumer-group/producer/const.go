package main

const (
	// Hashing
	PartitionHash       = "hash"       // compute the hash of the message and send the partition
	PartitionRand       = "rand"       // select a random partition
	PartitionRoundRobin = "roundrobin" // select partition using round robin i.e. (i+1)%n
	// Defaults
	DefaultPublishTopic = "test"        // The default topic to publish
	DefaultPartitioner  = PartitionHash // The default partitioner for kafka

)
