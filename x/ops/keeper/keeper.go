package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/ethermint/x/ops/types"
)

var (
	RecordIdPrefix = "Record-"
	RecordCounter  = 0
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func (k Keeper) SetNameRecord(ctx sdk.Context, msg types.MsgCreateNameRecord) {
	store := ctx.KVStore(k.storeKey)
	id := fmt.Sprintf("%s%d", RecordIdPrefix, RecordCounter)
	record := types.NameRecord{
		Id:   id,
		Name: msg.Name,
		Age:  msg.Age,
	}
	store.Set([]byte(id), k.cdc.MustMarshal(&record))
	RecordCounter++
}

func (k Keeper) UpdateNameRecord(ctx sdk.Context, msg types.MsgUpdateNameRecord) {
	store := ctx.KVStore(k.storeKey)
	record := types.NameRecord{
		Id:   msg.Id,
		Name: msg.Name,
		Age:  msg.Age,
	}
	store.Set([]byte(record.Id), k.cdc.MustMarshal(&record))
}

func (k Keeper) DeleteRecord(ctx sdk.Context, msg types.MsgDeleteNameRecord) {

}
