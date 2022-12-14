package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDepositPool(validatorAddr sdk.ValAddress, tokens sdk.Coin, shares sdk.Dec) DepositPool {
	return DepositPool{
		OperatorAddress: validatorAddr.String(),
		Tokens:          tokens,
		Shares:          shares,
	}
}

func (d DepositPool) SharesFromTokens(tokens sdk.Coin) (sdk.Dec, error) {
	if d.Tokens.IsZero() {
		return sdk.ZeroDec(), ErrInsufficientShares
	}
	return d.Shares.MulInt(tokens.Amount).QuoInt(d.GetTokens().Amount), nil
}

func (d DepositPool) SharesFromTokensTruncated(tokens sdk.Coin) (sdk.Dec, error) {
	if d.Tokens.IsZero() {
		return sdk.ZeroDec(), ErrInsufficientShares
	}
	return d.Shares.MulInt(tokens.Amount).QuoTruncate(sdk.NewDecFromInt(d.GetTokens().Amount)), nil
}

func (d DepositPool) TokensFromShares(shares sdk.Dec) sdk.Dec {
	return (shares.MulInt(d.Tokens.Amount)).Quo(d.Shares)
}
