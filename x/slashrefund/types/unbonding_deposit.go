package types

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AddEntry - append entry to the unbonding deposit
func (ubd *UnbondingDeposit) AddEntry(creationHeight int64, minTime time.Time, balance math.Int) {
	// Check the entries exists with creation_height and complete_time
	entryIndex := -1
	for index, ubdEntry := range ubd.Entries {
		if ubdEntry.CreationHeight == creationHeight && ubdEntry.CompletionTime.Equal(minTime) {
			entryIndex = index
			break
		}
	}
	// entryIndex exists
	if entryIndex != -1 {
		ubdEntry := ubd.Entries[entryIndex]
		ubdEntry.Balance = ubdEntry.Balance.Add(balance)
		ubdEntry.InitialBalance = ubdEntry.InitialBalance.Add(balance)

		// update the entry
		ubd.Entries[entryIndex] = ubdEntry
	} else {
		// append the new unbond deposit entry
		entry := NewUnbondingDepositEntry(creationHeight, minTime, balance)
		ubd.Entries = append(ubd.Entries, entry)
	}
}

func NewUnbondingDepositEntry(creationHeight int64, completionTime time.Time, balance math.Int) UnbondingDepositEntry {
	return UnbondingDepositEntry{
		CreationHeight: creationHeight,
		CompletionTime: completionTime,
		InitialBalance: balance,
		Balance:        balance,
	}
}

func NewUnbondingDeposit(
	depositorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
	creationHeight int64, minTime time.Time, balance math.Int,
) UnbondingDeposit {
	return UnbondingDeposit{
		DepositorAddress: depositorAddr.String(),
		ValidatorAddress: validatorAddr.String(),
		Entries: []UnbondingDepositEntry{
			NewUnbondingDepositEntry(creationHeight, minTime, balance),
		},
	}
}
