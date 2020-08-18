package main

import (

)

var router = handlers.Router

func main() {

	/*
		var err error
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error accessing environment variables")
		}
	*/

	router.Init(":80")

}