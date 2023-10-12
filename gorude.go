package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hello [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// Define vars
var (
	targets = flag.String("t", "", "List target host")
	request = flag.String("r", "", "Raw request")
)

func init() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()
}

func main() {

	fmt.Println(*targets)
	fmt.Println(*request)

	fmt.Println("tail:", flag.Args())
}
