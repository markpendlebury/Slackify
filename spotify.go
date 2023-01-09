package main

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

func GetCurrentlyPlaying(client *spotify.Client) *spotify.CurrentlyPlaying {

	// Use our spotify client to get the currently playing track data:
	if client != nil {
		currentlyPlaying, err := client.PlayerCurrentlyPlaying(context.Background())
		if err != nil {
			println(err.Error())
		}

		return currentlyPlaying
	}

	return nil

}
