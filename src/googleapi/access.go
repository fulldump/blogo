package googleapi

import (
	"encoding/json"
	"net/http"
)

type GoogleAccess struct {
	AccessToken      string  `json:"access_token"`
	TokenType        string  `json:"token_type"`
	ExpiresIn        int     `json:"expires_in"`
	IdToken          string  `json:"id_token"`
	Error            *string `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

func (ga *GoogleAccess) GetUserInfo() (user_info *GoogleUserInfo, err error) {

	u := "https://www.googleapis.com/oauth2/v2/userinfo"

	req, err := http.NewRequest("GET", u, nil)
	if nil != err {
		return
	}
	req.Header.Set("Authorization", "Bearer "+ga.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if nil != err {
		return
	}

	user_info = &GoogleUserInfo{}
	err = json.NewDecoder(res.Body).Decode(user_info)
	if nil != err {
		return
	}

	return
}
