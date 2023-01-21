package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: Environment Variables
var spotifyRedirectURI = os.Getenv("SPOTIFY_REDIRECT_URI")
var base64EncodedCredentials = os.Getenv("BASE64_ENCODED_SPOTIFY_CREDENTIALS")

// const spotifyState = "fd342dd83b219b5a6f6438b0dd588b12"

// This function creates a http server and awaits for a
// callback response from our request to authenticate
// with spotify
func CreateSpotifyListener() {

	http.HandleFunc("/spotify/callback", completeSpotifyAuth)
	http.HandleFunc("/spotify", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

}

// This function will use the response from the callback to
// authenticate our spotify client
func completeSpotifyAuth(w http.ResponseWriter, r *http.Request) {
	// Lets try to get the token from the response:
	spotifyAuthToken := r.FormValue("code")
	spotifyStateCode := r.FormValue("state")

	if len(spotifyAuthToken) == 0 {
		log.Fatal("Returned authToken was empty")
	}

	if len(spotifyStateCode) == 0 {
		log.Fatal("Returned state was empty")
	}

	// TODO: Validate state code here

	// ------------------------------

	// Now we exchange our authToken for an spotifyResponse:
	spotifyResponse := exchangeSpotifyAuthToken(spotifyAuthToken)

	fmt.Println(spotifyResponse.AccessToken)

	// Lets create some hacky js to close the window:
	js := `<script type="text/javascript"  charset="utf-8">
		window.close();
	</script>`
	// Send the hacky JS back to the responseWriter
	w.Write([]byte(js))
}

func exchangeSpotifyAuthToken(authToken string) SpotifyOpenIdAuthResponse {
	url := "https://accounts.spotify.com/api/token"
	method := "POST"

	payload := strings.NewReader("code=" + authToken + "&redirect_uri=" + spotifyRedirectURI + "&grant_type=authorization_code")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Basic "+base64EncodedCredentials)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	var responseModel SpotifyOpenIdAuthResponse
	_ = json.Unmarshal(body, &responseModel)

	return responseModel
}
