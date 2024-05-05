package service

import (
	"errors"
	"go-echo/internal/entity"
	"go-echo/internal/repository"
	"go-echo/pkg/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserService interface {
	Login(email string, password string) (string, error)
	FindAllUser() ([]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id uuid.UUID) (bool, error)
	FindUserByID(id uuid.UUID) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	tokenUseCase   token.TokenUseCase
}

func NewUserService(userRepository repository.UserRepository, tokenUseCase token.TokenUseCase) *userService {
	return &userService{
		userRepository: userRepository,
		tokenUseCase:   tokenUseCase,
	}
}

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("username/password are wrong, please check again")
	}

	if user.Password != password {
		return "", errors.New("username/password are wrong, please check again")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	claims := token.JwtCustomClaims{
		ID:      user.ID.String(),
		Email:   user.Email,
		Role:    user.Role,
		Address: user.Address,
		Phone:   user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-commerce",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)

	if err != nil {
		return "", errors.New("internal error system")
	}

	return token, nil
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	return s.userRepository.FindAllUser()
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	return s.userRepository.UpdateUser(user)
}
func (s *userService) DeleteUser(id uuid.UUID) (bool, error) {
	user, err := s.userRepository.FindUserByID(id)

	if err != nil {
		return false, err
	}

	return s.userRepository.DeleteUser(user)
}

func (s *userService) FindUserByID(id uuid.UUID) (*entity.User, error) {
	return s.userRepository.FindUserByID(id)
}
