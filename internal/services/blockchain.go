package services

import (
	"context"

	"github.com/jmsilvadev/atomics-commits-aio/internal/database"
	"github.com/jmsilvadev/atomics-commits-aio/internal/entities"
	"github.com/jmsilvadev/atomics-commits-aio/internal/repositories"
)

type blockchain struct {
	ctx context.Context
	db  database.DatabaseProvider
}

var _ Service = &blockchain{}

func New(ctx context.Context, db database.DatabaseProvider) *blockchain {
	return &blockchain{
		ctx: ctx,
		db:  db,
	}
}

func (s *blockchain) CreateTransaction(ctx context.Context, data *entities.Transaction) (*entities.Transaction, error) {
	repo := repositories.NewBlockchain(ctx, s.db)
	result, err := repo.Save(ctx, data)
	return result.(*entities.Transaction), err
}
