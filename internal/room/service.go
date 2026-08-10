package room

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/amanhasnainy/bingo-backend/internal/config"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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

func (s *Service) Create(
	hostID string,
	playlistID string,
) (*Room, error) {

	hostObjectID, err := bson.ObjectIDFromHex(
		hostID,
	)

	if err != nil {
		return nil, err
	}

	room := Room{
		Code: generateRoomCode(),

		PlaylistID: playlistID,

		HostID: hostObjectID,

		PlayerIDs: []bson.ObjectID{
			hostObjectID,
		},

		CalledNumbers: []int{},

		Status: StatusWaiting,

		CreatedAt: time.Now(),
	}

	err = s.repository.Create(
		room,
	)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func generateRoomCode() string {

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, 6)

	for i := range code {
		code[i] = characters[rand.Intn(len(characters))]
	}

	return string(code)
}

func (s *Service) Join(
	roomCode string,
	userID string,
) error {

	room, err := s.repository.FindByCode(
		roomCode,
	)

	if err != nil {
		return err
	}

	userObjectID, err := bson.ObjectIDFromHex(
		userID,
	)

	if err != nil {
		return err
	}

	room.PlayerIDs = append(
		room.PlayerIDs,
		userObjectID,
	)

	return s.repository.Update(
		*room,
	)
}

func (s *Service) GetAll() ([]Room, error) {

	return s.repository.FindAll()
}

func (s *Service) GetByCode(
	roomCode string,
) (*Room, error) {

	return s.repository.FindByCode(
		roomCode,
	)
}

func (s *Service) Start(
	roomCode string,
	userID string,
) error {

	room, err := s.repository.FindByCode(
		roomCode,
	)

	if err != nil {
		return err
	}

	if room.HostID.Hex() != userID {
		return errors.New(
			"only the host can start the game",
		)
	}

	room.Status = StatusPlaying

	return s.repository.Update(
		*room,
	)
}

func (s *Service) Leave(
	roomCode string,
	userID string,
) error {

	room, err := s.repository.FindByCode(
		roomCode,
	)

	if err != nil {
		return err
	}

	userObjectID, err := bson.ObjectIDFromHex(
		userID,
	)

	if err != nil {
		return err
	}

	var updatedPlayers []bson.ObjectID

	for _, playerID := range room.PlayerIDs {

		if playerID != userObjectID {
			updatedPlayers = append(
				updatedPlayers,
				playerID,
			)
		}
	}

	room.PlayerIDs = updatedPlayers

	return s.repository.Update(
		*room,
	)
}

func (s *Service) Next(
	roomCode string,
	userID string,
) (int, string, error) {

	room, err := s.repository.FindByCode(
		roomCode,
	)

	if err != nil {
		return 0, "", err
	}

	if room.HostID.Hex() != userID {
		return 0, "", errors.New(
			"only the host can call the next song",
		)
	}

	if room.Status != StatusPlaying {
		return 0, "", errors.New(
			"game has not started",
		)
	}

	if len(room.CalledNumbers) == 75 {
		return 0, "", errors.New(
			"all songs have been played",
		)
	}

	var number int

	for {

		number = rand.Intn(75) + 1

		found := false

		for _, called := range room.CalledNumbers {

			if called == number {
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	room.CalledNumbers = append(
		room.CalledNumbers,
		number,
	)

	err = s.repository.Update(
		*room,
	)

	if err != nil {
		return 0, "", err
	}

	url := fmt.Sprintf(
		"%s/%d.m4a",
		config.GetEnv("R2_PUBLIC_URL"),
		number,
	)

	return number, url, nil
}