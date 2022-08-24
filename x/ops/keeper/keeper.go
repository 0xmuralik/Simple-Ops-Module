package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/ethermint/x/ops/types"
)

var (
	RecordIDPrefix   = "Record-"
	RecordCounterKey = []byte{0x0}
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func getRecordId(index uint64) string {
	return fmt.Sprintf("%s%d", RecordIDPrefix, index)
}

func (k Keeper) getRecordCounter(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(RecordCounterKey)
	counter := binary.LittleEndian.Uint64(bz)
	return counter
}

func (k Keeper) incrementRecordCounter(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(RecordCounterKey)
	counter := binary.LittleEndian.Uint64(bz)
	counter++

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, counter)
	store.Set(RecordCounterKey, b)
}

func (k Keeper) SetNameRecord(ctx sdk.Context, name string, age uint64) types.NameRecord {
	store := ctx.KVStore(k.storeKey)
	id := getRecordId(k.getRecordCounter(ctx))
	record := types.NameRecord{
		Id:   id,
		Name: name,
		Age:  age,
	}
	store.Set([]byte(id), k.cdc.MustMarshal(&record))
	k.incrementRecordCounter(ctx)
	return record
}

func (k Keeper) UpdateNameRecord(ctx sdk.Context, id string, name string, age uint64) types.NameRecord {
	store := ctx.KVStore(k.storeKey)
	record := types.NameRecord{
		Id:   id,
		Name: name,
		Age:  age,
	}
	store.Set([]byte(record.Id), k.cdc.MustMarshal(&record))
	return record
}

func (k Keeper) DeleteRecord(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id))
}
