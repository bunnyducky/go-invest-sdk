package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type JupiterSwapAccount struct {
	Owner           solana.PublicKey
	InputTokenMint  solana.PublicKey
	OutputTokenMint solana.PublicKey
}

type JupiterSwapData struct {
	InputAmount float32
	Slippage    float32
}

func (c *Client) JupiterSwapData(ctx context.Context, accounts JupiterSwapAccount, data JupiterSwapData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "jup-ag/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    accounts.Owner,
			"inputTokenMint":  accounts.InputTokenMint,
			"outputTokenMint": accounts.OutputTokenMint,
		},
		Data: map[string]interface{}{
			"inputAmount": data.InputAmount,
			"slipage":     data.Slippage,
		},
	}, &result)

	return &result, err
}
