package spotifyweb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type dataBack struct {
	Token   string `json:"access_token"`
	Type    string `json:"token_type"`
	Expires int    `json:"expires_in"`
	Refresh string `json:"refresh_token"`
}

func (api SpotifyWeb) GetAuthUrl(callbackUrl string) string {
	return "https://accounts.spotify.com/authorize?" + "client_id=" + api.ClientID + "&response_type=code&redirect_uri=" + url.QueryEscape(callbackUrl)
}

func (api SpotifyWeb) GetToken(code, callbackUrl string) (string, string, error) {
	data := url.Values{
		"grant_type":    []string{"authorization_code"},
		"code":          []string{code},
		"redirect_uri":  []string{callbackUrl},
		"client_id":     []string{api.ClientID},
		"client_secret": []string{api.Secret},
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(req)
	var token dataBack
	err = json.NewDecoder(res.Body).Decode(&token)
	return token.Token, token.Refresh, err
}
