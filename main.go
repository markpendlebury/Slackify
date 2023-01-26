package main

import "fmt"

func main() {
	fmt.Println("Starting up...")

	CreateUserTable()

	CreateSlackListener()
	CreateSpotifyListener()

	startWebserver()

}
