package playlist

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll() ([]Playlist, error) {

	return s.repository.FindAll()
}