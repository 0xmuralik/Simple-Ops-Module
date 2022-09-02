package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tharsis/ethermint/x/ops/types"
)

const (
	recordCounterKey = "record_counter"
	recordsKey       = "records"
)

func genRecordCounter(r *rand.Rand) uint64 {
	counter, err := simtypes.RandPositiveInt(r, sdk.NewInt(100))
	if err != nil {
		panic(fmt.Errorf("error while generating random record counter, Error: %w", err))
	}
	return counter.Uint64()
}

func genRecords(r *rand.Rand, counter uint64) []*types.NameRecord {
	var records []*types.NameRecord
	for i := 0; i < int(counter); i++ {
		randStrLength := simtypes.RandIntBetween(r, 0, 100)
		randName := simtypes.RandStringOfLength(r, randStrLength)
		randAge, err := simtypes.RandPositiveInt(r, sdk.NewInt(100))
		if err != nil {
			panic(fmt.Errorf("cannot generate random age in genRecords genesis. Error: %w", err))
		}
		record := &types.NameRecord{
			Name: randName,
			Age:  randAge.Uint64(),
		}
		records = append(records, record)
	}
	return records
}

func RandomizedGenState(simState *module.SimulationState) {

	var (
		recordCounter uint64
		records       []*types.NameRecord
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, recordCounterKey, &recordCounter, simState.Rand, func(r *rand.Rand) {
			recordCounter = genRecordCounter(r)
		},
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, recordsKey, &records, simState.Rand, func(r *rand.Rand) {
			records = genRecords(r, recordCounter)
		},
	)

	opsGenesis := types.GenesisState{
		RecordCounter: recordCounter,
		Records:       records,
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&opsGenesis)
}
