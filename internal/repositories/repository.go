package repositories

import (
	"fmt"

	"github.com/Naumovets/go-auth/internal/entities"
	"github.com/Naumovets/go-auth/internal/utils"
	"github.com/go-pg/pg"
)

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) ExistsUser(username string) (bool, error) {
	user := &entities.User{}
	res, err := r.db.Model(user).Where("username = ?", username).Exists()
	if err != nil {
		return false, fmt.Errorf("get user: %s", err)
	}

	return res, nil
}

func (r *Repository) AddUser(user *entities.User) error {

	hashPass, err := utils.HashPassword(user.Password)

	if err != nil {
		return fmt.Errorf("hash pass: %s", err)
	}

	user.Password = hashPass

	_, err = r.db.Model(user).Insert()

	if err != nil {
		return fmt.Errorf("add user: %s", err)
	}

	return nil

}

func (r *Repository) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		return nil, fmt.Errorf("get user: %s", err)
	}

	return user, nil
}

func (r *Repository) GetUsersByIds(ids []string) ([]*entities.User, error) {
	users := make([]*entities.User, 0)
	err := r.db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

	if err != nil {
		return nil, fmt.Errorf("get users by ids: %s", err)
	}

	return users, nil

}
