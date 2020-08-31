package main

import "boilerplate_golangfront/handlers"

var server = handlers.Server{}

func main() {
	server.Init(":8080")
}
