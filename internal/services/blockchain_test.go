package services

import (
	"context"
	"testing"

	"github.com/jmsilvadev/atomics-commits-aio/internal/database"
	"github.com/jmsilvadev/atomics-commits-aio/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestUnit(t *testing.T) {
	ctx := context.Background()
	db := database.NewPostgres(ctx)
	b := New(ctx, db)
	itx := &entities.Transaction{
		Block: "1",
		Hash:  1,
	}
	r, err := b.CreateTransaction(ctx, itx)
	require.NoError(t, err)
	require.Equal(t, itx.Block, r.Block)
	require.Equal(t, itx.Hash, r.Hash)
}
