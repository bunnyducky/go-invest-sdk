package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

func (c *Client) RaydiumSwap(ctx context.Context, owner solana.PublicKey, sourceTokenMint solana.PublicKey, destTokenMint solana.PublicKey, amountIn uint64, minAmountout uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "raydium/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    owner,
			"sourceTokenMint": sourceTokenMint,
			"destTokenMint":   destTokenMint,
		},
		Data: map[string]interface{}{
			"amountIn":     amountIn,
			"minAmountOut": minAmountout,
		},
	}, &result)

	return &result, err
}
