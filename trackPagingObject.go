package spotifyweb

type TrackPagingObject struct {
	Href     string
	Items    []TrackSimple
	Limit    int
	Next     string
	Offset   int
	Previous string
	Total    int
}
