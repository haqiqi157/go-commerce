package service

import (
	"errors"
	"go-echo/internal/entity"
	"go-echo/internal/repository"
	"go-echo/pkg/encrypt"
	"go-echo/pkg/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	encryptTool    encrypt.EncryptTool
}

func NewUserService(userRepository repository.UserRepository, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool) *userService {
	return &userService{
		userRepository: userRepository,
		tokenUseCase:   tokenUseCase,
		encryptTool:    encryptTool,
	}
}

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	user.Address, _ = s.encryptTool.Decrypt(user.Address)
	user.Phone, _ = s.encryptTool.Decrypt(user.Phone)

	claims := token.JwtCustomClaims{
		ID:      user.ID.String(),
		Email:   user.Email,
		Role:    user.Role,
		Address: user.Address,
		Phone:   user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Go-Commerce",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("ada kesalahan dari sistem")
	}

	return token, nil
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	users, err := s.userRepository.FindAllUser()
	if err != nil {
		return nil, err
	}

	formattedUser := make([]entity.User, 0)
	for _, v := range users {
		v.Address, _ = s.encryptTool.Decrypt(v.Address)
		v.Phone, _ = s.encryptTool.Decrypt(v.Phone)
		formattedUser = append(formattedUser, v)
	}

	return formattedUser, nil
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	user.Address, _ = s.encryptTool.Encrypt(user.Address)
	user.Phone, _ = s.encryptTool.Encrypt(user.Phone)

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	newUser.Address, _ = s.encryptTool.Decrypt(newUser.Address)
	newUser.Phone, _ = s.encryptTool.Decrypt(newUser.Phone)

	return newUser, nil
}

func (s *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if user.Address != "" {
		user.Address, _ = s.encryptTool.Encrypt(user.Address)
	}
	if user.Phone != "" {
		user.Phone, _ = s.encryptTool.Encrypt(user.Phone)
	}
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
