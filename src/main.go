package main

func main() {
	print("F**k Go Package")
	getConfigFile()
	println(cfg.DatabaseHost)
	println(cfg.DatabasePassword)
	println(cfg.DatabaseUsername)
	println(cfg.ServerName)
	println(cfg.ServerPort)
}