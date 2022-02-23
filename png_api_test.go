package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestPNGAddLiquidity(t *testing.T) {
	infra := GetTestInfra(t)

	res, err := infra.Client.PNGAddLiquidity(context.Background(),
		PNGAddLiquidityAccount{
			Owner:      solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:      solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			TokenAMint: solana.MustPublicKeyFromBase58("PRT88RkA4Kg5z7pKnezeNH4mafTvtQdfFgpQTGRjz44"),  // PRT
			TokenBMint: solana.MustPublicKeyFromBase58("Ea5SjE2Y6yvCeW5dYTn7PYMuW5ikXkvbGdcmSnXeaLjS"), // PAI
		},
		PNGAddLiquidityData{
			TokenAAmount: 1_000_000, // 1 PRT
			TokenBAmount: 0,
		},
	)

	assert.NoError(t, err)
	fmt.Println(res)
}

func TestPNGAddLiquiditySimulate(t *testing.T) {
	infra := GetTestInfra(t)

	res, err := infra.Client.PNGAddLiquiditySimulate(context.Background(),
		PNGAddLiquidityAccount{
			Owner:      solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			Payer:      solana.MustPublicKeyFromBase58("6uHZK4v5fz9JZX6VaVyxus5AxFqJk7B8jiHw1gMYaSZg"),
			TokenAMint: solana.MustPublicKeyFromBase58("PRT88RkA4Kg5z7pKnezeNH4mafTvtQdfFgpQTGRjz44"),  // PRT
			TokenBMint: solana.MustPublicKeyFromBase58("Ea5SjE2Y6yvCeW5dYTn7PYMuW5ikXkvbGdcmSnXeaLjS"), // PAI
		},
		PNGAddLiquidityData{
			TokenAAmount: 1_000_000, // 1 PRT
			TokenBAmount: 0,
		},
	)

	assert.NoError(t, err)
	// LP token mint
	assert.Greater(t, res.Balances["BPY5oGvks5bR6zQyyWopQmijj7eZYUM1r7YJPmTJ5FKx"], 0)
}
