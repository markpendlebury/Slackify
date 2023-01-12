package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const spotifyRedirectURI = "http://localhost:8080/spotify/callback"

const slackClientId = "4089515761430.4096027862115"

const cert = "/home/mpendlebury/Documents/repos/Elesoft/Slackify/Slackify-CLI/localhost.crt"
const key = "/home/mpendlebury/Documents/repos/Elesoft/Slackify/Slackify-CLI/localhost.key"

var (
	spotifyAuth = spotifyauth.New(
		spotifyauth.WithRedirectURL(spotifyRedirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying),
		spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
		spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_CLIENT_SECRET")))
	ch    = make(chan *spotify.Client)
	state = "fd342dd83b219b5a6f6438b0dd588b12"
)

// This function creates a http server and awaits for a
// callback response from our request to authenticate
// with spotify
func CreateSpotifyListener() string {

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

	return spotifyAuth.AuthURL(state)

}

// This Function creates a http server and awaits for a
// callback response from out request to authenticate
// with slack
func CreateSlackListener() string {

	http.HandleFunc("/slack/callback", completeSlackAuth)
	http.HandleFunc("/slack", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServeTLS(":8181", cert, key, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return "https://slack.com/oauth/v2/authorize?user_scope=users.profile:write&client_id=" + slackClientId + "&redirect_uri=https://localhost:8181/slack/callback"
}

// This function opens the request url in our default browser and will either:
// 1. If your browser is already signed into spotify, grab a token and close the tab / window
// 2. If not already signed in, redirect to spotify's authentication mechanism THEN grab a token and close the window
func openbrowser(url string, target string) {
	fmt.Println("I've opened a browser window, you may need to use it to sign into " + target + ". If you're already signed in the window should close automatically")

	// TODO: only show this if debug is enabled
	fmt.Println(url)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// This function will use the response from the callback to
// authenticate our spotify client
func completeSpotifyAuth(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyAuth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(spotifyAuth.Client(r.Context(), token))
	js := `<script type="text/javascript"  charset="utf-8">
				window.close()
	</script>`

	w.Write([]byte(js))

	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

// This function will use the response from the call back to
// authenticate our slack client
func completeSlackAuth(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	headers := r.Header
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(b)
	fmt.Println(headers)

	js := `<script type="text/javascript"  charset="utf-8">
	window.close()
</script>`

	w.Write([]byte(js))

}
