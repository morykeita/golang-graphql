/**
 * @author Mory Keita on 2/8/20
 */
package database

import (
	"context"
	"fmt"
)

type DBlogger struct{}

func (d DBlogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBlogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery)
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
