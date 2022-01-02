package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type JupiterSwapAccount struct {
	owner           solana.PublicKey
	inputTokenMint  solana.PublicKey
	outputTokenMint solana.PublicKey
}

type JupiterSwapData struct {
	inputAmount uint64
	slippage    int16
}

func (c *Client) JupiterSwapData(ctx context.Context, accounts JupiterSwapAccount, data JupiterSwapData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "jup-ag/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    accounts.owner,
			"inputTokenMint":  accounts.inputTokenMint,
			"outputTokenMint": accounts.outputTokenMint,
		},
		Data: map[string]interface{}{
			"inputAmount": data.inputAmount,
			"slipage":     data.slippage,
		},
	}, &result)

	return &result, err
}
