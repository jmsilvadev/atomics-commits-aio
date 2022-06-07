package repositories

import (
	"context"

	"github.com/jmsilvadev/atomics-commits-aio/internal/database"
)

type blockchain struct {
	ctx context.Context
	db  database.DatabaseProvider
}

var _ Repository = &blockchain{}

func NewBlockchain(ctx context.Context, db database.DatabaseProvider) *blockchain {
	return &blockchain{
		ctx: ctx,
		db:  db,
	}
}

func (r *blockchain) Save(ctx context.Context, data interface{}) (interface{}, error) {
	result, err := r.db.Upsert(ctx, data)
	return result, err
}
