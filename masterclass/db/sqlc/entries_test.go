package db

import (
	"SimpleBank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Entry, Account) {
	acc_args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, _ := testQueries.CreateAccount(context.Background(), acc_args)

	args := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	assert.Equal(t, args.AccountID, entry.AccountID)
	assert.Equal(t, args.Amount, entry.Amount)
	require.WithinDuration(t, acc.CreatedAt.Time, entry.CreatedAt.Time, time.Second)

	return entry, acc
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	curr, _ := createRandomEntry(t)

	acc, err := testQueries.GetEntry(context.Background(), curr.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	assert.Equal(t, curr.ID, acc.ID)
	assert.Equal(t, curr.AccountID, acc.AccountID)
	assert.Equal(t, curr.Amount, acc.Amount)
	require.WithinDuration(t, curr.CreatedAt.Time, acc.CreatedAt.Time, time.Second)
}

func TestListEntry(t *testing.T) {
	_, acc := createRandomEntry(t)

	args := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     2,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
