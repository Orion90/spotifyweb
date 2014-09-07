//Package spotifyweb provides a GO wrapper for Spotify's Web API (https://developer.spotify.com/web-api).
package spotifyweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//SpotifyWeb provides setup for the app utilizing the API.
type SpotifyWeb struct {
	Endpoint,
	ClientID,
	Secret,
	Token, Refresh string
}

//DoBasic provides basic API calls, no access token required.
func (api *SpotifyWeb) DoBasic(call, method string, result interface{}) error {
	url := api.Endpoint + call
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func (api *SpotifyWeb) DoAuth(call, method string, result interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, api.Endpoint+call, nil)
	if err != nil {
		fmt.Println("err1")
		return err
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err2")
		return err
	}
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		fmt.Println("err3")
		return err
	}
	return nil
}
