package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type QuarryStakeAccounts struct {
	Owner       solana.PublicKey
	Payer       solana.PublicKey
	Rewarder    solana.PublicKey
	LpTokenMint solana.PublicKey
}

func (c *Client) QuarryStake(ctx context.Context, accounts QuarryStakeAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "quarry/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"rewarderKey":  accounts.Rewarder,
			"lpTokenMint":  accounts.LpTokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmount": stakeAmount,
		},
	}, &result)
	return &result, err
}

func (c *Client) QuarryUnstake(ctx context.Context, accounts QuarryStakeAccounts, unstakeAmount uint64) (*SimulationResponse, error) {
	result := SimulationResponse{}
	err := c.post(ctx, "quarry/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"rewarderKey":  accounts.Rewarder,
			"lpTokenMint":  accounts.LpTokenMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)
	return &result, err
}
