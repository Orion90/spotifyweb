package spotifyweb

type TrackSimple struct {
	Artists     []ArtistSimple
	Disc        int `json:"disc_number"`
	Duration    int `json:"duration_ms"` //The track lengths in milliseconds
	Explcit     bool
	ExternalUrl map[string]string `json:"external_urls"`
	Href        string
	Id          string
	Name        string
	PreviewUrl  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string
	Uri         string
}
type TrackPlaylist struct {
	AddedAt string
	AddedBy interface{}
	Track   TrackFull
}
type TrackFull struct {
	Album       AlbumSimple
	Artists     []ArtistSimple
	Markets     []string `json:available_markets"`
	Disc        int      `json:"disc_number"`
	Duration    int      `json:"duration_ms"` //The track lengths in milliseconds
	Explcit     bool
	ExternalUrl map[string]string `json:"external_urls"`
	Href        string
	Id          string
	Name        string
	PreviewUrl  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string
	Uri         string
}
type Tracks struct {
	Href  string
	Total int
}
