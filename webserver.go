package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func startWebserver() {
	fileServer := http.FileServer(http.Dir("./html"))
	http.Handle("/", fileServer)

	fmt.Printf("Webserver listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
