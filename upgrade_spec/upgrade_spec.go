package upgrade_spec

// We need this in order to process inclusions
type RawUpgradeSpec struct {
	Version  string   `yaml:"version"`
	Includes []string `yaml:"includes"`
	Upgrade  Upgrade  `yaml:"upgrade"`
	Finalize Finalize `yaml:"finalize"`
}

type UpgradeSpec struct {
	Version  string        `yaml:"version"`
	Includes []UpgradeSpec `yaml:"includes"`
	Upgrade  Upgrade       `yaml:"upgrade"`
	Finalize Finalize      `yaml:"finalize"`
}

type Upgrade struct {
	Phases []Phase `yaml:"phases"`
}

type Phase struct {
	Name        string    `yaml:"name"`
	Backend     string    `yaml:"backend"`
	Message     string    `yaml:"message"`
	Preinstall  []string  `yaml:"preinstall"`
	Packages    []Package `yaml:"packages"`
	Postinstall []string  `yaml:"postinstall"`
	Reboot      bool      `yaml:"reboot"`
}

type Package struct {
	Url           string `yaml:"url"`
	Hash          string `yaml:"hash"`
	HashAlgorithm string `yaml:"hash-algorithm"`
}

type Finalize struct {
	Shell      []string `yaml:"shell"`
	FileWrite  []File   `yaml:"file_write"`
	FileRemove []string `yaml:"file_remove"`
	CleanCache bool     `yaml:"clean-caches"`
	Reboot     bool     `yaml:"reboot"`
}

type File struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
}
