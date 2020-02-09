/**
 * @author Mory Keita on 2/8/20
 */
package database

import (
	"github.com/go-pg/pg/v9"
	"github.com/morykeita/graphql-golang/models"
)

type UserRepository struct {
	DB *pg.DB
}

func (u *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

