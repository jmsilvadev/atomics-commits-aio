package repositories

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
	b := NewBlockchain(ctx, db)
	itx := &entities.Transaction{
		Block: "1",
		Hash:  1,
	}
	_, err := b.Save(ctx, itx)
	require.NoError(t, err)
}
