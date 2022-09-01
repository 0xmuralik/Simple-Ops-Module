package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tharsis/ethermint/x/ops/client/cli"
	"github.com/tharsis/ethermint/x/ops/types"
)

func (s *IntegrationTestSuite) TestTxCreateNameRecord() {
	val := s.network.Validators[0]
	sr := s.Require()

	testCases := []struct {
		testName string
		name     string
		age      uint64
		err      bool
	}{
		{
			"create name record",
			"Murali",
			22,
			false,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			clientCtx := val.ClientCtx
			cmd := cli.NewCreateRecordCmd()
			args := []string{
				tc.name,
				fmt.Sprintf("%d", tc.age),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.accountName),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("3%s", s.cfg.BondDenom)),
			}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.err {
				sr.Error(err)
			} else {
				sr.NoError(err)
				var d sdk.TxResponse
				err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &d)
				sr.Nil(err)
				sr.NoError(err)
				sr.Zero(d.Code)

				queryArgs := []string{
					"Record-0",
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				queryOut, err := clitestutil.ExecTestCLICmd(clientCtx, cli.GetRecordByIDCmd(), queryArgs)
				sr.NoError(err)

				var queryResponse types.QueryRecordResponse
				err = clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryResponse)
				sr.NoError(err)

				sr.Equal(queryResponse, types.QueryRecordResponse{Record: &types.NameRecord{Id: "Record-0", Name: tc.name, Age: tc.age}})
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateRecord() {
	val := s.network.Validators[0]
	sr := s.Require()

	testCases := []struct {
		testName string
		id       string
		name     string
		age      uint64
		err      bool
	}{
		{
			"update record",
			"Record-0",
			"Sai",
			22,
			false,
		},
		{
			"wrong id",
			"Record-1",
			"Sai",
			22,
			true,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			clientCtx := val.ClientCtx
			cmd := cli.UpdateRecordCmd()

			s.creatNameRecord("Murali", 22)
			args := []string{
				tc.id,
				tc.name,
				fmt.Sprintf("%d", tc.age),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.accountName),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("3%s", s.cfg.BondDenom)),
			}
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.err {
				sr.Error(err)
			} else {
				sr.NoError(err)
				var d sdk.TxResponse
				err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &d)
				sr.Nil(err)
				sr.NoError(err)
				sr.Zero(d.Code)

				queryArgs := []string{
					"Record-0",
					fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				}
				queryOut, err := clitestutil.ExecTestCLICmd(clientCtx, cli.GetRecordByIDCmd(), queryArgs)
				sr.NoError(err)

				var queryResponse types.QueryRecordResponse
				err = clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryResponse)
				sr.NoError(err)

				sr.Equal(queryResponse, types.QueryRecordResponse{Record: &types.NameRecord{Id: "Record-0", Name: tc.name, Age: tc.age}})
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxDeleteRecord() {
	val := s.network.Validators[0]
	sr := s.Require()

	testCases := []struct {
		testName string
		id       string
		err      bool
	}{
		{
			"delete record",
			"Record-0",
			false,
		},
		{
			"wrong id",
			"Record-10",
			true,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			clientCtx := val.ClientCtx
			cmd := cli.DeleteRecordCmd()

			s.creatNameRecord("Murali", 22)
			args := []string{
				tc.id,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.accountName),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
				fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("3%s", s.cfg.BondDenom)),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.err {
				sr.Error(err)
			} else {
				sr.NoError(err)
				var d sdk.TxResponse
				err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &d)
				sr.Nil(err)
				sr.NoError(err)
				sr.Zero(d.Code)

				queryArgs := []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)}
				queryOut, err := clitestutil.ExecTestCLICmd(clientCtx, cli.ListRecorsCmd(), queryArgs)
				sr.NoError(err)

				var queryResponse types.QueryAllRecordsResponse
				err = clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryResponse)
				sr.NoError(err)
				sr.Equal(0, len(queryResponse.Records))
			}
		})
	}
}
