package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type PNGAddLiquidityAccount struct {
	Owner      solana.PublicKey
	Payer      solana.PublicKey
	TokenAMint solana.PublicKey
	TokenBMint solana.PublicKey
}

type PNGAddLiquidityData struct {
	TokenAAmount uint64
	TokenBAmount uint64
}

func (c *Client) PNGAddLiquidityData(ctx context.Context, accounts PNGAddLiquidityAccount, data PNGAddLiquidityData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "png-fi/add-liquidity", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenAMint":   accounts.TokenAMint,
			"tokenBMint":   accounts.TokenBMint,
		},
		Data: map[string]interface{}{
			"tokenAAmount": data.TokenAAmount,
			"tokenBAmount": data.TokenBAmount,
		},
	}, &result)

	return &result, err
}
