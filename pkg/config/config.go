package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	gcfg "gopkg.in/gcfg.v1"
)

// Init Read and process config file
func Init() AppConfig {

	appconfig := AppConfig{}
	ok := readConfig(&appconfig, "./development", "stockbit")
	if !ok {
		log.Fatal("Failed to read config file")
	}
	return appconfig
}

// readConfig is file handler for reading configuration files into variable
// Return: - boolean
func readConfig(ac *AppConfig, path string, module string) bool {
	parts := []string{"main", "vendor"}
	var configString []string

	for _, v := range parts {
		fname := path + "/" + module + "." + v + ".ini"
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Reading", fname)

		config, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Println("function readConfig", err)
			return false
		}

		configString = append(configString, string(config))
	}

	err := gcfg.ReadStringInto(ac, strings.Join(configString, "\n\n"))
	if err != nil {
		log.Println("function readConfig", err)
		return false
	}

	return true
}
