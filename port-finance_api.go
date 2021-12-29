package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

func (c *Client) PortInit(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, tokenMint solana.PublicKey, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/init", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
	}, &result)

	return &result, err
}

func (c *Client) PortStake(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, tokenMint solana.PublicKey, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/stake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
		Data: map[string]interface{}{
			"stakeAmount": stakeAmount,
		},
	}, &result)

	return &result, err
}

func (c *Client) PortUnstake(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "port/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"tokenMint":    tokenMint,
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
