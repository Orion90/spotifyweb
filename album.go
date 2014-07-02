package spotifyweb

import (
	"strings"
)

type MultiAlbumContainer struct {
	Albums []AlbumContainer
}

//AlbumContainer holds information on an album
type AlbumContainer struct {
	AlbumType        string            `json:"album_type"`             //The type of the album: one of "album", "single", or "compilation".
	Artists          []ArtistSimple    `json:"artists"`                //The artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	Markets          []string          `json:"available_markets"`      //The markets in which the album is available: ISO 3166-1 alpha-2 country codes. Note: Album is available when at least 1 of its tracks is available in that market.
	ExternalID       map[string]string `json:"external_ids"`           //Known external IDs for the album.
	ExternalUr       map[string]string `json:"external_urls"`          //Known external URLs for this album.
	Genres           []string          `json:"genres"`                 //A list of the genres used to classify the album. For example: "Prog Rock", "Post-Grunge". (If not yet classified, the array is empty.)
	Link             string            `json:"href"`                   //A link to the Web API endpoint providing full details of the album.
	ID               string            `json:"id"`                     //The Spotify ID for the album.
	Images           []Image           `json:"images"`                 //The cover art for the album in various sizes, widest first.
	Name             string            `json:"name"`                   //The name of the album.
	Popularity       int               `json:"popularity"`             //The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
	Released         string            `json:"release_date"`           //The date the album was first released, for example "1981-12-15". Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleasePrecision string            `json:"release_date_precision"` //The precision with which release_date value is known: "year", "month", or "day".
	Tracks           TrackPagingObject `json:"tracks"`
	Type             string            `json:"type"`
	Uri              string            `json:"uri"`
}
type AlbumSimple struct {
	AlbumType  string            `json:"album_type"`        //The type of the album: one of "album", "single", or "compilation".
	Markets    []string          `json:"available_markets"` //The markets in which the album is available: ISO 3166-1 alpha-2 country codes. Note: Album is available when at least 1 of its tracks is available in that market.
	ExternalUr map[string]string `json:"external_urls"`     //Known external URLs for this album.
	Link       string            `json:"href"`              //A link to the Web API endpoint providing full details of the album.
	ID         string            `json:"id"`                //The Spotify ID for the album.
	Images     []Image           `json:"images"`            //The cover art for the album in various sizes, widest first.
	Name       string            `json:"name"`              //The name of the album.
	Type       string            `json:"type"`
	Uri        string            `json:"uri"`
}

func (api SpotifyWeb) Album(albumId string) (AlbumContainer, error) {
	var album AlbumContainer
	err := api.DoBasic("albums/"+albumId, "GET", &album)
	return album, err
}

func (api SpotifyWeb) Albums(albumIds ...string) (MultiAlbumContainer, error) {
	ids := strings.Join(albumIds, ",")
	var albums MultiAlbumContainer
	err := api.DoBasic("albums?ids="+ids, "GET", &albums)
	return albums, err
}

func (api SpotifyWeb) AlbumTracks(albumId string) (TrackPagingObject, error) {
	var album AlbumContainer
	err := api.DoBasic("albums/"+albumId, "GET", &album)
	return album.Tracks, err
}
