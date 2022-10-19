package types

const (
	// ModuleName defines the module name
	ModuleName = "slashrefund"

	// StoreKey defines the primary module store key
	kStoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_slashrefund"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	UnbondingDepositKey      = "UnbondingDeposit/value/"
	UnbondingDepositCountKey = "UnbondingDeposit/count/"
)
