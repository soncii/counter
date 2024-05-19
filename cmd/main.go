package main

import (
	"counter/internal/app"
	"flag"
	"fmt"
	"os"
)

func main() {
	numWorkers, filePath := parseFlags()
	app.Run(numWorkers, filePath)
}

func parseFlags() (int, string) {
	numWorkers := flag.Int("w", 1, "number of goroutines to use")
	flag.Parse()
	fmt.Printf("Number of workers: %d\n", *numWorkers)
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide the JSON file path")
		os.Exit(1)
	}

	filePath := flag.Args()[0]
	return *numWorkers, filePath
}
