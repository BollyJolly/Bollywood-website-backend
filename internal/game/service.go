package game

// Service defines business operations for bingo gameplay.
type Service interface {
	GetCard(roomID, playerID string) (*Card, error)
	Mark(roomID, playerID string, songID int) (*Card, error)
	ClaimBingo(roomID, playerID string) (*Game, error)
}

// GameService contains gameplay business logic.
// It will later depend on room/song repositories for starting games and calling songs.
type GameService struct {
	repo Repository
}

// NewService wires a Repository into the game service.
func NewService(repo Repository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) GetCard(roomID, playerID string) (*Card, error) {
	// TODO: load or generate the player's bingo card
	return nil, nil
}

func (s *GameService) Mark(roomID, playerID string, songID int) (*Card, error) {
	// TODO: mark a song on the player's card
	return nil, nil
}

func (s *GameService) ClaimBingo(roomID, playerID string) (*Game, error) {
	// TODO: validate bingo claim using engine.CheckWin
	return nil, nil
}
