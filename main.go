package main

import (
	"fmt"
	"goserver/config"
	"goserver/controller"
	"goserver/router"
	"log"
	"net/http"
)

func main() {

	config.InitialMigration()
	controller.Init() // THIS IS TO MAKE SURE INIT FUNCTION IN CONTROLLER RUN BEFRORE
	r := router.GetRouter()

	fmt.Println("Starting server on port 4040")
	err := http.ListenAndServe(":4040", r)
	if err != nil {
		fmt.Println("Failed to start server")
		log.Fatal(err)
		return
	}

	fmt.Println("Server successfully started")
}
