package types_test

import (
	"testing"

	"github.com/made-in-block/slash-refund/x/slashrefund/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DepositList: []types.Deposit{
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
					{
						DepositorAddress: "1",
						ValidatorAddress: "1",
					},
				},
				UnbondingDepositList: []types.UnbondingDeposit{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				UnbondingDepositCount: 2,
				DepositPoolList: []types.DepositPool{
					{
						OperatorAddress: "0",
					},
					{
						OperatorAddress: "1",
					},
				},
				UnbondingDepositList: []types.UnbondingDeposit{
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
					{
						DepositorAddress: "1",
						ValidatorAddress: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated unbondingDeposit",
			genState: &types.GenesisState{
				UnbondingDepositList: []types.UnbondingDeposit{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid unbondingDeposit count",
			genState: &types.GenesisState{
				UnbondingDepositList: []types.UnbondingDeposit{
					{
						Id: 1,
					},
				},
				UnbondingDepositCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated depositPool",
			genState: &types.GenesisState{
				DepositPoolList: []types.DepositPool{
					{
						OperatorAddress: "0",
					},
					{
						OperatorAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated unbondingDeposit",
			genState: &types.GenesisState{
				UnbondingDepositList: []types.UnbondingDeposit{
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
					{
						DepositorAddress: "0",
						ValidatorAddress: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
