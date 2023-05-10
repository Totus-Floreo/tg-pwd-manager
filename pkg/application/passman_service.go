package application

import (
	"errors"
	"strings"

	"github.com/Totus-Floreo/tg-pwd-manager/pkg/model"
	"github.com/Totus-Floreo/tg-pwd-manager/pkg/repository"
)

type PassManService interface {
	CreateUser(int64) error
	CreateCredentials(*model.User, string) error
	GetUserByID(int64) (*model.User, error)
	GetUserCredentials(*model.User, string) (*model.Credentials, error)
	DeleteCredentials(*model.User, string) error
	GetUserExists(int64) (bool, error)
}

type passManService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) PassManService {
	return &passManService{userRepo}
}

func (s *passManService) CreateUser(userID int64) error {
	user := &model.User{
		UserID:  userID,
		Storage: make(map[string]*model.Credentials),
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *passManService) CreateCredentials(user *model.User, message string) error {
	strings := strings.Split(message, " ")
	if len(strings) != 4 {
		return errors.New("Not like a /set service login password")
	}
	cred := &model.Credentials{Login: strings[2], Password: strings[3]}
	user.Storage[strings[1]] = cred
	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *passManService) GetUserByID(userID int64) (*model.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (s *passManService) GetUserCredentials(user *model.User, message string) (*model.Credentials, error) {
	strings := strings.Split(message, " ")
	cred, ok := user.Storage[strings[1]]
	if !ok {
		return &model.Credentials{}, errors.New("Not Found Credentials by name service")
	}
	return cred, nil
}

func (s *passManService) DeleteCredentials(user *model.User, key string) error {
	delete(user.Storage, key)
	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *passManService) GetUserExists(userID int64) (bool, error) {
	exists, err := s.userRepo.UserExists(userID)
	if err != nil {
		return true, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
