package main

var cfg ConfigFile

func init()	{
	LH := getLoggerHandler("Config")
	cfg.DatabaseHost     = "localhost"   // Edited value
	cfg.DatabasePassword = "postgres"    // Edited value
	cfg.DatabaseUsername = "postgres"    // Edited value
	cfg.ServerName       = "MyServer"    // Edited value
	cfg.ServerPort       = 8080          // Edited value
	LH.Info("Config package initialized with default values.")
	defer LH.Sync()
}