# Welcome

Slackify is a collection of web and backend microsservices (overengineerd yet i know) used to query Spotify and get your currently listening to track and update your status within Slack.

This repository is the webserver side of the application, the rest is still in development and will be available soon

Example: 
<br>
<p align="center">
<!-- ![Slack profile preview](https://user-images.githubusercontent.com/63231900/211337042-b812ded7-9a24-4d28-b4b9-2a7c63991a19.png) -->
<img src="https://user-images.githubusercontent.com/63231900/211337042-b812ded7-9a24-4d28-b4b9-2a7c63991a19.png">
</p>
<p align="center">
<img src="https://user-images.githubusercontent.com/63231900/213915058-24630098-eabb-461a-8b7e-fa5b57c8c994.png">
</p>
<br>

<p align="center">
  <img src="https://github.com/markpendlebury/Slackify/workflows/Build/badge.svg">
  <img src="https://github.com/markpendlebury/Slackify/workflows/Release/badge.svg">
</p>


# Getting started

Before you start you will need to set the following environment variables: 

```
SLACK_CLIENT_ID="YOUR_SLACK_APP_CLIENT_ID_HERE",
SPOTIFY_CLIENT_ID="YOUR_SPOTIFY_APP_CLIENT_ID_HERE",
SLACK_REDIRECT_URI="https://localhost:8080/slack/callback",
SPOTIFY_REDIRECT_URI="http://localhost:8181/spotify/callback",
BASE64_ENCODED_SPOTIFY_CREDENTIALS="YOUR_SPOTIFY_CLIENT_ID_AND_SECRET_TOGETHER_SEPERATED_BY_A_:_BASED64_ENCODED"
SLACK_CLIENT_SECRET="YOUR_SLACK_APP_CLIENT_SECRET"
AWS_ACCESS_KEY_ID="YOUR_AWS_ACCESS_KEY_ID"
AWS_SECRET_ACCESS_KEY="YOUR_AWS_SECRET_ACCESS_KEY"
AWS_REGION="THE_REGION_OF_YOUR_DYNAMO_DB"
```

You will also need an aws account and a user configured with cli access (access key / secret etc) and access to a dynamodb instance in your preferred region.


When the web application starts it will actually start 3 webservers: 
- The main webserver listening on `1234` 
- The Spotify callback handler listening on port `8080`
- The Slack callback handler listening on port `8181`


### Docker

```
git clone git@github.com:markpendlebury/Slackify.git slackify-ui
cd slackify-ui
docker build -t slackify-ui:latest .
docker run -p 1234:1234 -p 8080:8080 -p 8181 slacikfy-ui:latest
```

## Compiled from source
```
git clone git@github.com:markpendlebury/Slackify.git slackify-ui
cd slackify-ui
go build .
./slackify
```

    
# Changelog:
- v0.5-alpha
  - Exanding on the user model to contain more personalised information such as username / currently listening to
  - Added mechanisms and updated auth flow to grab required data during the user auth flow process
  - Updated index.html to contain logic to only show parts of the auth flow that are required (when data is missing for example) 
  - Mapping (some of) this data to data pulled from slack / spotify
  
- v0.4-alpha
  - Implementing DynamoDB functionality
  - Created a UserModel to store in the db
  - Implemented a (first attempt) flow for requesting and storing the UserModel
  
- v0.3.1-alpha
  - Fixed a small templating issue

- v0.3-alpha
  - Refactored the entire application to now serve over html rather than a console application
  - Adding templating mechanism to help keep secrets out of html (clientid/secrets etc) 
  - Added a basic dockerfile to serve the app 
  - Moved slack and spotify related listeners and auth flows to their respective files

- v0.2-alpha 
  - Added readme
  - General improvements to main.go 
  - Added a timestamp output to help debug [this](https://github.com/markpendlebury/Slackify/issues/8) issue
  - Added nil check to GetCurrentlyPlaying
  
- V0.1-alpha 
  - Initial release



# Known issues and bug reporting

Please use githubs [issue tracker](https://github.com/markpendlebury/Slackify/issues) to report any issues/bugs or make any suggestions you may have
