package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

func (c *Client) SunnyInitVault(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey, tokenMint solana.PublicKey, stakeAmountA uint64, stakeAmountB uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/init-vault", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyStake(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, payer solana.PublicKey, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/stake", OperationRequest{
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

func (c *Client) SunnyUnstake(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, payer solana.PublicKey, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyClaimSaberRewards(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, payer solana.PublicKey) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/claim-saber-rewards", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyClaimSunnyRewards(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, payer solana.PublicKey) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/claim-sunny-rewards", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyWithdrawRewardsFromVault(ctx context.Context, owner solana.PublicKey, tokenMint solana.PublicKey, payer solana.PublicKey, platform string) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/withdraw-rewards-from-vault", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
			"tokenMint":    tokenMint,
		},
		Platform: platform,
	}, &result)

	return &result, err
}

func (c *Client) SunnyRedeemSaber(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/redeem-saber", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyRedeemSunny(ctx context.Context, owner solana.PublicKey, payer solana.PublicKey) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/redeem-sunny", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": owner,
			"payer":        payer,
		},
	}, &result)

	return &result, err
}
