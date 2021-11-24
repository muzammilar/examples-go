// The geomserver is TODO application
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	version string
	commit  string
)

func main() {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	for {
		fmt.Printf("Hello! This program was compiled on `%s`.\n", commit)
		time.Sleep(30 * time.Second)
	}

}
