package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tharsis/ethermint/testutil/network"
	"github.com/tharsis/ethermint/x/ops/client/cli"
	"github.com/tharsis/ethermint/x/ops/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg            network.Config
	network        *network.Network
	accountName    string
	accountAddress string
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("Setting up integration test suite")

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.accountName = "accountName"
	s.createAccountWithBalance(s.accountName)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) createAccountWithBalance(accountName string) {
	val := s.network.Validators[0]
	sr := s.Require()
	consPrivKey := ed25519.GenPrivKey()
	consPubKeyBz, err := s.cfg.Codec.MarshalInterfaceJSON(consPrivKey.PubKey())
	sr.NoError(err)
	sr.NotNil(consPubKeyBz)
	info, _, err := val.ClientCtx.Keyring.NewMnemonic(accountName, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	sr.NoError(err)

	newAddr, _ := info.GetAddress()
	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(200))),
		fmt.Sprintf("--%s", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	sr.NoError(err)
	s.accountAddress = newAddr.String()
}

func (s *IntegrationTestSuite) creatNameRecord(name string, age uint64) {
	val := s.network.Validators[0]
	sr := s.Require()
	createOpsCmd := cli.NewCreateRecordCmd()
	args := []string{
		name,
		fmt.Sprintf("%d", age),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.accountName),
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("3%s", s.cfg.BondDenom)),
	}
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, createOpsCmd, args)
	sr.NoError(err)
	var d sdk.TxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &d)
	sr.NoError(err)
	sr.Zero(d.Code)
	err = s.network.WaitForNextBlock()
	sr.NoError(err)
}

func (s *IntegrationTestSuite) TestListNameRecords() {
	val := s.network.Validators[0]

	testCases := []struct {
		name    string
		records []*types.NameRecord
	}{
		{
			"No records",
			[]*types.NameRecord{},
		},
		{
			"2 records",
			[]*types.NameRecord{
				{
					Id:   "Record-0",
					Name: "Murali",
					Age:  22,
				},
				{
					Id:   "Record-1",
					Name: "Sai",
					Age:  28,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			clientCtx := val.ClientCtx
			for _, record := range tc.records {
				s.creatNameRecord(record.Name, record.Age)
			}
			cmd := cli.ListRecorsCmd()
			args := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)
			var queryResponse types.QueryAllRecordsResponse
			err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &queryResponse)
			s.Require().NoError(err)
			s.Require().Equal(tc.records, queryResponse.Records)
		})
	}
}

func (s *IntegrationTestSuite) TestGetRecordByID() {
	val := s.network.Validators[0]

	testCases := []struct {
		name    string
		record  types.NameRecord
		id      string
		success bool
	}{
		{
			"Get record by id",
			types.NameRecord{
				Id:   "Record-0",
				Name: "Murali",
				Age:  22,
			},
			"Record-0",
			true,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			clientCtx := val.ClientCtx
			s.creatNameRecord(tc.record.Name, tc.record.Age)

			cmd := cli.GetRecordByIDCmd()
			args := []string{
				tc.id,
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			s.Require().NoError(err)

			var queryResponse types.QueryRecordResponse
			err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &queryResponse)
			s.Require().NoError(err)
		})
	}
}
