package spotifyweb

type TrackSimple struct {
	Artists     []ArtistSimple
	Disc        int `json:"disc_number"`
	Duration    int `json:"duration_ms"` //The track lengths in milliseconds
	Explcit     bool
	ExternalUrl ExternalUrl
	Href        string
	Id          string
	Name        string
	PreviewUrl  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string
	Uri         string
}
