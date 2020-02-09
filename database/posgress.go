/**
 * @author Mory Keita on 2/8/20
 */
package database

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
)

type DBlogger struct{}

func (d DBlogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBlogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {

	var db *pg.DB =  pg.Connect(opts)
	if db == nil{
		fmt.Print("failed to connected to dabase", )
	}
	return db
}
