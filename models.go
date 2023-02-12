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
	ApplicationName      string
	SlackClientId        string
	SpotifyClientId      string
	SlackRedirectUri     string
	SpotifyRedirectUri   string
	SpotifyState         string
	SlackState           string
	SlackUserExists      bool
	SpotifyUserExists    bool
	UserExists           bool
	UserName             string
	ProfilePicture       string
	CurrentlyListeningTo string
}

type UserModel struct {
	SlackTeamId              string
	SlackUserId              string
	SpotifyUserId            string
	SlackToken               string
	SpotifyToken             string
	SpotifyRefreshToken      string
	SpotifyTokenExpiresIn    int
	UserName                 string
	UserProfilePicture       string
	UserCurrentlyListeningTo string
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

type SpotifyListeningToModel struct {
	Timestamp int64 `json:"timestamp"`
	Context   struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"context"`
	ProgressMs int `json:"progress_ms"`
	Item       struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Resuming bool `json:"resuming"`
		} `json:"disallows"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}
