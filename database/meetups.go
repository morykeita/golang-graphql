/**
 * @author Mory Keita on 2/8/20
 */
package database

import (
	"github.com/go-pg/pg/v9"
	"github.com/morykeita/graphql-golang/models"
)

type MeetupsRepository struct {
	DB *pg.DB
}

func (m *MeetupsRepository) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Select()

	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupsRepository) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_,err := m.DB.Model(meetup).Returning("*").Insert()
	return meetup,err
}

