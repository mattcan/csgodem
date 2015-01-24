package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("f","", "Name of the demo file")

	flag.Parse()

	fmt.Println(*filename)
}
