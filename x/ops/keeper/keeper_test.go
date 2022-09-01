package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tharsis/ethermint/app"
	opsKeeper "github.com/tharsis/ethermint/x/ops/keeper"
	"github.com/tharsis/ethermint/x/ops/types"
)

type KeeperTestSuite struct {
	suite.Suite
	app         *app.EthermintApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	testApp := app.Setup(suite.T(), false, func(ea *app.EthermintApp, genesis simapp.GenesisState) simapp.GenesisState {
		return genesis
	})
	ctx := testApp.BaseApp.NewContext(false, tmproto.Header{})

	querier := opsKeeper.Querier{Keeper: testApp.OpsKeeper}

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, testApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app, suite.ctx, suite.queryClient = testApp, ctx, queryClient
}

func (suite *KeeperTestSuite) TestRecords() {
	records := suite.app.OpsKeeper.ListRecords(suite.ctx)
	require.Equal(suite.T(), len(records), 0)
}

func (suite *KeeperTestSuite) TestCreateRecord() {
	newRecord := types.NameRecord{
		Name: "Murali",
		Age:  22,
	}
	recordCreated := suite.app.OpsKeeper.SetNameRecord(suite.ctx, newRecord.Name, newRecord.Age)
	newRecord.Id = recordCreated.Id
	require.Equal(suite.T(), newRecord, recordCreated)
}

func (suite *KeeperTestSuite) TestGetRecordById() {
	newRecord := types.NameRecord{
		Name: "Murali",
		Age:  22,
	}
	recordCreated := suite.app.OpsKeeper.SetNameRecord(suite.ctx, newRecord.Name, newRecord.Age)
	record := suite.app.OpsKeeper.GetNameRecordByID(suite.ctx, recordCreated.Id)
	require.Equal(suite.T(), record, recordCreated)
}

func (suite *KeeperTestSuite) TestListRecords() {
	newRecords := []*types.NameRecord{
		{
			Name: "Murali",
			Age:  22,
		},
		{
			Name: "Sai",
			Age:  27,
		},
	}
	for _, record := range newRecords {
		recordCreated := suite.app.OpsKeeper.SetNameRecord(suite.ctx, record.Name, record.Age)
		record.Id = recordCreated.Id
	}

	recordList := suite.app.OpsKeeper.ListRecords(suite.ctx)
	require.Equal(suite.T(), newRecords, recordList)
}
func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
