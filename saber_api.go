package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type SaberStakeAccounts struct {
	Owner       solana.PublicKey
	LpTokenMint solana.PublicKey
}

type SaberStakeData struct {
	StakeAmountA uint64
	StakeAmountB uint64
}

func (c *Client) SaberStake(ctx context.Context, accounts SaberStakeAccounts, data SaberStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"lpTokenMint":  accounts.LpTokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmountA": data.StakeAmountA,
			"stakeAmountB": data.StakeAmountB,
		},
	}, &result)

	return &result, err
}

type SaberUnstakeAccounts struct {
	Owner                   solana.PublicKey
	LpTokenMint             solana.PublicKey
	WithdrawMint            solana.PublicKey
	DestinationTokenAccount solana.PublicKey
}

func (c *Client) SaberUnstake(ctx context.Context, accounts SaberUnstakeAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            accounts.Owner,
			"lpTokenMint":             accounts.LpTokenMint,
			"withdrawMint":            accounts.WithdrawMint,
			"destinationTokenAccount": accounts.DestinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}
