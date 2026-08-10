package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/amanhasnainy/bingo-backend/internal/session"
	"github.com/amanhasnainy/bingo-backend/internal/user"
)

type Service struct {
	userRepository    *user.Repository
	sessionRepository *session.Repository
}
func NewService(
	userRepository *user.Repository,
	sessionRepository *session.Repository,
) *Service {

	return &Service{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *Service) Register(
	req RegisterRequest,
) error {

	existingUser, _ := s.userRepository.FindByEmail(
		req.Email,
	)

	if existingUser != nil {
		return errors.New(
			"user already exists",
		)
	}

	hashedPassword, err := HashPassword(
		req.Password,
	)

	fmt.Println("Input password:", req.Password)
    fmt.Println("Hashed password:", hashedPassword)

	if err != nil {
		return err
	}

	newUser := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Stars:    30,
	}

	err = s.userRepository.Create(
		newUser,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(
	req LoginRequest,
) (*AuthResponse, error) {

	existingUser, err := s.userRepository.FindByEmail(
		req.Email,
	)

	if err != nil {
		return nil, errors.New(
			"invalid email or password",
		)
	}

	err = ComparePassword(
		req.Password,
		existingUser.Password,
	)

	if err != nil {
		return nil, errors.New(
			"invalid email or password",
		)
	}

	refreshToken, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	newSession := session.Session{
		UserID:       existingUser.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(30 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	savedSession, err := s.sessionRepository.Create(
		newSession,
	)

	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateAccessToken(
		existingUser.ID.Hex(),
		savedSession.ID.Hex(),
	)

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Refresh(
	req RefreshRequest,
) (*AuthResponse, error) {

	session, err := s.sessionRepository.FindByRefreshToken(
		req.RefreshToken,
	)

	if err != nil {
		return nil, errors.New(
			"invalid refresh token",
		)
	}

	if time.Now().After(
		session.ExpiresAt,
	) {
		return nil, errors.New(
			"refresh token expired",
		)
	}

	accessToken, err := GenerateAccessToken(
		session.UserID.Hex(),
		session.ID.Hex(),
	)

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken: accessToken,
	}, nil
}