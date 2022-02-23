package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type SarosStakeAccounts struct {
	Payer       solana.PublicKey
	Owner       solana.PublicKey
	LPTokenMint solana.PublicKey
}

type SarosStakeData struct {
	StakeAmountA uint64
	StakeAmountB uint64
}

func (c *Client) SarosStake(ctx context.Context, accounts SarosStakeAccounts, data SarosStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saros/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"payer":        accounts.Payer,
			"ownerAccount": accounts.Owner,
			"lpTokenMint":  accounts.LPTokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmountA": data.StakeAmountA,
			"stakeAmountB": data.StakeAmountB,
		},
	}, &result)
	return &result, err
}

type SarosUnstakeAccounts struct {
	Payer        solana.PublicKey
	Owner        solana.PublicKey
	LPTokenMint  solana.PublicKey
	WithdrawMint solana.PublicKey
}

func (c *Client) SarosUnstake(ctx context.Context, accounts SarosUnstakeAccounts, unstakeAmount uint64) (*SimulationResponse, error) {
	result := SimulationResponse{}
	err := c.post(ctx, "saros/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"payer":        accounts.Payer,
			"ownerAccount": accounts.Owner,
			"lpTokenMint":  accounts.LPTokenMint,
			"withdrawMint": accounts.WithdrawMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)
	return &result, err
}
