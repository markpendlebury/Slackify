package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
)

func OAuth() {
	ctx := context.TODO()

	bot := slack.OAuthResponseBot{
		BotUserID:      "BotUserId",
		BotAccessToken: "BotAccessToken",
	}

	client := &http.Client{}

	token, scope, bot, err := slack.GetBotOAuthTokenContext(ctx, client, "1152673223634.4086809297906", "ecc057a088aa9bf2c4f13b7b3bedb05f", "", "http://localhost:8181")

	fmt.Println(token, scope, err, bot.BotAccessToken)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// const slackProfileEndpoint = "https://slack.com/api/users.profile.set"

// func setSlackStatus(status string, delay int) {

// 	now := time.Now()
// 	output := fmt.Sprintf("[%s] Settings Slack status to: %s", now.Format("15:04:05"), status)

// 	fmt.Println(output)
// 	// Get an expiry time in unix time format
// 	expiry := GetExpiryTime(delay)

// 	// Set our request method
// 	method := "POST"

// 	// Create our payload
// 	payloadString := fmt.Sprintf(`
//   {
// 	  "profile": {
// 		  "status_text": "%s",
// 		  "status_emoji": ":notes:",
// 		  "status_expiration": "%d"
// 	  }
//   }
//   `, status, expiry)

// 	// Create a reader for our payloads
// 	payload := strings.NewReader(payloadString)

// 	// Create a httpclient
// 	client := &http.Client{}

// 	// Create a http request
// 	req, err := http.NewRequest(method, slackProfileEndpoint, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// create our authToken
// 	authToken := fmt.Sprintf("Bearer %s", os.Getenv("SLACK_TOKEN"))

// 	// Add our required request headers
// 	req.Header.Add("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Add("Authorization", authToken)

// 	// use our client to submit the request
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	// Read the response body:
// 	// body, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	var sr SlackResponse

// 	derp := json.NewDecoder(res.Body).Decode(&sr)

// 	if !sr.Ok {
// 		fmt.Println("Error updating slack status: ")

// 		if sr.Error == "ratelimited" {
// 			time.Sleep(30000)
// 		}
// 		fmt.Println(sr.Error)
// 	}

// 	if derp != nil {
// 		fmt.Println(err)
// 	}

// 	// This needs improvement, currently just swalling the response
// 	// We should read the response and handle any errors here:
// 	// fmt.Sprintf("%s", body)

// }

// func GetExpiryTime(delay int) int64 {
// 	// Get epoc time in 5 mins for status expiry:
// 	now := time.Now()

// 	// Convert delay (miliseconds stored in an integer)
// 	// to time.Duration so we can then convert it to
// 	// unix time
// 	delayTime := time.Duration(delay)

// 	// Take now (current time in unix time) and add
// 	// our delaytime (current track time - track length) in miliseconds
// 	// to get an expiry time in unix time for our slack status
// 	expiry := now.Add(delayTime * time.Millisecond)

// 	return expiry.Unix()
// }
