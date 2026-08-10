package category

// Category groups songs into themed playlists (e.g. Wedding Songs).
type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PlaylistURL string `json:"playlistUrl"`
}
