package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const slackProfileEndpoint = "https://slack.com/api/users.profile.set"

func GetSlackUserId(authToken string) SlackUserModel {
	url := "https://slack.com/api/openid.connect.userInfo?pretty=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+authToken)

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

	slackUser := SlackUserModel{}

	_ = json.Unmarshal(body, &slackUser)

	if slackUser.Ok {
		return slackUser
	}
	return SlackUserModel{}
}

func SetSlackStatus(status string, delay int, userToken string) {

	now := time.Now()
	output := fmt.Sprintf("[%s] Settings Slack status to: %s", now.Format("15:04:05"), status)

	fmt.Println(output)
	// Get an expiry time in unix time format
	expiry := GetExpiryTime(delay)

	// Set our request method
	method := "POST"

	// Create our payload
	payloadString := fmt.Sprintf(`
  {
	  "profile": {
		  "status_text": "%s",
		  "status_emoji": ":notes:",
		  "status_expiration": "%d"
	  }
  }
  `, status, expiry)

	// Create a reader for our payloads
	payload := strings.NewReader(payloadString)

	// Create a httpclient
	client := &http.Client{}

	// Create a http request
	req, err := http.NewRequest(method, slackProfileEndpoint, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	// create our authToken
	authToken := fmt.Sprintf("Bearer %s", userToken)

	// Add our required request headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", authToken)

	// use our client to submit the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response body:
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	var sr SlackResponse

	derp := json.NewDecoder(res.Body).Decode(&sr)

	if !sr.Ok {
		fmt.Println("Error updating slack status: ")

		if sr.Error == "ratelimited" {
			time.Sleep(30000)
		}
		fmt.Println(sr.Error)
	}

	if derp != nil {
		fmt.Println(err)
	}

	// This needs improvement, currently just swallowing the response
	// We should read the response and handle any errors here:
	// fmt.Sprintf("%s", body)

}

func GetExpiryTime(delay int) int64 {
	// Get epoc time in 5 mins for status expiry:
	now := time.Now()

	// Convert delay (miliseconds stored in an integer)
	// to time.Duration so we can then convert it to
	// unix time
	delayTime := time.Duration(delay)

	// Take now (current time in unix time) and add
	// our delaytime (current track time - track length) in miliseconds
	// to get an expiry time in unix time for our slack status
	expiry := now.Add(delayTime * time.Millisecond)

	return expiry.Unix()
}
