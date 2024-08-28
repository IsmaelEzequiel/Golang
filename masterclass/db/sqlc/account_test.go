package db

import (
	"SimpleBank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	assert.Equal(t, acc.Owner, args.Owner)
	assert.Equal(t, acc.Balance, args.Balance)
	assert.Equal(t, acc.Currency, args.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	curr := createRandomAccount(t)

	acc, err := testQueries.GetAccount(context.Background(), curr.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	assert.Equal(t, curr.ID, acc.ID)
	assert.Equal(t, curr.Owner, acc.Owner)
	assert.Equal(t, curr.Currency, acc.Currency)
	assert.Equal(t, curr.Balance, acc.Balance)
	require.WithinDuration(t, curr.CreatedAt.Time, acc.CreatedAt.Time, time.Second)
}

func TestUpdatedAccount(t *testing.T) {
	curr := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      curr.ID,
		Balance: util.RandomMoney(),
	}

	acc, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	assert.Equal(t, curr.ID, acc.ID)
	assert.Equal(t, args.Balance, acc.Balance)
	require.WithinDuration(t, curr.CreatedAt.Time, acc.CreatedAt.Time, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	curr := createRandomAccount(t)

	err := testQueries.DeleteAccouunt(context.Background(), curr.ID)
	require.NoError(t, err)

	acc, err := testQueries.GetAccount(context.Background(), curr.ID)
	require.Error(t, err)
	require.EqualError(t, err, err.Error())
	require.Empty(t, acc)
}

func TestListAccount(t *testing.T) {
	args := ListAccountsParams{
		Limit:  2,
		Offset: 0,
	}

	accs, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, accs)

	for _, account := range accs {
		require.NotEmpty(t, account)
	}
}
