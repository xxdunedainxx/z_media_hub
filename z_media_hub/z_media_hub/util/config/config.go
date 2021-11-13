package config

import (
		"os"
    "fmt"
    "encoding/json"
)

type Config struct {
    Kasa struct {
    	KasaEnabled bool `json:"kasaEnabled"`
    	Endpoints []string `json:"endpoints"`
    }`json:"kasa"`
}

func LoadConfiguration() Config {
    var config Config
    configFile, err := os.Open("conf.json")
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}