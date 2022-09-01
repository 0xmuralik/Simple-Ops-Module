package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/ethermint/x/ops/types"
)

type Querier struct {
	Keeper
}

// NewQuerierImpl returns an implementation of the ops QueryServer interface for the provided Keeper.
func NewQuerierImpl(keeper Keeper) types.QueryServer {
	return &Querier{Keeper: keeper}
}

func (q Querier) ListRecords(c context.Context, _ *types.QueryAllRecordsRequest) (*types.QueryAllRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	records := q.Keeper.ListRecords(ctx)
	return &types.QueryAllRecordsResponse{
		Records: records,
	}, nil
}

func (q Querier) GetRecord(c context.Context, req *types.QueryRecordRequest) (*types.QueryRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	record := q.Keeper.GetNameRecordByID(ctx, req.Id)
	return &types.QueryRecordResponse{
		Record: &record,
	}, nil
}

func (q Querier) GetRecordCounter(c context.Context, req *types.QueryRecordCounterRequest) (*types.QueryRecordCounterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	counter := q.Keeper.GetRecordCounter(ctx)
	return &types.QueryRecordCounterResponse{
		Counter: counter,
	}, nil
}
