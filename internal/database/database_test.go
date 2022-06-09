package database

import (
	"context"
	"testing"

	"github.com/jmsilvadev/atomics-commits-aio/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestUnit(t *testing.T) {
	ctx := context.Background()
	db := NewPostgres(ctx)
	itx := &entities.Transaction{
		Block: "1",
		Hash:  1,
	}
	_, err := db.Upsert(ctx, itx)
	require.NoError(t, err)
}
