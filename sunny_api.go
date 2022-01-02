package investsdk

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type SunnyInitVaultAccounts struct {
	Owner     solana.PublicKey
	Payer     solana.PublicKey
	TokenMint solana.PublicKey
}

func (c *Client) SunnyInitVault(ctx context.Context, accounts SunnyInitVaultAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/init-vault", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
	}, &result)

	return &result, err
}

type SunnyStakeAccounts struct {
	Owner     solana.PublicKey
	TokenMint solana.PublicKey
	Payer     solana.PublicKey
}

func (c *Client) SunnyStake(ctx context.Context, accounts SunnyStakeAccounts, stakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/stake", OperationRequest{
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

type SunnyUnstakeAccounts struct {
	Owner     solana.PublicKey
	TokenMint solana.PublicKey
	Payer     solana.PublicKey
}

func (c *Client) SunnyUnstake(ctx context.Context, accounts SunnyUnstakeAccounts, unstakeAmount uint64) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/unstake", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
		Data: map[string]interface{}{
			"unstakeAmount": unstakeAmount,
		},
	}, &result)

	return &result, err
}

type SunnyClaimAccounts struct {
	Owner     solana.PublicKey
	TokenMint solana.PublicKey
	Payer     solana.PublicKey
}

func (c *Client) SunnyClaimSaberRewards(ctx context.Context, accounts SunnyClaimAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/claim-saber-rewards", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyClaimSunnyRewards(ctx context.Context, accounts SunnyClaimAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/claim-sunny-rewards", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
	}, &result)

	return &result, err
}

type SunnyWithdrawRewardsFromVaultAccounts struct {
	Owner     solana.PublicKey
	TokenMint solana.PublicKey
	Payer     solana.PublicKey
}

func (c *Client) SunnyWithdrawRewardsFromVault(ctx context.Context, accounts SunnyWithdrawRewardsFromVaultAccounts, platform string) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/withdraw-rewards-from-vault", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
			"tokenMint":    accounts.TokenMint,
		},
		Platform: platform,
	}, &result)

	return &result, err
}

type SunnyRedeemAccounts struct {
	Owner solana.PublicKey
	Payer solana.PublicKey
}

func (c *Client) SunnyRedeemSaber(ctx context.Context, accounts SunnyRedeemAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/redeem-saber", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
		},
	}, &result)

	return &result, err
}

func (c *Client) SunnyRedeemSunny(ctx context.Context, accounts SunnyRedeemAccounts) (*OperationResponse, error) {
	result := OperationResponse{}
	err := c.post(ctx, "sunny/redeem-sunny", OperationRequest{
		Accounts: map[string]interface{}{
			"ownerAccount": accounts.Owner,
			"payer":        accounts.Payer,
		},
	}, &result)

	return &result, err
}
