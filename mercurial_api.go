package investsdk

import (
	"context"

	solana "github.com/gagliardetto/solana-go"
)

func (c *Client) TriggerMerClaim(ctx context.Context, pda solana.PublicKey) error {
	result := map[string]interface{}{}
	return c.post(ctx, "mercurial/claim", map[string]solana.PublicKey{"address": pda}, &result)
}

func (c *Client) MercurialStakeData(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, pool string, sourceTokenAccounts []solana.PublicKey, stakeAmounts []uint64, minMintAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":        owner,
			"payer":               payer,
			"pool":                pool,
			"sourceTokenAccounts": sourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  stakeAmounts,
			"minMintAmount": minMintAmount,
		},
	}, &result)
	return &result, err
}

type MercurialStakeEstimateResponse struct {
	MinAmount uint32 `json:"minAmount"`
}

func (c *Client) MercurialStakeEstimate(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, pool string, sourceTokenAccounts []solana.PublicKey, stakeAmounts []uint64, minMintAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/stake/estimate", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":        owner,
			"payer":               payer,
			"pool":                pool,
			"sourceTokenAccounts": sourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  stakeAmounts,
			"minMintAmount": minMintAmount,
		},
	}, &result)
	return &result, err
}

func (c *Client) MercurialUnstakeData(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, pool string, destinationTokenMint solana.PublicKey, destinationTokenAccount solana.PublicKey, unstakeAmount uint64, miniumOutAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            owner,
			"payer":                   payer,
			"pool":                    pool,
			"destinationTokenMint":    destinationTokenMint,
			"destinationTokenAccount": destinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount":   unstakeAmount,
			"miniumOutAmount": miniumOutAmount,
		},
	}, &result)

	return &result, err
}
