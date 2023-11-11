package main

import (
	"github.com/timdufrane/configurationgenerator/pkg/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

const VERSION = "0.1a"

func main() {
	baseConfig := flag.StringP("config", "c", "./config", "base configuration file")
	varConfig := flag.StringP("seedconfig", "s", "./seed.yml", "seed (variable) configuration file")
	outputDir := flag.StringP("directory", "d", "./out", "output directory for writing generated configuration files")
	//verbose := flag.BoolP("verbose", "v", false, "show verbose output")
	version := flag.BoolP("version", "V", false, "show version information")

	help := flag.BoolP("help", "h", false, "show this help message")

	flag.Parse()

	if *help {
		PrintVersion()
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		PrintVersion()
		os.Exit(0)
	}

	if _, err := os.Stat(*baseConfig); err != nil {
		panic("Invalid base configuration file specified")
	}

	if _, err := os.Stat(*varConfig); err != nil {
		panic("Invalid variable configuration file specified")
	}

	baseBytes, err := ioutil.ReadFile(*baseConfig)

	if err != nil {
		panic(fmt.Errorf("Base configuration file could not be read: %v", err))
	}

	err = config.LoadSeedConfigFromFile(*varConfig)

	if err != nil {
		fmt.Errorf("Variable configuration file could not be read: %v", err)
		os.Exit(5)
	}

	log.Println("Generating configs, hang tight")
	configs := config.SeedConfig.GenerateConfigsFromSeed(baseBytes)

	// Create output directory if it doesn't exist
	if _, err := os.Stat(*outputDir); err != nil {
		// Doesn't exist, try to create it
		err = os.MkdirAll(*outputDir, 0644)

		if err != nil {
			os.Exit(8)
		}

		log.Println("Created output directory, as it did not exist")
	}

	for k, v := range configs {
		log.Printf("Writing %d of %d config(s) to files\n", k+1, len(configs))
		ioutil.WriteFile(filepath.Join(*outputDir, fmt.Sprintf("out.%d", k)), []byte(v), 0644)
	}
}

func PrintVersion() {
	fmt.Printf("%s Version %s\n\n", os.Args[0], VERSION)
}
