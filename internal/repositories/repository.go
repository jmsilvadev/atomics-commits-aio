package repositories

import "context"

type Repository interface {
	Save(ctx context.Context, data interface{}) (interface{}, error)
}
