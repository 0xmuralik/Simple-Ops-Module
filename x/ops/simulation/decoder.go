package simulation

import (
	"bytes"
	"fmt"

	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/tharsis/ethermint/x/ops/keeper"
)

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], keeper.RecordCounterKey):
			counterA := binary.LittleEndian.Uint64(kvA.Value)
			counterB := binary.LittleEndian.Uint64(kvB.Value)
			return fmt.Sprintf("%v\n%v", counterA, counterB)
		default:
			panic(fmt.Sprintf("invalid ops key %X", kvA.Key))
		}
	}
}
