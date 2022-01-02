package investsdk

import (
	"context"

	solana "github.com/gagliardetto/solana-go"
)

func (c *Client) TriggerMerClaim(ctx context.Context, pda solana.PublicKey) error {
	result := map[string]interface{}{}
	return c.post(ctx, "mercurial/claim", map[string]solana.PublicKey{"address": pda}, &result)
}

type MercurialStakeAccounts struct {
	owner               solana.PublicKey
	payer               solana.PublicKey
	pool                string
	sourceTokenAccounts []solana.PublicKey
}

type MercurialStakeData struct {
	stakeAmounts  []uint64
	minMintAmount uint64
}

func (c *Client) MercurialStake(ctx context.Context, accounts MercurialStakeAccounts, data MercurialStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":        accounts.owner,
			"payer":               accounts.payer,
			"pool":                accounts.pool,
			"sourceTokenAccounts": accounts.sourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  data.stakeAmounts,
			"minMintAmount": data.minMintAmount,
		},
	}, &result)
	return &result, err
}

type MercurialStakeEstimateResponse struct {
	MinAmount uint32 `json:"minAmount"`
}

func (c *Client) MercurialStakeEstimate(ctx context.Context, accounts MercurialStakeAccounts, data MercurialStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/stake/estimate", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":        accounts.owner,
			"payer":               accounts.payer,
			"pool":                accounts.pool,
			"sourceTokenAccounts": accounts.sourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  data.stakeAmounts,
			"minMintAmount": data.minMintAmount,
		},
	}, &result)
	return &result, err
}

type MercurialUnstakeAccounts struct {
	owner                   solana.PublicKey
	payer                   solana.PublicKey
	pool                    string
	destinationTokenMint    solana.PublicKey
	destinationTokenAccount solana.PublicKey
}

type MercurialUnstakeData struct {
	unstakeAmount   uint64
	miniumOutAmount uint64
}

func (c *Client) MercurialUnstake(ctx context.Context, accounts MercurialUnstakeAccounts, data MercurialUnstakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            accounts.owner,
			"payer":                   accounts.payer,
			"pool":                    accounts.pool,
			"destinationTokenMint":    accounts.destinationTokenMint,
			"destinationTokenAccount": accounts.destinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount":   data.unstakeAmount,
			"miniumOutAmount": data.miniumOutAmount,
		},
	}, &result)

	return &result, err
}
