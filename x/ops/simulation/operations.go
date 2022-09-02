package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tharsis/ethermint/x/ops/keeper"
	"github.com/tharsis/ethermint/x/ops/types"
)

const (
	OpWeightMsgCreateNameRecord = "op_weight_msg_create_name_record"
	OpWeightMsgUpdateNameRecord = "op_weight_msg_update_name_record"
	OpWeightMsgDeleteNameRecord = "op_weight_msg_delete_name_record"
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak simulation.AccountKeeper, bk simulation.BankKeeper, k *keeper.Keeper,
) simulation.WeightedOperations {

	var (
		weightMsgCreateNameRecord int
		weightMsgUpdateNameRecord int
		weightMsgDeleteNameRecord int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateNameRecord, &weightMsgCreateNameRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNameRecord = simappparams.DefaultWeightMsgSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateNameRecord, &weightMsgUpdateNameRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNameRecord = simappparams.DefaultWeightMsgSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteNameRecord, &weightMsgDeleteNameRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteNameRecord = simappparams.DefaultWeightMsgSend
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateNameRecord,
			SimulateMsgCreateNameRecord(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateNameRecord,
			SimulateMsgUpdateNameRecord(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDeleteNameRecord,
			SimulateMsgDeleteNameRecord(ak, bk, k),
		),
	}
}

func SimulateMsgCreateNameRecord(ak simulation.AccountKeeper, bk simulation.BankKeeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accounts)

		randStrLength := simtypes.RandIntBetween(r, 0, 100)
		randName := simtypes.RandStringOfLength(r, randStrLength)
		randAge, err := simtypes.RandPositiveInt(r, sdk.NewInt(100))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.MsgCreateNameRecord{}.Type(), "unable to generate random age"), nil, nil
		}
		msg := types.NewMsgCreateRecord(simAccount.Address, randName, randAge.Uint64())

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           &msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			ModuleName:    types.ModuleName,
			AccountKeeper: ak,
			Bankkeeper:    bk,
		}

		return simulation.GenAndDeliverTx(txCtx, sdk.Coins{})
	}
}

func SimulateMsgUpdateNameRecord(ak simulation.AccountKeeper, bk simulation.BankKeeper, k *keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accounts)

		records := k.ListRecords(ctx)
		randRecordIndex := simtypes.RandIntBetween(r, 0, len(records)-1)

		randStrLength := simtypes.RandIntBetween(r, 0, 100)
		randName := simtypes.RandStringOfLength(r, randStrLength)
		randAge, err := simtypes.RandPositiveInt(r, sdk.NewInt(100))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.MsgUpdateNameRecord{}.Type(), "unable to generate random age"), nil, nil
		}

		msg := types.NewMsgUpdateRecord(simAccount.Address, records[randRecordIndex].Id, randName, randAge.Uint64())

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           &msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			ModuleName:    types.ModuleName,
			AccountKeeper: ak,
			Bankkeeper:    bk,
		}

		return simulation.GenAndDeliverTx(txCtx, sdk.Coins{})
	}
}

func SimulateMsgDeleteNameRecord(ak simulation.AccountKeeper, bk simulation.BankKeeper, k *keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accounts)

		records := k.ListRecords(ctx)
		randRecordIndex := simtypes.RandIntBetween(r, 0, len(records)-1)

		msg := types.NewMsgDeleteRecord(simAccount.Address, records[randRecordIndex].Id)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           &msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			ModuleName:    types.ModuleName,
			AccountKeeper: ak,
			Bankkeeper:    bk,
		}

		return simulation.GenAndDeliverTx(txCtx, sdk.Coins{})
	}
}
