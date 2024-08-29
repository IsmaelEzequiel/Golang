package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	amount := int64(10)

	transfer, err := storeDB.CreateTransfer(context.Background(), CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	assert.Equal(t, acc1.ID, transfer.FromAccountID)
	assert.Equal(t, acc2.ID, transfer.ToAccountID)
	assert.Equal(t, amount, transfer.Amount)
	assert.NotZero(t, transfer.ID)
}
