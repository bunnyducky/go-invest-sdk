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
	Owner               solana.PublicKey
	Payer               solana.PublicKey
	Pool                string
	LpTokenMint         solana.PublicKey
	SourceTokenAccounts []solana.PublicKey
}

type MercurialStakeData struct {
	StakeAmounts  []uint64
	MinMintAmount uint64
}

func (c *Client) MercurialStake(ctx context.Context, accounts MercurialStakeAccounts, data MercurialStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":        accounts.Owner,
			"payer":               accounts.Payer,
			"pool":                accounts.Pool,
			"lpTokenMint":         accounts.LpTokenMint,
			"sourceTokenAccounts": accounts.SourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  data.StakeAmounts,
			"minMintAmount": data.MinMintAmount,
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
			"ownerAccount":        accounts.Owner,
			"payer":               accounts.Payer,
			"pool":                accounts.Pool,
			"lpTokenMint":         accounts.LpTokenMint,
			"sourceTokenAccounts": accounts.SourceTokenAccounts,
		},
		Data: map[string]interface{}{
			"stakeAmounts":  data.StakeAmounts,
			"minMintAmount": data.MinMintAmount,
		},
	}, &result)
	return &result, err
}

type MercurialUnstakeAccounts struct {
	Owner                   solana.PublicKey
	Payer                   solana.PublicKey
	Pool                    string
	LpTokenMint             solana.PublicKey
	DestinationTokenMint    solana.PublicKey
	DestinationTokenAccount solana.PublicKey
}

type MercurialUnstakeData struct {
	UnstakeAmount   uint64
	MiniumOutAmount uint64
}

func (c *Client) MercurialUnstake(ctx context.Context, accounts MercurialUnstakeAccounts, data MercurialUnstakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "mercurial/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            accounts.Owner,
			"payer":                   accounts.Payer,
			"pool":                    accounts.Pool,
			"lpTokenMint":             accounts.LpTokenMint,
			"destinationTokenMint":    accounts.DestinationTokenMint,
			"destinationTokenAccount": accounts.DestinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount":   data.UnstakeAmount,
			"miniumOutAmount": data.MiniumOutAmount,
		},
	}, &result)

	return &result, err
}
