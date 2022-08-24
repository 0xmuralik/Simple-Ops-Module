package ops

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tharsis/ethermint/x/ops/keeper"
	"github.com/tharsis/ethermint/x/ops/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {

	k.SetRecordCounter(ctx, data.RecordCounter)

	for _, record := range data.Records {
		k.SaveRecord(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	records := k.ListRecords(ctx)
	counter := k.GetRecordCounter(ctx)

	return types.GenesisState{Records: records, RecordCounter: counter}
}
