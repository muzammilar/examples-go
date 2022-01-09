// The module producer contains the code for a sample data producer

package main

const (
	// Producer Types (Async)
	ProducerSync  = iota // 0
	ProducerAsync        // 1
)

const (
	// User ID range
	UserIDMax = 5000 // The max possible value of user ID (not inclusive)
	UserIDMin = 50   // The min possible value of user ID (inclusive)
)

var UserIDRange = UserIDMax - UserIDMin

const (
	// Hashing
	PartitionHash       = "hash"       // compute the hash of the message and send the partition
	PartitionRand       = "rand"       // select a random partition
	PartitionRoundRobin = "roundrobin" // select partition using round robin i.e. (i+1)%n
)

const (
	// Defaults
	DefaultPublishTopic          = "trees"       // The default topic to publish
	DefaultPartitioner           = PartitionHash // The default partitioner for kafka
	DefaultKafkaVersion          = "2.8.1"       //The default kafka version
	DefaultLoggingLevel          = "info"        // The default logging level of the application
	DefaultConnectionBackoffMs   = 5000          // The defualt time in milliseconds before retrying to connect to kafka when creating producers
	DefaultMessageSendIntervalMs = 100           // The default interval between sending messages
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
