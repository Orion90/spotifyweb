package spotifyweb

import (
	"fmt"
)

//Me is the current logged in user's basic info.
type Me struct {
	DisplayName string `json:"display_name"`
	Country,
	Email,
	Href,
	Id,
	Product,
	Type,
	Uri string
	ExternalUrls map[string]string
	Images       []Image
	Playlists    PlaylistPagingObject `json:"-"`
	Api          SpotifyWeb           `json:"-"`
}

func (api SpotifyWeb) Me() (Me, error) {
	var me Me
	err := api.DoAuth("me", "GET", &me)
	me.Api = api
	me.GetPlayLists()
	return me, err
}

func (me *Me) GetPlayLists() {
	err := me.Api.DoAuth("users/"+me.Id+"/playlists", "GET", &me.Playlists)
	if err != nil {
		fmt.Println(err)
	}
	var items []PlaylistSimple
	for _, pl := range me.Playlists.Items {
		pl.Api = me.Api
		items = append(items, pl)
		if err != nil {
			fmt.Println(err)
		}
	}
	me.Playlists.Items = items
}
