package main

import (
	"os"
	
	"github.com/pelletier/go-toml/v2"
)

var cfg ConfigFile

func getConfigFile() {
	LH := getLoggerHandler("ConfigDecoder")
	LH.Info("ConfigDecoder package initialized")
	defer LH.Sync()
	_, err := os.Stat("config/config.toml"); isConfigFileExist := os.IsNotExist(err)
	if (isConfigFileExist) {
		file, rErr := os.ReadFile("config/config.toml")
		if rErr != nil {
			LH.Fatal("Error reading config file: " + rErr.Error())
		}
		// decode config file
		dErr := toml.Unmarshal(file, &cfg)
		if dErr != nil {
			LH.Fatal("Error unmarshaling config file: " + dErr.Error())
		}
	}else{
		LH.Fatal("Config file does not exist.")
	}
	LH.Info("Config file loaded successfully.")
}