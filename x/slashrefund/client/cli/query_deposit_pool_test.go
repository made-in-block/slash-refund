package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/made-in-block/slash-refund/testutil/network"
	"github.com/made-in-block/slash-refund/testutil/nullify"
	"github.com/made-in-block/slash-refund/x/slashrefund/client/cli"
	"github.com/made-in-block/slash-refund/x/slashrefund/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithDepositPoolObjects(t *testing.T, n int) (*network.Network, []types.DepositPool) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		depositPool := types.DepositPool{
			OperatorAddress: strconv.Itoa(i),
		}
		nullify.Fill(&depositPool)
		state.DepositPoolList = append(state.DepositPoolList, depositPool)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.DepositPoolList
}

func TestShowDepositPool(t *testing.T) {
	net, objs := networkWithDepositPoolObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc              string
		idOperatorAddress string

		args []string
		err  error
		obj  types.DepositPool
	}{
		{
			desc:              "found",
			idOperatorAddress: objs[0].OperatorAddress,

			args: common,
			obj:  objs[0],
		},
		{
			desc:              "not found",
			idOperatorAddress: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOperatorAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowDepositPool(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetDepositPoolResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DepositPool)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DepositPool),
				)
			}
		})
	}
}

func TestListDepositPool(t *testing.T) {
	net, objs := networkWithDepositPoolObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDepositPool(), args)
			require.NoError(t, err)
			var resp types.QueryAllDepositPoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DepositPool), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DepositPool),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDepositPool(), args)
			require.NoError(t, err)
			var resp types.QueryAllDepositPoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DepositPool), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DepositPool),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDepositPool(), args)
		require.NoError(t, err)
		var resp types.QueryAllDepositPoolResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.DepositPool),
		)
	})
}
