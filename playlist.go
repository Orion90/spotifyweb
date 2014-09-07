package spotifyweb

type PlaylistSimple struct {
	Collaborative bool
	ExternalUrls  map[string]string
	Href          string
	Id            string
	Name          string
	Owner         User
	Public        bool
	Tracks        Tracks
	TrackData     TrackFullPagingObject
	Type          string
	Uri           string
	Api           SpotifyWeb `json:"-"`
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

func (pl *PlaylistSimple) GetFullTracks() (TrackFullPagingObject, error) {
	err := pl.Api.DoAuth("users/"+pl.Owner.Id+"/playlists/"+pl.Id+"/tracks", "GET", &pl.TrackData)

	if err != nil {
		return pl.TrackData, err
	}
	return pl.TrackData, nil
}
