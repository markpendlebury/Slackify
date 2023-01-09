# Welcome

Slackify is a local client used to query Spotify and get your currently listening to track and update your status within Slack.

Example: 
![enter image description here](https://user-images.githubusercontent.com/63231900/194503307-cd64595d-1ab8-40dd-aa0a-b55728512d6a.png)

    
# Getting Started:
In order for slackify to update your Slack status you will need to create an application within the slack space you want to update your status in: 

**Instuctions for this process coming soon**

Once you have this setup you'll need to: 
```
export SLACK_TOKEN=xxxxxxxx
```


# Usage
1. Download the correct binary for your architecture from the [releases](https://github.com/markpendlebury/Slackify/releases) page
2. Place the binary in a location accessible by your `$PATH` usually somewhere like `/usr/local/bin` for example
3. run `slackify` from a terminal, it will attempt to open a web browser and you will be required to sign into spotify (if you're not already signed in) 
  If not there's a url provided, simply copy & paste this url into your web browser and sign into slack manually

Start listening via spotify. 

*note*
Because we're using the spotify API, you don't need to listen on the same device as slackify :wink:
  
# Changelog:
- v0.2-alpha 
  - Added readme
  - General improvements to main.go 
  - Added a timestamp output to help debug [this](https://github.com/markpendlebury/Slackify/issues/8) issue
  - Added nil check to GetCurrentlyPlaying
  
- V0.1-alpha 
  - Initial release



# Known issues and bug reporting

Please use githubs [issue tracker](https://github.com/markpendlebury/Slackify/issues) to report any issues/bugs or make any suggestions you may have