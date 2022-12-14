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

// NewKeeper creates new instances of the ops Keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func getRecordID(index uint64) string {
	return fmt.Sprintf("%s%d", RecordIDPrefix, index)
}

func (k Keeper) GetRecordCounter(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(RecordCounterKey)
	counter := binary.LittleEndian.Uint64(bz)
	return counter
}

func (k Keeper) SetRecordCounter(ctx sdk.Context, counter uint64) {
	store := ctx.KVStore(k.storeKey)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, counter)
	store.Set(RecordCounterKey, b)
	bz := store.Get(RecordCounterKey)
	counterGet := binary.LittleEndian.Uint64(bz)
	fmt.Println("Counter: ", counterGet)
}

func (k Keeper) SaveRecord(ctx sdk.Context, record *types.NameRecord) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(record.Id), k.cdc.MustMarshal(record))
}

func (k Keeper) SetNameRecord(ctx sdk.Context, name string, age uint64) types.NameRecord {
	store := ctx.KVStore(k.storeKey)
	fmt.Println("In keeper")
	counter := k.GetRecordCounter(ctx)
	fmt.Println("Counter: ", counter)
	id := getRecordID(counter)
	fmt.Println("id: ", id)
	record := types.NameRecord{
		Id:   id,
		Name: name,
		Age:  age,
	}

	fmt.Println("record: ", record)

	store.Set([]byte(id), k.cdc.MustMarshal(&record))
	k.SetRecordCounter(ctx, counter+1)
	return record
}

func (k Keeper) GetNameRecordByID(ctx sdk.Context, id string) types.NameRecord {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(id))
	var record types.NameRecord
	k.cdc.MustUnmarshal(bz, &record)
	return record
}

func (k Keeper) UpdateNameRecord(ctx sdk.Context, id string, name string, age uint64) types.NameRecord {
	store := ctx.KVStore(k.storeKey)
	record := types.NameRecord{
		Id:   id,
		Name: name,
		Age:  age,
	}
	if store.Has([]byte(id)) {
		store.Set([]byte(id), k.cdc.MustMarshal(&record))
		return record
	}
	return types.NameRecord{}
}

func (k Keeper) DeleteRecord(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id))
}

func (k Keeper) ListRecords(ctx sdk.Context) []*types.NameRecord {
	var records []*types.NameRecord

	store := ctx.KVStore(k.storeKey)
	itr := sdk.KVStorePrefixIterator(store, []byte(RecordIDPrefix))
	defer itr.Close()
	for ; itr.Valid(); itr.Next() {
		bz := store.Get(itr.Key())
		if bz != nil {
			var record types.NameRecord
			k.cdc.MustUnmarshal(bz, &record)
			records = append(records, &record)
		}
	}
	return records
}
