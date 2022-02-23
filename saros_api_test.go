package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestSarosStake(t *testing.T) {
	infra := GetTestInfra(t)
	// LP C98-CASH
	res, err := infra.Client.SarosStake(context.Background(),
		SarosStakeAccounts{
			Payer:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Owner:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			LPTokenMint: solana.MustPublicKeyFromBase58("4g7DwZVn8edcSYR34zE4Dm13YLH6ZvtXCNnaUyQSFy9w"),
		},
		SarosStakeData{
			StakeAmountA: 0,
			StakeAmountB: 100_000, //0.1 CASH
		},
	)
	assert.NoError(t, err)
	fmt.Println(res)
}

func TestSarosUnstakeSimulate(t *testing.T) {
	infra := GetTestInfra(t)

	res, err := infra.Client.SarosUnstake(context.Background(),
		SarosUnstakeAccounts{
			Payer:        solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Owner:        solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			LPTokenMint:  solana.MustPublicKeyFromBase58("4g7DwZVn8edcSYR34zE4Dm13YLH6ZvtXCNnaUyQSFy9w"),
			WithdrawMint: solana.MustPublicKeyFromBase58("C98A4nkJXhpVZNAZdHUA95RpTF3T4whtQubL3YobiUX9"), // C98
		},
		100_000, // 1 PRT
	)
	assert.NoError(t, err)
	fmt.Println(res)
}
