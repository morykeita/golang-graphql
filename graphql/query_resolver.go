/**
 * @author Mory Keita on 2/9/20
 */
package graphql

import (
	"context"
	"github.com/morykeita/graphql-golang/models"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }


func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepository.GetMeetups()
}
