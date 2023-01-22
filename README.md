# Welcome

Slackify is a local client used to query Spotify and get your currently listening to track and update your status within Slack.

This repository is the webserver side of the application, the rest is still in development and will be available soon

Example: <br>
![enter image description here](https://user-images.githubusercontent.com/63231900/211337042-b812ded7-9a24-4d28-b4b9-2a7c63991a19.png)



<br>
<br>


![Build](https://github.com/markpendlebury/Slackify/workflows/Build/badge.svg) ![Release](https://github.com/markpendlebury/Slackify/workflows/Release/badge.svg)




    
# Changelog:
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