package spotifyweb

import (
	"strings"
)

type ArtistSimple struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href,
	Id,
	Name,
	Type,
	Uri string
}
type MultiArtistContainer struct {
	Artists []Artist
}
type Artist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href,
	Id,
	Name,
	Type,
	Uri string
	Images     []Image
	Genres     []string
	Popularity int
}

func (api SpotifyWeb) Artist(artistId string) (Artist, error) {
	var artist Artist
	err := api.DoBasic("artists/"+artistId, "GET", &artist)
	return artist, err
}

func (api SpotifyWeb) Artists(artistIds ...string) (MultiArtistContainer, error) {
	ids := strings.Join(artistIds, ",")
	var artists MultiArtistContainer
	err := api.DoBasic("artists?ids="+ids, "GET", &artists)
	return artists, err
}
