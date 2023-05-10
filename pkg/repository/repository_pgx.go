package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/Totus-Floreo/tg-pwd-manager/pkg/model"
)

type UserRepositoryInterface interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByID(int64) (*model.User, error)
	UserExists(int64) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	data, err := json.Marshal(user.Storage)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users (user_id, storage) VALUES ($1, $2)", user.UserID, data)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	data, err := json.Marshal(user.Storage)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE users SET storage = $1 WHERE user_id = $2", data, user.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	var data []byte
	err := r.db.QueryRow("SELECT storage FROM users WHERE user_id = $1", id).Scan(&data)
	if err != nil {
		return nil, err
	}

	user := &model.User{UserID: id, Storage: make(map[string]*model.Credentials)}
	err = json.Unmarshal(data, &user.Storage)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UserExists(userID int64) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
