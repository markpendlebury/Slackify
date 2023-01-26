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
var baseUrl = os.Getenv("BASE_URL")

// This Function creates a http server and awaits for a
// callback response from out request to authenticate
// with slack
func CreateSlackListener() {

	http.HandleFunc("/slack/callback", completeSlackAuth)
	http.HandleFunc("/slack", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
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

	// TODO: validate state code here
	//
	// ------------------------------

	// Now exchange the slackAuthToken for a slackAccessToken:
	slackResponse := exchangeSlackAuthToken(slackAuthToken)

	// Now get the userId of this user:
	slackUser := GetSlackUserId(slackResponse.AccessToken)

	newUser := UserModel{
		SlackUserId:        slackUser.HTTPSSlackComUserID,
		SlackTeamId:        slackUser.HTTPSSlackComTeamID,
		SlackToken:         slackResponse.AccessToken,
		UserName:           fmt.Sprintf("%s %s", slackUser.GivenName, slackUser.FamilyName),
		UserProfilePicture: slackUser.Picture,
	}

	// Lets check if we already have a user with this SlackId:
	existingUser := GetUserBySlackUserId(newUser)
	if len(existingUser.SlackToken) == 0 || len(existingUser.SlackTeamId) == 0 || len(existingUser.SlackUserId) == 0 {
		CreateUser(newUser)
	} else {
		newUser.SpotifyToken = existingUser.SpotifyToken
		newUser.SpotifyUserId = existingUser.SpotifyUserId
		newUser.UserName = fmt.Sprintf("%s %s", slackUser.GivenName, slackUser.FamilyName)
		// UpdateUser(newUser)
	}

	// Lets create some hacky js to close the window:
	js := fmt.Sprintf(`<script type="text/javascript"  charset="utf-8">
		window.location.replace("%s/?slackUserId=%s&slackTeamId=%s");
	</script>`, baseUrl, newUser.SlackUserId, newUser.SlackTeamId)
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

	var slackResponse SlackOpenIdAuthResponse

	_ = json.Unmarshal(body, &slackResponse)

	return slackResponse
}
