package keeper_test

// import (
// 	"strconv"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/query"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	keepertest "github.com/made-in-block/slash-refund/testutil/keeper"
// 	"github.com/made-in-block/slash-refund/testutil/nullify"
// 	"github.com/made-in-block/slash-refund/x/slashrefund/types"
// 	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
// )

// // Prevent strconv unused error
// var _ = strconv.IntSize

// func TestDepositQuerySingle(t *testing.T) {
// 	keeper, ctx := keepertest.SlashrefundKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	_ = wctx
// 	msgs := createNDeposit(keeper, ctx, 2)
// 	for _, tc := range []struct {
// 		desc     string
// 		request  *types.QueryGetDepositRequest
// 		response *types.QueryGetDepositResponse
// 		err      error
// 	}{
// 		{
// 			desc: "First",
// 			request: &types.QueryGetDepositRequest{
// 				DepositorAddress: msgs[0].DepositorAddress,
// 				ValidatorAddress: msgs[0].ValidatorAddress,
// 			},
// 			response: &types.QueryGetDepositResponse{Deposit: msgs[0]},
// 		},
// 		{
// 			desc: "Second",
// 			request: &types.QueryGetDepositRequest{
// 				DepositorAddress: msgs[1].DepositorAddress,
// 				ValidatorAddress: msgs[1].ValidatorAddress,
// 			},
// 			response: &types.QueryGetDepositResponse{Deposit: msgs[1]},
// 		},
// 		{
// 			desc: "KeyNotFound",
// 			request: &types.QueryGetDepositRequest{
// 				DepositorAddress: strconv.Itoa(100000),
// 				ValidatorAddress: strconv.Itoa(100000),
// 			},
// 			err: status.Error(codes.NotFound, "not found"),
// 		},
// 		{
// 			desc: "InvalidRequest",
// 			err:  status.Error(codes.InvalidArgument, "invalid request"),
// 		},
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			valAddr, _ := sdk.ValAddressFromBech32(tc.request.ValidatorAddress)
// 			depAddr, _ := sdk.AccAddressFromBech32(tc.request.DepositorAddress)
// 			validator, _ := types.StakingKeeper.GetValidator(ctx, valAddr)
// 			depCoin := sdk.NewCoin("stake",sdk.NewInt(10))
// 			response, err := keeper.Deposit(ctx, depAddr, depCoin, validator)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t,
// 					nullify.Fill(tc.response),
// 					nullify.Fill(response),
// 				)
// 			}
// 		})
// 	}
// }

// func TestDepositQueryPaginated(t *testing.T) {
// 	keeper, ctx := keepertest.SlashrefundKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	msgs := createNDeposit(keeper, ctx, 5)

// 	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDepositRequest {
// 		return &types.QueryAllDepositRequest{
// 			Pagination: &query.PageRequest{
// 				Key:        next,
// 				Offset:     offset,
// 				Limit:      limit,
// 				CountTotal: total,
// 			},
// 		}
// 	}
// 	t.Run("ByOffset", func(t *testing.T) {
// 		step := 2
// 		for i := 0; i < len(msgs); i += step {
// 			resp, err := keeper.DepositAll(wctx, request(nil, uint64(i), uint64(step), false))
// 			require.NoError(t, err)
// 			require.LessOrEqual(t, len(resp.Deposit), step)
// 			require.Subset(t,
// 				nullify.Fill(msgs),
// 				nullify.Fill(resp.Deposit),
// 			)
// 		}
// 	})
// 	t.Run("ByKey", func(t *testing.T) {
// 		step := 2
// 		var next []byte
// 		for i := 0; i < len(msgs); i += step {
// 			resp, err := keeper.DepositAll(wctx, request(next, 0, uint64(step), false))
// 			require.NoError(t, err)
// 			require.LessOrEqual(t, len(resp.Deposit), step)
// 			require.Subset(t,
// 				nullify.Fill(msgs),
// 				nullify.Fill(resp.Deposit),
// 			)
// 			next = resp.Pagination.NextKey
// 		}
// 	})
// 	t.Run("Total", func(t *testing.T) {
// 		resp, err := keeper.DepositAll(wctx, request(nil, 0, 0, true))
// 		require.NoError(t, err)
// 		require.Equal(t, len(msgs), int(resp.Pagination.Total))
// 		require.ElementsMatch(t,
// 			nullify.Fill(msgs),
// 			nullify.Fill(resp.Deposit),
// 		)
// 	})
// 	t.Run("InvalidRequest", func(t *testing.T) {
// 		_, err := keeper.DepositAll(wctx, nil)
// 		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
// 	})
// }
