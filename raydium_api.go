package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type RaydiumSwapAccounts struct {
	Owner           solana.PublicKey
	SourceTokenMint solana.PublicKey
	DestTokenMint   solana.PublicKey
}

type RaydiumSwapData struct {
	AmountIn     uint64
	MinAmountout uint64
}

func (c *Client) RaydiumSwap(ctx context.Context, accounts RaydiumSwapAccounts, data RaydiumSwapData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "raydium/swap", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":    accounts.Owner,
			"sourceTokenMint": accounts.SourceTokenMint,
			"destTokenMint":   accounts.DestTokenMint,
		},
		Data: map[string]interface{}{
			"amountIn":     data.AmountIn,
			"minAmountOut": data.MinAmountout,
		},
	}, &result)

	return &result, err
}
