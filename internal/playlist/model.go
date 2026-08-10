package playlist

type Playlist struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
	EntryFee    int    `json:"entryFee" bson:"entryFee"`
	TotalSongs  int    `json:"totalSongs" bson:"totalSongs"`
}