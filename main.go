package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	listenUrl := CreateListener()

	openbrowser(listenUrl)

	// Await auth completion
	client := <-ch

	// Get the current logged in user
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Show who we're logged in as:
	fmt.Println("Logged in as: ", user.DisplayName)

	// Begin listening for track changes:

	for {
		// initialie delay in miliseconds
		var delay int = 30000

		// Get currentlyPlaying data from spotify:
		currentlyPlaying := GetCurrentlyPlaying(client)

		if currentlyPlaying.Item != nil {
			// Create a formatted status update
			status := fmt.Sprintf("Currently Listening to: %s - %s", currentlyPlaying.Item.Artists[0].Name, currentlyPlaying.Item.Name)

			// Get the progress (current timestamp) of the track we're listening to:
			progress := currentlyPlaying.Progress
			// Get the full length of the track we're listening to:
			length := currentlyPlaying.Item.Duration

			// Get the time remaining of the track in miliseconds
			delay = length - progress + 3

			// Output to terminal our current status, here \r ensures only use a single line
			fmt.Printf("\r " + status)

			// Send our status to slack:
			setSlackStatus(status, delay)
		} else {
			// We're not listening to anything, Tell the user;
			fmt.Printf("\r Nothing playing...")
		}

		// Here we sleep for the above calculated delay plus 2 seconds
		// This helps reduce the chances of hitting spotify's rate limits
		// whilst at the same time keeping the slack updates as accurage
		// as possible, open to ideas for a better solution to this:
		time.Sleep(time.Duration(delay)*time.Millisecond + time.Duration(2))
	}
}
