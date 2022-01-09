package main

import (
	"fmt"
	"time"
)

func main() {
	// print messages
	for {
		fmt.Printf("Hey There!\n")
		time.Sleep(5 * time.Second)
	}
}

// The consumer reads json as the following
var data map[string]interface{}
