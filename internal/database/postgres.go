package database

import (
	"context"
)

type postgresDB struct {
	ctx context.Context
}

var _ DatabaseProvider = &postgresDB{}

func NewPostgres(ctx context.Context) *postgresDB {
	return &postgresDB{
		ctx: ctx,
	}
}

func (db *postgresDB) Upsert(ctx context.Context, data interface{}) (interface{}, error) {
	return data, nil
}
