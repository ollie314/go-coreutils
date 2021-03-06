// exit program for go-coreutils - fka

// This file should be built to run correctly.
// Because it finds the parent pid to kill.
// If you use `go run` the `go` will be the parent...

package main

import (
	"log"
	"os"
)

// find the parent pid of the current program.
var process = os.Getppid()

func main() {
	proc, err := os.FindProcess(process)
	if err != nil {
		log.Fatalln(err)
	} else {
		proc.Kill()
	}
}
