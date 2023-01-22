package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Set the listen port for the main webserver
const port = 1234

// Here we start the main webserver
func startWebserver() {

	// Create a listener for / and send
	//  any requests to the homePage Function
	http.HandleFunc("/", homePage)
	// Start listening
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}

}

// Any requests that hit / (root) will be directed here
func homePage(writer http.ResponseWriter, request *http.Request) {
	// Only accept GET on this route
	if request.Method == "GET" {
		// Read our html file (index.html) into memory
		tmplt, _ := template.ParseFiles("./html/index.html")

		// Create an instance of HtmlContext
		// and populate it with the required data
		// This allows us to avoid storing sensitive data
		// in code (see os.GetEnv("")) <--- we read from the
		// environment
		context := HtmlContext{
			ApplicationName:    "Slackify",
			SlackClientId:      os.Getenv("SLACK_CLIENT_ID"),
			SpotifyClientId:    os.Getenv("SPOTIFY_CLIENT_ID"),
			SlackRedirectUri:   os.Getenv("SLACK_REDIRECT_URI"),
			SpotifyRedirectUri: os.Getenv("SPOTIFY_REDIRECT_URI"),
			SlackState:         "TODO, GENERATE SLACK STATE",
			SpotifyState:       "TODO, GENERATE SPOTIFY STATE",
		}

		// Build the template and write it to the http response
		err := tmplt.Execute(writer, context)

		if err != nil {
			log.Fatal(err)
		}
	}
}
