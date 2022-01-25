package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestQuarryStake(t *testing.T) {
	infra := GetTestInfra(t)

	// Saros PRT-CASH LP
	res, err := infra.Client.QuarryStake(context.Background(),
		QuarryStakeAccounts{
			Owner:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Rewarder:    solana.MustPublicKeyFromBase58("5LAZ5rUe1CLJoKYauyVXdbG6e7nBmY2j5mJ8PnesCA8z"), // PRT
			LpTokenMint: solana.MustPublicKeyFromBase58("EoNJJWQMdvMscCL7V6wNrGaDLi791sPZH1hSzHcwfsDj"), // PAI
		},
		1_000_000, // 1 LP
	)

	assert.NoError(t, err)
	fmt.Println(res)
}

func TestQuarryUnstake(t *testing.T) {
	infra := GetTestInfra(t)

	// Saros PRT-CASH LP
	res, err := infra.Client.QuarryUnstake(context.Background(),
		QuarryStakeAccounts{
			Owner:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Rewarder:    solana.MustPublicKeyFromBase58("5LAZ5rUe1CLJoKYauyVXdbG6e7nBmY2j5mJ8PnesCA8z"), // PRT
			LpTokenMint: solana.MustPublicKeyFromBase58("EoNJJWQMdvMscCL7V6wNrGaDLi791sPZH1hSzHcwfsDj"), // PAI
		},
		1_000_000, // 1 LP
	)

	assert.NoError(t, err)
	fmt.Println(res)
}
