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
