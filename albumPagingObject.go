package spotifyweb

type AlbumPagingObject struct {
	Href     string
	Items    []AlbumSimple
	Limit    int
	Next     string
	Offset   int
	Previous string
	Total    int
}
