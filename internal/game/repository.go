package game

// Repository defines data-access methods for games and bingo cards.
type Repository interface {
	Create(g *Game) error
	FindByRoomID(roomID string) (*Game, error)
	Update(g *Game) error
	SaveCard(roomID string, card *Card) error
	FindCard(roomID, playerID string) (*Card, error)
}

// MemoryRepository stores games and cards in maps.
// Cards are keyed as "roomID:playerID" for quick lookup.
type MemoryRepository struct {
	games map[string]*Game // keyed by room ID
	cards map[string]*Card // keyed by "roomID:playerID"
}

// NewMemoryRepository creates an empty in-memory game store.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		games: make(map[string]*Game),
		cards: make(map[string]*Card),
	}
}

func (r *MemoryRepository) Create(g *Game) error {
	// TODO: implement create game
	return nil
}

func (r *MemoryRepository) FindByRoomID(roomID string) (*Game, error) {
	// TODO: implement find game by room ID
	return nil, nil
}

func (r *MemoryRepository) Update(g *Game) error {
	// TODO: implement update game
	return nil
}

func (r *MemoryRepository) SaveCard(roomID string, card *Card) error {
	// TODO: implement save card
	return nil
}

func (r *MemoryRepository) FindCard(roomID, playerID string) (*Card, error) {
	// TODO: implement find card
	return nil, nil
}
