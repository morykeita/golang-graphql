//go:generate go run github.com/99designs/gqlgen -v

package graphql_golang

import (
	"context"
	"errors"
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

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct {
	*Resolver
}

type meetupResolver struct {
	*Resolver
}

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return m.UserRepository.GetUserById(obj.UserID)
}

type userResolver struct {
	*Resolver
}

func (m mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3{
		return nil, errors.New("name should be at least 3 characters long")
	}
	if len(input.Description) < 3{
		return nil, errors.New("description should be at least 3 characters long")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}
	return m.MeetupsRepository.CreateMeetup(meetup)
}
func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (u userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup
	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}
	return m, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepository.GetMeetups()
}
