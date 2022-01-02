package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type SaberStakeAccounts struct {
	owner       solana.PublicKey
	lpTokenMint solana.PublicKey
}

type SaberStakeData struct {
	stakeAmountA uint64
	stakeAmountB uint64
}

func (c *Client) SaberStake(ctx context.Context, accounts SaberStakeAccounts, data SaberStakeData) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.owner,
			"lpTokenMint":  accounts.lpTokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmountA": data.stakeAmountA,
			"stakeAmountB": data.stakeAmountB,
		},
	}, &result)

	return &result, err
}

type SaberUnstakeAccounts struct {
	owner                   solana.PublicKey
	lpTokenMint             solana.PublicKey
	withdrawMint            solana.PublicKey
	destinationTokenAccount solana.PublicKey
}

func (c *Client) SaberUnstake(ctx context.Context, accounts SaberUnstakeAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            accounts.owner,
			"lpTokenMint":             accounts.lpTokenMint,
			"withdrawMint":            accounts.withdrawMint,
			"destinationTokenAccount": accounts.destinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}
