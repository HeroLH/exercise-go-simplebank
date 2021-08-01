package db

import (
	"context"
	"database/sql"
	"exercise-go-simplebank/db/sqlc/db"
	"exercise-go-simplebank/infrastructure/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func RandomCurrency() string {
	Currencies := []string{"EUR", "USD", "CAD", "RMB"}
	n := len(Currencies)
	return Currencies[utils.RandomInt(0, int64(n-1))]
}

func RandomMoney() int64 {
	return utils.RandomInt(0, 1000)
}

func RandomOwner() string {
	return utils.RandomString(6)
}

func createRandomAccount(t *testing.T) db.Account {
	args := db.CreateAccountParams{
		Owner: RandomOwner(),
		Balance: RandomMoney(),
		Currency: RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := db.UpdateAccountParams{
		ID:      account1.ID,
		Balance: RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	// 检查两个时间戳是否相同, 最多有点增量持续时间, 这边设置增量持续时间为 1 秒
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	var lastAccount db.Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	arg := db.ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}