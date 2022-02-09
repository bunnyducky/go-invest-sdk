package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type QuarryAccounts struct {
	Owner     solana.PublicKey
	Payer     solana.PublicKey
	Rewarder  solana.PublicKey
	TokenMint solana.PublicKey
}

func (c *Client) QuarryStake(ctx context.Context, accounts QuarryAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "quarry/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"rewarderKey":  accounts.Rewarder,
			"tokenMint":    accounts.TokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmount": stakeAmount,
		},
	}, &result)
	return &result, err
}

func (c *Client) QuarryUnstake(ctx context.Context, accounts QuarryAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "quarry/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"rewarderKey":  accounts.Rewarder,
			"tokenMint":    accounts.TokenMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)
	return &result, err
}

func (c *Client) QuarryClaim(ctx context.Context, accounts QuarryAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "quarry/claim", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"rewarderKey":  accounts.Rewarder,
			"tokenMint":    accounts.TokenMint,
		},
	}, &result)
	return &result, err
}
