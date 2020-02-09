/**
 * @author Mory Keita on 2/9/20
 */
package graphql

import (
	"context"
	"errors"
	"github.com/morykeita/graphql-golang/models"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct {
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

