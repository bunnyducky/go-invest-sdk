package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type SarosAddLiquidityAccount struct {
	Owner       solana.PublicKey
	Payer       solana.PublicKey
	LPTokenMint solana.PublicKey
}

type SarosAddLiquidityData struct {
	TokenAAmount uint64
	TokenBAmount uint64
}

func (c *Client) SarosAddLiquidity(ctx context.Context, accounts SarosAddLiquidityAccount, data SarosAddLiquidityData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saros/add-liquidity", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"lpTokenMint":  accounts.LPTokenMint,
		},
		Data: map[string]interface{}{
			"tokenAAmount": data.TokenAAmount,
			"tokenBAmount": data.TokenBAmount,
		},
	}, &result)
	return &result, err
}

func (c *Client) SarosAddLiquiditySimulate(ctx context.Context, accounts SarosAddLiquidityAccount, data SarosAddLiquidityData) (*SimulationResponse, error) {
	result := SimulationResponse{}
	err := c.post(ctx, "saros/add-liquidity/simulate", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"lpTokenMint":  accounts.LPTokenMint,
		},
		Data: map[string]interface{}{
			"tokenAAmount": data.TokenAAmount,
			"tokenBAmount": data.TokenBAmount,
		},
	}, &result)
	return &result, err
}
