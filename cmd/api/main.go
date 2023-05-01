package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {}


func main(){
	app := Config{}
	
	srv := &http.Server{
		Addr: ":80",
		Handler: app.routes(),
	}

	ServerPort := 80

	fmt.Printf("Starting server on port %v", ServerPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}