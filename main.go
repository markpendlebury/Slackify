package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying),
		spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
		spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_CLIENT_SECRET")))

	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	openbrowser(url)

	// wait for auth to complete
	client := <-ch
	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Logged in as: ", user.DisplayName)

	for {
		currentlyListeningTo, err := client.PlayerCurrentlyPlaying(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		if currentlyListeningTo.Item != nil {
			status := fmt.Sprintf("Currently Listening to: %s - %s", currentlyListeningTo.Item.Artists[0].Name, currentlyListeningTo.Item.Name)
			fmt.Println(status)
			setSlackStatus(status)
		}

		time.Sleep(60 * time.Second)
	}

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	js := `<script type="text/javascript"  charset="utf-8">
				window.close()
	</script>`

	w.Write([]byte(js))

	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func openbrowser(url string) {
	fmt.Println("Opening: ", url)
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

func setSlackStatus(status string) {

	// Get epoc time in 5 mins for status expiry:
	now := time.Now()

	expiry := now.Add(300 * time.Second)

	url := "https://slack.com/api/users.profile.set"
	method := "POST"

	payloadString := fmt.Sprintf(`
  {
	  "profile": {
		  "status_text": "%s",
		  "status_emoji": ":musical_note:",
		  "status_expiration": "%d"
	  }
  }
  `, status, expiry.Unix())

	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	authToken := fmt.Sprintf("Bearer %s", os.Getenv("SLACK_TOKEN"))

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", authToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Sprintf("%s", body)

}
