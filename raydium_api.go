package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type RaydiumSwapAccounts struct {
	owner           solana.PublicKey
	sourceTokenMint solana.PublicKey
	destTokenMint   solana.PublicKey
}

type RaydiumSwapData struct {
	amountIn     uint64
	minAmountout uint64
}

func (c *Client) RaydiumSwap(ctx context.Context, accounts RaydiumSwapAccounts, data RaydiumSwapData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "raydium/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    accounts.owner,
			"sourceTokenMint": accounts.sourceTokenMint,
			"destTokenMint":   accounts.destTokenMint,
		},
		Data: map[string]interface{}{
			"amountIn":     data.amountIn,
			"minAmountOut": data.minAmountout,
		},
	}, &result)

	return &result, err
}
