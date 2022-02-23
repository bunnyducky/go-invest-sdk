package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestSarosAddLiquidity(t *testing.T) {
	infra := GetTestInfra(t)
	// LP C98-CASH
	res, err := infra.Client.SarosAddLiquidity(context.Background(),
		SarosAddLiquidityAccount{
			Owner:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			LPTokenMint: solana.MustPublicKeyFromBase58("4g7DwZVn8edcSYR34zE4Dm13YLH6ZvtXCNnaUyQSFy9w"), // PAI
		},
		SarosAddLiquidityData{
			TokenAAmount: 0,
			TokenBAmount: 100_000, //0.1 CASH
		},
	)
	assert.NoError(t, err)
	fmt.Println(res)
}

func TestSarosAddLiquiditySimulate(t *testing.T) {
	infra := GetTestInfra(t)

	res, err := infra.Client.SarosAddLiquiditySimulate(context.Background(),
		SarosAddLiquidityAccount{
			Owner:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:       solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			LPTokenMint: solana.MustPublicKeyFromBase58("4g7DwZVn8edcSYR34zE4Dm13YLH6ZvtXCNnaUyQSFy9w"), // PAI
		},
		SarosAddLiquidityData{
			TokenAAmount: 1_000_000, // 1 PRT
			TokenBAmount: 0,
		},
	)

	assert.NoError(t, err)
	// LP token mint
	assert.Greater(t, res.Balances["4g7DwZVn8edcSYR34zE4Dm13YLH6ZvtXCNnaUyQSFy9w"], 0)
}
