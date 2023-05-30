package main

import "github.com/ropehapi/api-go-expert/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
