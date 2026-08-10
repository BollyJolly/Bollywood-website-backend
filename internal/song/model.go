package song

// Song is a single track belonging to a category.
// Each category is expected to hold 75 songs for bingo card generation.

type Song struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}
