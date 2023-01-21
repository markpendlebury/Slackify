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
var slackClientId = os.Getenv("SLACK_CLIENT_ID")
var slackClientSecret = os.Getenv("SLACK_CLIENT_SECRET")
var slackRedirectUri = os.Getenv("SLACK_REDIRECT_URI")

// This Function creates a http server and awaits for a
// callback response from out request to authenticate
// with slack
func CreateSlackListener() {

	http.HandleFunc("/slack/callback", completeSlackAuth)
	http.HandleFunc("/slack", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		// err := http.ListenAndServeTLS(":8181", "localhost.crt", "localhost.key", nil)
		err := http.ListenAndServe(":8181", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

// This function will use the response from the call back to
// authenticate our slack client
func completeSlackAuth(w http.ResponseWriter, r *http.Request) {

	slackAuthToken := r.FormValue("code")
	// slackStateCode := r.FormValue("state")

	if len(slackAuthToken) == 0 {
		log.Fatal("Returned authToken was empty")
	}

	// if len(slackStateCode) == 0 {
	// 	log.Fatal("Returned state was empty")
	// }

	// TODO: validate state code here
	//
	// ------------------------------

	// Now exchange the slackAuthToken for a slackAccessToken:

	slackResponse := exchangeSlackAuthToken(slackAuthToken)

	fmt.Println(slackResponse)

	// Lets create some hacky js to close the window:
	js := `<script type="text/javascript"  charset="utf-8">
		window.location.replace("/");
	</script>`
	// Send the hacky JS back to the responseWriter
	w.Write([]byte(js))

}

func exchangeSlackAuthToken(slackAuthToken string) SlackOpenIdAuthResponse {
	url := "https://slack.com/api/openid.connect.token"
	method := "POST"

	payload := strings.NewReader("code=" + slackAuthToken + "&client_id=" + slackClientId + "&client_secret=" + slackClientSecret + "&redirect_uri=" + slackRedirectUri)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
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

	var slackResponse SlackOpenIdAuthResponse

	_ = json.Unmarshal(body, &slackResponse)

	return slackResponse
}
