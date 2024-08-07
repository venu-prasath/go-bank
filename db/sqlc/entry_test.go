package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/venu-prasath/go-bank/util"
)

func createRandomEntry(t *testing.T, id int64) Entry {
	arg := CreateEntryParams {
		AccountID: id,
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func createXEntries() {
	var i int64
	for ; i<10; i++ {
		arg := CreateEntryParams {
			AccountID: 5,
			Amount: util.RandomMoney(),
		}
		_, err := testQueries.CreateEntry(context.Background(), arg)
		if err != nil {
			log.Println("Error occured: ", err)
		}

	}
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t, 51)
}

func TestCreateXEntries(t *testing.T) {
	createXEntries()
}

func TestGetEntry(t *testing.T) {
	//Create Entry
	entry1 := createRandomEntry(t, 52)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.NotZero(t, entry2.ID)
	require.NotZero(t, entry2.CreatedAt)
}

func TestListEntry(t *testing.T) {
	// for i:=0; i<10; i++ {
	// 	createRandomEntry(t, 53)
	// }

	arg := ListEntriesParams {
		Limit: 5,
		Offset: 5,
		AccountID: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}


}