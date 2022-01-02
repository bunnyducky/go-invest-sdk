package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type PortInitAccounts struct {
	Owner     solana.PublicKey
	Payer     solana.PublicKey
	TokenMint solana.PublicKey
}

func (c *Client) PortInit(ctx context.Context, accounts PortInitAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/init", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
	}, &result)

	return &result, err
}

type PortStakeAccounts struct {
	Owner     solana.PublicKey
	Payer     solana.PublicKey
	TokenMint solana.PublicKey
}

func (c *Client) PortStake(ctx context.Context, accounts PortStakeAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmount": stakeAmount,
		},
	}, &result)

	return &result, err
}

type PortUnstakeAccounts struct {
	Owner     solana.PublicKey
	TokenMint solana.PublicKey
}

func (c *Client) PortUnstake(ctx context.Context, accounts PortUnstakeAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"tokenMint":    accounts.TokenMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}

func (c *Client) PortClaim(ctx context.Context, owner solana.PublicKey) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/claim", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
		},
	}, &result)

	return &result, err
}
