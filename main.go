package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/ohaiibuzzle/archupgrade-go/upgrade_spec"
)

func usage() {
	fmt.Println("usage: archupgradego <filename>")
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

	// We change the working dir into the directory of the file
	// (because it will use relative paths in its includes)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = os.Chdir(filepath.Dir(args[0]))
	if err != nil {
		panic(err)
	}
	defer os.Chdir(dir)

	fp, err := os.Open(filepath.Base(args[0]))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// parse the file into a raw spec first
	spec := upgrade_spec.RawUpgradeSpec{}
	err = yaml.NewDecoder(fp).Decode(&spec)
	if err != nil {
		panic(err)
	}

	// Process inclusions
	real_spec, err := ProcessInclusions(&spec)
	if err != nil {
		panic(err)
	}

	// Dump it
	fmt.Printf("Dumping to stdout\n")
	yaml.NewEncoder(os.Stdout).Encode(real_spec)
}
