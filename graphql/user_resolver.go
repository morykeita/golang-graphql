/**
 * @author Mory Keita on 2/8/20
 */
package graphql

import (
	"context"
	"github.com/morykeita/graphql-golang/models"
)

type userResolver struct {
	*Resolver
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
