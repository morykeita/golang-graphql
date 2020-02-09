//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"github.com/morykeita/graphql-golang/database"
	"github.com/morykeita/graphql-golang/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "first meetup",
		Description: "first meetup description",
		UserID:      "2",
	},
	{
		ID:          "2",
		Name:        "second meetup",
		Description: "second meetup description",
		UserID:      "1",
	},
}

type Resolver struct {
	database.MeetupsRepository
	database.UserRepository
}

