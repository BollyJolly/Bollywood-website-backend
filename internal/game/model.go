package game

// Cell is one square on a bingo card.
// IsFree marks the center free space when using a classic 5x5 layout.
type Cell struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Marked bool   `json:"marked"`
	IsFree bool   `json:"isFree"`
}

// Card is a player's unique bingo board for a game.
type Card struct {
	PlayerID string `json:"playerId"`
	Cells    []Cell `json:"cells"`
}

// Game tracks live match state for a room.
// Status values will typically be: "active", "finished".
type Game struct {
	ID             string `json:"id"`
	RoomID         string `json:"roomId"`
	CurrentSongID  int    `json:"currentSongId"`
	CalledSongIDs  []int  `json:"calledSongIds"`
	WinnerPlayerID string `json:"winnerPlayerId"`
	Status         string `json:"status"`
}
