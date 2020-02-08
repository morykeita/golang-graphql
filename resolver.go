//go:generate go run github.com/99designs/gqlgen -v

package graphql_golang

import (
	"context"

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

var user = []*models.User{
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
	panic("implement me")
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
	return nil, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}
