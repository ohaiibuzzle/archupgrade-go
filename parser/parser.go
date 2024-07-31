package parser

import (
	"os"
	"path/filepath"

	"github.com/ohaiibuzzle/archupgrade-go/upgrade_spec"
	"gopkg.in/yaml.v3"
)

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

func ParseSpec(path string) (*upgrade_spec.UpgradeSpec, error) {
	// Here we change the working dir into the directory of the file
	// (because it will use relative paths in its includes)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = os.Chdir(filepath.Dir(path))
	if err != nil {
		panic(err)
	}
	defer os.Chdir(dir)

	fp, err := os.Open(filepath.Base(path))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

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
	return real_spec, nil
}
