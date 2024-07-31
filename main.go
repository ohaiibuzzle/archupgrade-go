package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ohaiibuzzle/archupgrade-go/parser"
)

func usage() {
	// For now it's a simple parser. Do nothing.
	fmt.Println("usage: archupgrade-go <filename>")
	flag.PrintDefaults()
	os.Exit(2)
}
func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	// Simple stuff: Read a file from argv[0] and dump it
	fmt.Printf("Reading from %s\n", args[0])

	spec, err := parser.ParseSpec(args[0])
	if err != nil {
		panic(err)
	}

	// Dump it
	fmt.Printf("Dumping to stdout\n")
	yaml.NewEncoder(os.Stdout).Encode(spec)
}
