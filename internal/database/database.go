package database

import "context"

type DatabaseProvider interface {
	Upsert(ctx context.Context, data interface{}) (interface{}, error)
}
