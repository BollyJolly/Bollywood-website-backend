package song

// Service defines business operations for songs.
type Service interface {
	ListByCategory(categoryID string) ([]Song, error)
	GetByID(id int) (*Song, error)
}

// SongService contains song business logic.
type SongService struct {
	repo Repository
}

// NewService wires a Repository into the song service.
func NewService(repo Repository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) ListByCategory(categoryID string) ([]Song, error) {
	// TODO: implement list songs for a category
	return nil, nil
}

func (s *SongService) GetByID(id int) (*Song, error) {
	// TODO: implement get song by ID
	return nil, nil
}
