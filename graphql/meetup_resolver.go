/**
 * @author Mory Keita on 2/8/20
 */
package graphql

import (
	"context"
	"github.com/morykeita/graphql-golang/models"
)


type meetupResolver struct {
	*Resolver
}

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}



