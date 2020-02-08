//go:generate go run github.com/99designs/gqlgen -v

package graphql_golang

import (
	"context"
	"errors"

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

var users = []*models.User{
	{
		ID:       "1",
		Username: "magicmory",
		Email:    "magicmory@gmail.com",
	},
	{
		ID:       "2",
		Username: "johndoe",
		Email:    "johndoe@gmail.com",
	},
}

type Resolver struct {
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

func (m meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {

	user := new(models.User)
	for _, u := range users {
		if u.ID == obj.ID {
			user = u
			break
		}
	}
	if user == nil {
		return nil, errors.New("User with id does not exist")
	}
	return user, nil

}

type userResolver struct {
	*Resolver
}

func (m mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("implement me")
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
	return meetups, nil
}
