package investsdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestPNGApi(t *testing.T) {
	infra := GetTestInfra(t)

	res, err := infra.Client.PNGAddLiquidityData(context.Background(),
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
