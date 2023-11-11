package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

var SeedConfig Seed

func LoadSeedConfigFromFile(path string) error {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, &SeedConfig)
}
