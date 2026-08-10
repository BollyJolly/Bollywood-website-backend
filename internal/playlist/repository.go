package playlist

type Repository struct {
}

func (r *Repository) FindAll() ([]Playlist, error) {

	return []Playlist{
		{
			ID:          "bollywood-classics",
			Name:        "Bollywood Classics",
			Description: "75 Evergreen Bollywood Songs",
			Thumbnail:   "https://pub-49cc62f340ac4f6a85f3f323289a437b.r2.dev/playlist.jpg",
			EntryFee:    5,
			TotalSongs:  75,
		},
	}, nil
}