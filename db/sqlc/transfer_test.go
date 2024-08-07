package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/venu-prasath/go-bank/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams {
		FromAccountID: 1,
		ToAccountID: 2,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	t1 := createRandomTransfer(t)
	t2, err := testQueries.GetTransfer(context.Background(), 1)
	
	require.NoError(t, err)
	require.NotEmpty(t, t2)

	//require.Equal(t, t1.Amount, t2.Amount)
	require.Equal(t, t1.FromAccountID, t2.FromAccountID)
	require.Equal(t, t1.ToAccountID, t2.ToAccountID)
}

func TestListTranfers(t *testing.T) {
	arg := ListTransfersParams {
		Limit: 5,
		Offset: 5,
		FromAccountID: 1,
		ToAccountID: 2,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}