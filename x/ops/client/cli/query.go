package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tharsis/ethermint/x/ops/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	opsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the ops module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	opsQueryCmd.AddCommand(
		GetRecordByIDCmd(),
		ListRecorsCmd(),
		GetRecordCounterCmd(),
	)

	return opsQueryCmd
}

// GetRecordByIdCmd gets record by Id
func GetRecordByIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Get record by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			id := args[0]

			res, err := queryClient.GetRecord(cmd.Context(), &types.QueryRecordRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// ListRecorsCmd lists all existing records
func ListRecorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List all records",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ListRecords(cmd.Context(), &types.QueryAllRecordsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetRecordCounterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "counter",
		Args:  cobra.NoArgs,
		Short: "get record counter",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetRecordCounter(cmd.Context(), &types.QueryRecordCounterRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
