package config

import (
	"fmt"
	"log"
	"regexp"
)

type Seed struct {
	Parameters map[string][]string `yaml:",inline"`
}

func (s Seed) GenerateConfigsFromSeed(config []byte) []string {
	// Determine largest array size
	var max int
	for _, v := range s.Parameters {
		if len(v) > max {
			max = len(v)
		}
	}

	stringConfig := string(config)

	// Allocate n configurations from passed config
	configs := make([]string, max)

	// Set each to base config
	for i := range configs {
		configs[i] = stringConfig

		searchRegexp := regexp.MustCompile("{{\\s*([^}]+)\\s*}}")

		matches := searchRegexp.FindAllStringSubmatch(configs[i], -1)

		for _, v := range matches {
			// Verify the map object exists
			if _, ok := s.Parameters[v[1]]; ok != true {
				log.Printf("WARNING: Hit a parameter (%s) that wasn't found in the configuration, skipping", v[1])
				continue
			}

			// Replace this config's values with the parameter mapping % i (toroidal in the case of incongruous mappings)
			replacementVal := s.Parameters[v[1]][i%len(s.Parameters[v[1]])]
			replacementRegexp := regexp.MustCompile(fmt.Sprintf("{{\\s*%s\\s*}}", v[1]))

			configs[i] = replacementRegexp.ReplaceAllString(configs[i], replacementVal)
		}
	}

	return configs
}
