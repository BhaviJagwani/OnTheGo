package main

import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    LogFilePath  string
	ServerHost   string
	ServerPort   string
}

func LoadConfiguration() Configuration {
	file, _ := os.Open("config/default.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}