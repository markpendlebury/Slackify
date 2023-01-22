package main

import "fmt"

func main() {
	fmt.Println("Starting up...")
	CreateSlackListener()
	CreateSpotifyListener()

	startWebserver()

	//  TODO:
	// Add creation of users, merge the slack and spotify data
	// Together along with a way of identifying the user on
	// both sides
	//
	// Then write this to a cookie and / dynamodb
}
