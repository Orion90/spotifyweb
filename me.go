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
}

func (api SpotifyWeb) Profile() (Me, error) {
	var me Me
	err := api.DoAuth("me", "GET", &me)
	me.GetPlayLists(api)
	return me, err
}

func (me *Me) GetPlayLists(api SpotifyWeb) {
	err := api.DoAuth("users/"+me.Id+"/playlists", "GET", &me.Playlists)
	if err != nil {
		fmt.Println(err)
	}
	// var items []PlaylistSimple
	// for _, pl := range me.Playlists.Items {
	// 	if !pl.Collaborative {
	// 		continue
	// 	}

	// 	data, err := pl.GetFullTracks(api, me.Id)
	// 	pl.TrackData = data
	// 	items = append(items, pl)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// me.Playlists.Items = items
}
