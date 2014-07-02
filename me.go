package spotifyweb

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
}

func (api SpotifyWeb) Profile() (Me, error) {
	var me Me
	err := api.DoAuth("me", "GET", &me)
	return me, err
}
