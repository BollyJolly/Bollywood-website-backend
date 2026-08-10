package game

import "github.com/amanhasnainy/bingo-backend/internal/song"

// GenerateCard builds a bingo card for a player from a pool of songs.
// Logic will live here later so the game engine stays focused on rules.
func GenerateCard(playerID string, songs []song.Song) (*Card, error) {
	// TODO: pick random songs, build cells (including free center if needed)
	return nil, nil
}
