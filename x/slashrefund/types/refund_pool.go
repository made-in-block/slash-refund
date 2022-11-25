package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDRefundPool(validatorAddr sdk.ValAddress, tokens sdk.Coin, shares sdk.Dec) RefundPool {
	return RefundPool{
		OperatorAddress: validatorAddr.String(),
		Tokens:          tokens,
		Shares:          shares,
	}
}

func (p RefundPool) SharesFromTokens(tokens sdk.Coin) (sdk.Dec, error) {
	if p.Tokens.IsZero() {
		return sdk.ZeroDec(), ErrInsufficientShares
	}
	return p.Shares.MulInt(tokens.Amount).QuoInt(p.GetTokens().Amount), nil
}

func (p RefundPool) SharesFromTokensTruncated(tokens sdk.Coin) (sdk.Dec, error) {
	if p.Tokens.IsZero() {
		return sdk.ZeroDec(), ErrInsufficientShares
	}
	return p.Shares.MulInt(tokens.Amount).QuoTruncate(sdk.NewDecFromInt(p.GetTokens().Amount)), nil
}

func (p RefundPool) TokensFromShares(shares sdk.Dec) sdk.Dec {
	return (shares.MulInt(p.Tokens.Amount)).Quo(p.Shares)
}