package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCurrentlyListeningTo(authToken string) string {
	url := "https://api.spotify.com/v1/me/player/currently-playing"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	currentlyListeningTo := SpotifyListeningToModel{}

	_ = json.Unmarshal(body, &currentlyListeningTo)

	currentlyListeningToString := "Nothing"

	if len(currentlyListeningTo.Item.Artists) > 0 {
		if len(currentlyListeningTo.Item.Name) > 0 {
			currentlyListeningToString = fmt.Sprintf("%s - %s", currentlyListeningTo.Item.Artists[0].Name, currentlyListeningTo.Item.Name)
		}
	}

	return currentlyListeningToString

}
