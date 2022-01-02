package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type PortInitAccounts struct {
	owner     solana.PublicKey
	payer     solana.PublicKey
	tokenMint solana.PublicKey
}

func (c *Client) PortInit(ctx context.Context, accounts PortInitAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/init", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.owner,
			"payer":        accounts.payer,
			"tokenMint":    accounts.tokenMint,
		},
	}, &result)

	return &result, err
}

type PortStakeAccounts struct {
	owner     solana.PublicKey
	payer     solana.PublicKey
	tokenMint solana.PublicKey
}

func (c *Client) PortStake(ctx context.Context, accounts PortStakeAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.owner,
			"payer":        accounts.payer,
			"tokenMint":    accounts.tokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmount": stakeAmount,
		},
	}, &result)

	return &result, err
}

type PortUnstakeAccounts struct {
	owner     solana.PublicKey
	tokenMint solana.PublicKey
}

func (c *Client) PortUnstake(ctx context.Context, accounts PortUnstakeAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.owner,
			"tokenMint":    accounts.tokenMint,
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
