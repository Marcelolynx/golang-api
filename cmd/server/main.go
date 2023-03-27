package main

import config "github.com/marcelolynx/golang-api/configs"

func main() {
	//...
	config, _ := config.LoadConfig(".")
	println(config.DBDriver)
}
