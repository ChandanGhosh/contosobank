package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/chandanghosh/contosobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)
	require.Equal(t, arg.Owner, acc.Owner)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	acc, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc1.Balance, acc.Balance)
	require.Equal(t, acc1.ID, acc.ID)
	require.Equal(t, acc1.Owner, acc.Owner)
	require.Equal(t, acc1.Currency, acc.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: util.RandomMoney(),
		ID:      acc1.ID,
	}

	acc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.ID, acc.ID)
	require.Equal(t, acc1.Owner, acc.Owner)
	require.Equal(t, acc1.Currency, acc.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5, Offset: 5})
	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
