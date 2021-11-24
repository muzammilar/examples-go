// The geomserver is TODO application
package main

import (
	"fmt"
	"time"
)

var (
	version string
	commit  string
)

func main() {
	for {
		fmt.Printf("Hello! This program was compiled on `%s`.\n", commit)
		time.Sleep(30 * time.Second)
	}

}
