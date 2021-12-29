package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

func (c *Client) SaberStake(ctx context.Context, owner solana.PublicKey, lpTokenMint solana.PublicKey, stakeAmountA uint64, stakeAmountB uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"lpTokenMint":  lpTokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmountA": stakeAmountA,
			"stakeAmountB": stakeAmountB,
		},
	}, &result)

	return &result, err
}

func (c *Client) SaberUnstake(ctx context.Context, owner solana.PublicKey, lpTokenMint solana.PublicKey, withdrawMint solana.PublicKey, destinationTokenAccount solana.PublicKey, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "saber/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount":            owner,
			"lpTokenMint":             lpTokenMint,
			"withdrawMint":            withdrawMint,
			"destinationTokenAccount": destinationTokenAccount,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}
