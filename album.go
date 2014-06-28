package spotifyweb

/*
images | array of image objects | The cover art for the album in various sizes, widest first.
name | string | The name of the album.
popularity | integer | The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
release_date | string | The date the album was first released, for example "1981-12-15". Depending on the precision, it might be shown as "1981" or "1981-12".
release_date_precision | string | The precision with which release_date value is known: "year", "month", or "day".
tracks | array of simplified track objects inside a paging object | The tracks of the album.
type | string | The object type: "album"
uri	| string | The Spotify URI for the album.
*/
//AlbumContainer holds information on an album
type AlbumContainer struct {
	AlbumType  string             `json:"album_type"`        //The type of the album: one of "album", "single", or "compilation".
	Artists    []ArtistSimplified `json:"artists"`           //The artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	Markets    []string           `json:"available_markets"` //The markets in which the album is available: ISO 3166-1 alpha-2 country codes. Note: Album is available when at least 1 of its tracks is available in that market.
	ExternalID ExternalID         `json:"external_ids"`      //Known external IDs for the album.
	ExternalUr ExternalUrl        `json:"external_urls"`     //Known external URLs for this album.
	Genres     []string           `json:"genres"`            //A list of the genres used to classify the album. For example: "Prog Rock", "Post-Grunge". (If not yet classified, the array is empty.)
	Link       string             `json:"href"`              //A link to the Web API endpoint providing full details of the album.
	ID         string             `json:"id"`                //The Spotify ID for the album.
	Images     []Image            `json:"images"`            //The cover art for the album in various sizes, widest first.
}

func (api *SpotifyWeb) Album(id string) (AlbumContainer, error) {
	var album AlbumContainer
	err := api.DoBasic("albums/"+id, "GET", &album)
	return album, err
}
