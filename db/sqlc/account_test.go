package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sarangi/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Accounts {
	accountParam := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), accountParam)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, accountParam.Owner, account.Owner)
	require.Equal(t, accountParam.Balance, account.Balance)
	require.Equal(t, accountParam.Currency, account.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, account2)
	assert.Equal(t, account.ID, account2.ID)
	assert.Equal(t, account.Owner, account2.Owner)
	assert.Equal(t, account.Balance, account2.Balance)
	assert.Equal(t, account.Currency, account2.Currency)
	assert.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	params := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomBalance(),
	}

	account3, err := testQueries.UpdateAccount(context.Background(), params)
	assert.NoError(t, err)
	assert.NotEmpty(t, account3)
	assert.Equal(t, account3.ID, params.ID)
	assert.Equal(t, account3.Balance, params.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, account2)

	testQueries.DeleteAccount(context.Background(), account.ID)
	account3, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, account3)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	params := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), params)
	assert.NoError(t, err)
	assert.Len(t, accounts, 5)

	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}

}
