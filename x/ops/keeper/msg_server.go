package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/ethermint/x/ops/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bond MsgServer interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Create(c context.Context, msg *types.MsgCreateNameRecord) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	resp := k.Keeper.SetNameRecord(ctx, msg.Name, msg.Age)
	return &types.MsgCreateResponse{
		Id:   resp.Id,
		Name: resp.Name,
		Age:  resp.Age,
	}, nil
}

func (k msgServer) Update(c context.Context, msg *types.MsgUpdateNameRecord) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	resp := k.Keeper.UpdateNameRecord(ctx, msg.Id, msg.Name, msg.Age)
	return &types.MsgUpdateResponse{
		Id:   resp.Id,
		Name: resp.Name,
		Age:  resp.Age,
	}, nil
}

func (k msgServer) Delete(c context.Context, msg *types.MsgDeleteNameRecord) (*types.MsgDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	k.Keeper.DeleteRecord(ctx, msg.Id)
	return &types.MsgDeleteResponse{}, nil
}
