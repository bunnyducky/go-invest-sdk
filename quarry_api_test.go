package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestQuarry_Stake(t *testing.T) {
	infra := GetTestInfra(t)

	// Saros PRT-CASH LP
	res, err := infra.Client.QuarryStake(context.Background(),
		QuarryAccounts{
			Owner:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Rewarder:  solana.MustPublicKeyFromBase58("5LAZ5rUe1CLJoKYauyVXdbG6e7nBmY2j5mJ8PnesCA8z"), // PRT
			TokenMint: solana.MustPublicKeyFromBase58("EoNJJWQMdvMscCL7V6wNrGaDLi791sPZH1hSzHcwfsDj"), // PAI
		},
		1_000_000, // 1 LP
	)

	assert.NoError(t, err)
	fmt.Println(res)
}

func TestQuarry_Unstake(t *testing.T) {
	infra := GetTestInfra(t)

	// Saros PRT-CASH LP
	res, err := infra.Client.QuarryUnstake(context.Background(),
		QuarryAccounts{
			Owner:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Rewarder:  solana.MustPublicKeyFromBase58("5LAZ5rUe1CLJoKYauyVXdbG6e7nBmY2j5mJ8PnesCA8z"), // PRT
			TokenMint: solana.MustPublicKeyFromBase58("EoNJJWQMdvMscCL7V6wNrGaDLi791sPZH1hSzHcwfsDj"), // PAI
		},
		1_000_000, // 1 LP
	)

	assert.NoError(t, err)
	fmt.Println(res)
}

func TestQuarry_Claim(t *testing.T) {
	infra := GetTestInfra(t)

	// Saros PRT-CASH LP
	res, err := infra.Client.QuarryClaim(context.Background(),
		QuarryAccounts{
			Owner:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:     solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Rewarder:  solana.MustPublicKeyFromBase58("5LAZ5rUe1CLJoKYauyVXdbG6e7nBmY2j5mJ8PnesCA8z"), // PRT
			TokenMint: solana.MustPublicKeyFromBase58("EoNJJWQMdvMscCL7V6wNrGaDLi791sPZH1hSzHcwfsDj"), // PAI
		},
	)

	assert.NoError(t, err)
	fmt.Println(res)
}
