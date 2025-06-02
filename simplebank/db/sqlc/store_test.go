package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranferTx(t *testing.T) {
	store := NewStore(testPool)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	fmt.Println(">>before:", account1.Balance, account2.Balance)

	amount := int64(10)
	n := 5

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: int64(account1.ID),
				ToAccountID:   int64(account2.ID),
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	//check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err, "Error in transfer transaction")

		result := <-results
		require.NotEmpty(t, result, "Transfer transaction result should not be empty")

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer, "Transfer should not be empty")
		require.Equal(t, account1.ID, transfer.FromAccountID, "FromAccountID should match")
		require.Equal(t, account2.ID, transfer.ToAccountID, "ToAccountID should match")
		require.Equal(t, amount, transfer.Amount, "Transfer amount should match")
		require.NotZero(t, transfer.ID, "Transfer ID should not be zero")
		require.NotZero(t, transfer.CreatedAt, "Transfer created_at should not be zero")

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err, "Error getting transfer after transaction")

		//check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry, "From entry should not be empty")
		require.Equal(t, account1.ID, fromEntry.AccountID, "From entry AccountID should match")
		require.Equal(t, -amount, fromEntry.Amount, "From entry amount should be negative")
		require.NotZero(t, fromEntry.ID, "From entry ID should not be zero")
		require.NotZero(t, fromEntry.CreatedAt, "From entry created_at should not be zero")
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err, "Error getting from entry after transaction")

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry, "To entry should not be empty")
		require.Equal(t, account2.ID, toEntry.AccountID, "To entry AccountID should match")
		require.Equal(t, amount, toEntry.Amount, "To entry amount should match")
		require.NotZero(t, toEntry.ID, "To entry ID should not be zero")
		require.NotZero(t, toEntry.CreatedAt, "To entry created_at should not be zero")
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err, "Error getting to entry after transaction")

		//check accounts
		fromAccount := result.AccountFrom
		require.NotEmpty(t, fromAccount, "From account should not be empty")
		require.Equal(t, account1.ID, fromAccount.ID, "From account ID should match")
		toAccount := result.AccountTo
		require.NotEmpty(t, toAccount, "To account should not be empty")
		require.Equal(t, account2.ID, toAccount.ID, "To account ID should match")
		//check balances
		fmt.Println(">>tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0, "Balance difference should be positive")
		require.True(t, diff1%amount == 0, "Balance difference should be a multiple of the transfer amount")
	}
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	expectedFromBalance := account1.Balance - int64(n)*amount
	expectedToBalance := account2.Balance + int64(n)*amount

	fmt.Println(">>after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, expectedFromBalance, updatedAccount1.Balance)
	require.Equal(t, expectedToBalance, updatedAccount2.Balance)
}
