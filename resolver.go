//go:generate go run github.com/99designs/gqlgen -v

package graphql_golang

import (
	"context"

	"github.com/morykeita/graphql-golang/models"
)

var meetups = []models.Meetup{
	{
		ID:          "1",
		Name:        "first meetup",
		Description: "first meetup description",
	},
	{
		ID:          "2",
		Name:        "second meetup",
		Description: "second meetup description",
	},
}

type Resolver struct{}

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

func (m mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("implement me")
}

type userResolver struct {
	*Resolver
}

func (u userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	panic("implement me")
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	panic("not implemented")
}
