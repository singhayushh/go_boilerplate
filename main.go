package main

import "boilerplate_golangfront/handlers"

var server = handlers.Server{}

func main() {

	/*
		var err error
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error accessing environment variables")
		}
	*/

	server.Init(":80")

}
