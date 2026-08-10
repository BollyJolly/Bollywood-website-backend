package song

// Repository defines data-access methods for songs.
type Repository interface {
	FindByCategoryID(categoryID string) ([]Song, error)
	FindByID(id int) (*Song, error)
	Create(s *Song) error
	CreateMany(songs []Song) error
}

// MemoryRepository stores songs in a slice for now.
type MemoryRepository struct {
	songs []Song
}

// NewMemoryRepository creates an empty in-memory song store.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		songs: make([]Song, 0),
	}
}

func (r *MemoryRepository) FindByCategoryID(categoryID string) ([]Song, error) {
	// TODO: implement find songs by category
	return nil, nil
}

func (r *MemoryRepository) FindByID(id int) (*Song, error) {
	// TODO: implement find song by ID
	return nil, nil
}

func (r *MemoryRepository) Create(s *Song) error {
	// TODO: implement create
	return nil
}

func (r *MemoryRepository) CreateMany(songs []Song) error {
	// TODO: implement bulk create (useful for seeding 75 songs per category)
	return nil
}
