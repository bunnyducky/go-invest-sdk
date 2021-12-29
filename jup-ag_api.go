package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

func (c *Client) JupiterSwapData(ctx context.Context, owner solana.PublicKey, inputTokenMint solana.PublicKey, outputTokenMint solana.PublicKey, inputAmount uint64, slippage int16) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "jup-ag/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    owner,
			"inputTokenMint":  inputTokenMint,
			"outputTokenMint": outputTokenMint,
		},
		Data: map[string]interface{}{
			"inputAmount": inputAmount,
			"slipage":     slippage,
		},
	}, &result)

	return &result, err
}
