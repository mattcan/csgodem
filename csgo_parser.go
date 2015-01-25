package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("f", "", "Name of the demo file")

	flag.Parse()

	fmt.Println(*filename)

	// use DemoFileDump to open
	var demo DemoFile
	if demo.Open(*filename) {
		// run the dump routine from DemoFileDump
		fmt.Println("Demo file is good")
	}
}
