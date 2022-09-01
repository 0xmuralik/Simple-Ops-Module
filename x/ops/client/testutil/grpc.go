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
					Name: "Murali",
					Age:  22,
				},
				{
					Name: "Sai",
					Age:  28,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			listBefore := s.getRecordList()
			for _, record := range tc.records {
				s.creatNameRecord(record.Name, record.Age)
				record.Id = fmt.Sprintf("Record-%d", s.getRecordCounter()-1)
			}
			resp, err := rest.GetRequest(tc.url)
			if tc.err {
				sr.Contains(string(resp), tc.errMsg)
			} else {
				sr.NoError(err)
				var response types.QueryAllRecordsResponse
				err := val.ClientCtx.Codec.UnmarshalJSON(resp, &response)
				sr.NoError(err)
				listAfter := append(listBefore, tc.records...)
				sr.Equal(response.Records, listAfter)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGRPCGetRecordByID() {
	val := s.network.Validators[0]
	sr := s.Require()
	reqUrl := fmt.Sprintf("%s/cosmos/ops/v1beta1/record/Record-", val.APIAddress)

	testCases := []struct {
		testName string
		url      string
		err      bool
		errMsg   string
		record   types.NameRecord
	}{
		{
			"empty list",
			fmt.Sprintf(reqUrl, "asdasd"),
			true,
			"",
			types.NameRecord{},
		},
		{
			"valid request",
			fmt.Sprintf(reqUrl+"%d", s.getRecordCounter()),
			false,
			"",
			types.NameRecord{Id: fmt.Sprintf("Record-%d", s.getRecordCounter()), Name: "GRPCGet", Age: 22},
		},
	}
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.testName), func() {
			if tc.record.Name != "" {
				s.creatNameRecord(tc.record.Name, tc.record.Age)
			}
			resp, _ := rest.GetRequest(tc.url)
			if tc.err {
				sr.Contains(string(resp), tc.errMsg)
			} else {
				var response types.QueryRecordResponse
				err := val.ClientCtx.Codec.UnmarshalJSON(resp, &response)
				sr.NoError(err)
				sr.Equal(&tc.record, response.Record)
			}
		})
	}
}
