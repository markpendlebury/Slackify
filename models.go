package main

type SlackResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Profile struct {
		Title                 string `json:"title"`
		Phone                 string `json:"phone"`
		Skype                 string `json:"skype"`
		RealName              string `json:"real_name"`
		RealNameNormalized    string `json:"real_name_normalized"`
		DisplayName           string `json:"display_name"`
		DisplayNameNormalized string `json:"display_name_normalized"`
		Fields                struct {
			Xf02V24SGHRQ struct {
				Value string `json:"value"`
				Alt   string `json:"alt"`
			} `json:"Xf02V24SGHRQ"`
			Xf02V2F46P3L struct {
				Value string `json:"value"`
				Alt   string `json:"alt"`
			} `json:"Xf02V2F46P3L"`
			Xf027SL3PGAH struct {
				Value string `json:"value"`
				Alt   string `json:"alt"`
			} `json:"Xf027SL3PGAH"`
			Xf029ANKHNJU struct {
				Value string `json:"value"`
				Alt   string `json:"alt"`
			} `json:"Xf029ANKHNJU"`
			Xf028J1N4945 struct {
				Value string `json:"value"`
				Alt   string `json:"alt"`
			} `json:"Xf028J1N4945"`
		} `json:"fields"`
		StatusText  string `json:"status_text"`
		StatusEmoji string `json:"status_emoji"`
	} `json:"profile"`
}

type SlackOpenIdAuthResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	IDToken     string `json:"id_token"`
}

type SpotifyOpenIdAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type HtmlContext struct {
	ApplicationName    string
	SlackClientId      string
	SpotifyClientId    string
	SlackRedirectUri   string
	SpotifyRedirectUri string
	SpotifyState       string
	SlackState         string
}

type UserModel struct {
	SlackTeamId   string
	SlackUserId   string
	SpotifyUserId string
	SlackToken    string
	SpotifyToken  string
}

type SlackUserModel struct {
	Ok                            bool   `json:"ok"`
	Sub                           string `json:"sub"`
	HTTPSSlackComUserID           string `json:"https://slack.com/user_id"`
	HTTPSSlackComTeamID           string `json:"https://slack.com/team_id"`
	Email                         string `json:"email"`
	EmailVerified                 bool   `json:"email_verified"`
	DateEmailVerified             int    `json:"date_email_verified"`
	Name                          string `json:"name"`
	Picture                       string `json:"picture"`
	GivenName                     string `json:"given_name"`
	FamilyName                    string `json:"family_name"`
	Locale                        string `json:"locale"`
	HTTPSSlackComTeamName         string `json:"https://slack.com/team_name"`
	HTTPSSlackComTeamDomain       string `json:"https://slack.com/team_domain"`
	HTTPSSlackComUserImage24      string `json:"https://slack.com/user_image_24"`
	HTTPSSlackComUserImage32      string `json:"https://slack.com/user_image_32"`
	HTTPSSlackComUserImage48      string `json:"https://slack.com/user_image_48"`
	HTTPSSlackComUserImage72      string `json:"https://slack.com/user_image_72"`
	HTTPSSlackComUserImage192     string `json:"https://slack.com/user_image_192"`
	HTTPSSlackComUserImage512     string `json:"https://slack.com/user_image_512"`
	HTTPSSlackComUserImage1024    string `json:"https://slack.com/user_image_1024"`
	HTTPSSlackComTeamImage34      string `json:"https://slack.com/team_image_34"`
	HTTPSSlackComTeamImage44      string `json:"https://slack.com/team_image_44"`
	HTTPSSlackComTeamImage68      string `json:"https://slack.com/team_image_68"`
	HTTPSSlackComTeamImage88      string `json:"https://slack.com/team_image_88"`
	HTTPSSlackComTeamImage102     string `json:"https://slack.com/team_image_102"`
	HTTPSSlackComTeamImage132     string `json:"https://slack.com/team_image_132"`
	HTTPSSlackComTeamImage230     string `json:"https://slack.com/team_image_230"`
	HTTPSSlackComTeamImageDefault bool   `json:"https://slack.com/team_image_default"`
}

type SpotifyUserModel struct {
	DisplayName  string `json:"display_name"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}
