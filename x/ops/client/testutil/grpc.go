package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/testutil/rest"
	"github.com/tharsis/ethermint/x/ops/types"
)

func (s *IntegrationTestSuite) TestGRPCListRecords() {
	val := s.network.Validators[0]
	sr := s.Require()
	reqUrl := fmt.Sprintf("%s/cosmos/ops/v1beta1/records", val.APIAddress)

	testCases := []struct {
		testName string
		url      string
		err      bool
		errMsg   string
		records  []*types.NameRecord
	}{
		{
			"invalid request with headers",
			reqUrl + "asdasdas",
			true,
			"",
			[]*types.NameRecord{},
		},
		{
			"no records",
			reqUrl,
			false,
			"",
			[]*types.NameRecord{},
		},
		{
			"valid request",
			reqUrl,
			false,
			"",
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
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			resp, err := rest.GetRequest(tc.url)
			if tc.err {
				sr.Error(err)
				sr.Contains(string(resp), tc.errMsg)
			} else {
				sr.NoError(err)
				var response types.QueryAllRecordsResponse
				err := val.ClientCtx.Codec.UnmarshalJSON(resp, &response)
				sr.NoError(err)
				sr.Equal(response.Records, tc.records)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGRPCGetRecordByID() {
	val := s.network.Validators[0]
	sr := s.Require()
	reqUrl := fmt.Sprintf("%s/cosmos/ops/v1beta1/record/", val.APIAddress) + "%s"

	testCases := []struct {
		testName string
		url      string
		err      bool
		record   types.NameRecord
	}{
		{
			"empty list",
			fmt.Sprintf(reqUrl, "asdasd"),
			true,
			types.NameRecord{},
		},
		{
			"valid request",
			fmt.Sprintf(reqUrl, "Record-0"),
			false,
			types.NameRecord{Name: "Murali", Age: 22},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			resp, err := rest.GetRequest(tc.url)
			if tc.err {
				sr.Error(err)
			} else {
				sr.NoError(err)
				var response types.QueryRecordResponse
				err := val.ClientCtx.Codec.UnmarshalJSON(resp, &response)
				sr.NoError(err)
				sr.Equal(response.Record, tc.record)
			}
		})
	}
}
