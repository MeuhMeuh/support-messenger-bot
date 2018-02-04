package main

import (
	"encoding/json"
	"os"
)

// Configuration handles the whole configuration of the project.
type Configuration struct {
	token []string
}

func getConfiguration() (Configuration, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	return configuration, err
}
