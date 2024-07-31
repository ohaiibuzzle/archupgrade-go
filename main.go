package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ohaiibuzzle/archupgrade-go/parser"
	"github.com/ohaiibuzzle/archupgrade-go/upgrade_spec"
)

func usage() {
	fmt.Println("usage: archupgrade-go <filename>")
	flag.PrintDefaults()
	os.Exit(2)
}

func ProcessInclusions(spec *upgrade_spec.RawUpgradeSpec) (*upgrade_spec.UpgradeSpec, error) {
	real_spec := &upgrade_spec.UpgradeSpec{
		Version:  spec.Version,
		Upgrade:  spec.Upgrade,
		Finalize: spec.Finalize,
	}

	// If we have inclusion, recursively process them as well
	if len(spec.Includes) != 0 {
		for _, include := range spec.Includes {
			sub_spec := upgrade_spec.RawUpgradeSpec{}

			fp, err := os.Open(include)
			if err != nil {
				return nil, err
			}
			defer fp.Close()

			err = yaml.NewDecoder(fp).Decode(&sub_spec)
			if err != nil {
				return nil, err
			}

			real_sub_spec, err := ProcessInclusions(&sub_spec)
			if err != nil {
				return nil, err
			}

			real_spec.Includes = append(real_spec.Includes, *real_sub_spec)
		}
	}
	return real_spec, nil
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
