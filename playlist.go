package spotifyweb

type PlaylistSimple struct {
	Collaborative bool
	ExternalUrls  map[string]string
	Href          string
	Id            string
	Name          string
	Owner         interface{}
	Public        bool
	Tracks        Tracks
	TrackData     TrackFullPagingObject
	Type          string
	Uri           string
}

type PlaylistPagingObject struct {
	Href     string
	Items    []PlaylistSimple
	Limit    int
	Next     string
	Offset   int
	Previous string
	Total    int
}

func (pl *PlaylistSimple) GetFullTracks(api SpotifyWeb, user string) (TrackFullPagingObject, error) {
	err := api.DoAuth("users/"+user+"/playlists/"+pl.Id+"/tracks", "GET", &pl.TrackData)

	if err != nil {
		return pl.TrackData, err
	}
	return pl.TrackData, nil
}
