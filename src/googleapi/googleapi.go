package googleapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type GoogleApi struct {
	ClientId     string `usage:"Google Client ID"`
	ClientSecret string `usage:"Google Client Secret"`
	RedirectUri  string `usage:"Google Redirect URI"`
}

/*
Following these steps:
https://developers.google.com/identity/protocols/OAuth2WebServer#obtainingaccesstokens

Example:
https://accounts.google.com/o/oauth2/v2/auth?
scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fdrive.metadata.readonly&
access_type=offline&
include_granted_scopes=true&
state=state_parameter_passthrough_value&
redirect_uri=http%3A%2F%2Foauth2.example.com%2Fcallback&
response_type=code&
client_id=client_id
*/
func (ga *GoogleApi) CreateLink(state string) string {
	return "https://accounts.google.com/o/oauth2/v2/auth?" +
		"client_id=" + url.QueryEscape(ga.ClientId) +
		"&redirect_uri=" + url.QueryEscape(ga.RedirectUri) +
		"&scope=" + url.QueryEscape("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile") +
		//"&access_type=" + url.QueryEscape(access_type) +
		//"&include_granted_scopes=" + url.QueryEscape(include_granted_scopes) +
		"&state=" + url.QueryEscape(state) +
		"&prompt=select_account" +
		"&response_type=" + url.QueryEscape("code")
}

func (ga *GoogleApi) CreateLinkWithHost(state, host string) string {

	r, err := url.Parse(ga.RedirectUri)
	if nil != err {
		panic(err)
	}

	r.Host = host

	return "https://accounts.google.com/o/oauth2/v2/auth?" +
		"client_id=" + url.QueryEscape(ga.ClientId) +
		"&redirect_uri=" + url.QueryEscape(r.String()) +
		"&scope=" + url.QueryEscape("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile") +
		//"&access_type=" + url.QueryEscape(access_type) +
		//"&include_granted_scopes=" + url.QueryEscape(include_granted_scopes) +
		"&state=" + url.QueryEscape(state) +
		"&prompt=select_account" +
		"&response_type=" + url.QueryEscape("code")
}

/*
To exchange an authorization code for an access token, call the https://www.googleapis.com/oauth2/v4/token endpoint and set the following parameters:

Fields:
code	The authorization code returned from the initial request.
client_id	The client ID obtained from the API Console.
client_secret	The client secret obtained from the API Console.
redirect_uri	One of the redirect URIs listed for your project in the API Console.
grant_type	As defined in the OAuth 2.0 specification, this field must contain a value of authorization_code.

Response:
access_token: // string
token_type: // string
expires_in: // int
id_token: // string
error: // string
error_description: // string
*/
func (ga *GoogleApi) GetAccessToken(code string) (access *GoogleAccess, err error) {
	endpoint := "https://www.googleapis.com/oauth2/v4/token"

	payload := "code=" + url.QueryEscape(code) +
		"&client_id=" + url.QueryEscape(ga.ClientId) +
		"&client_secret=" + url.QueryEscape(ga.ClientSecret) +
		"&redirect_uri=" + url.QueryEscape(ga.RedirectUri) +
		"&grant_type=" + url.QueryEscape("authorization_code")

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(payload))
	if nil != err {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if nil != err {
		return
	}

	access = &GoogleAccess{}
	if err = json.NewDecoder(res.Body).Decode(&access); nil != err {
		return
	}

	if nil != access.Error {
		err = errors.New(*access.Error + ":" + access.ErrorDescription)
		return
	}

	return
}

func (ga *GoogleApi) GetAccessTokenWithHost(code, host string) (access *GoogleAccess, err error) {

	r, err := url.Parse(ga.RedirectUri)
	if nil != err {
		panic(err)
	}

	r.Host = host

	endpoint := "https://www.googleapis.com/oauth2/v4/token"

	payload := "code=" + url.QueryEscape(code) +
		"&client_id=" + url.QueryEscape(ga.ClientId) +
		"&client_secret=" + url.QueryEscape(ga.ClientSecret) +
		"&redirect_uri=" + url.QueryEscape(r.String()) +
		"&grant_type=" + url.QueryEscape("authorization_code")

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(payload))
	if nil != err {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if nil != err {
		return
	}

	access = &GoogleAccess{}
	if err = json.NewDecoder(res.Body).Decode(&access); nil != err {
		return
	}

	if nil != access.Error {
		err = errors.New(*access.Error + ":" + access.ErrorDescription)
		return
	}

	return
}
