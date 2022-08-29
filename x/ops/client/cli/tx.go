package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tharsis/ethermint/server/flags"
	"github.com/tharsis/ethermint/x/ops/types"
)

// NewTxCmd returns a root CLI command handler for all x/ops transaction commands.
func NewTxCmd() *cobra.Command {
	opsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "ops transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	opsTxCmd.AddCommand(
		NewCreateRecordCmd(),
		UpdateRecordCmd(),
		DeleteRecordCmd(),
	)

	return opsTxCmd
}

// NewCreateRecordCmd is the CLI command for creating a record.
func NewCreateRecordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name] [age]",
		Short: "Create record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			age, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("cannot convert age to uint64: %w", err)
			}

			msg := types.NewMsgCreateRecord(clientCtx.GetFromAddress(), name, uint64(age))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	_, _ = flags.AddTxFlags(cmd)

	return cmd
}

// UpdateRecordCmd is the CLI command for updating a record.
func UpdateRecordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id] [name] [age]",
		Short: "Update record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			id := args[0]
			name := args[1]
			age, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("cannot convert age to uint64: %w", err)
			}

			msg := types.NewMsgUpdateRecord(clientCtx.GetFromAddress(), id, name, uint64(age))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	_, _ = flags.AddTxFlags(cmd)
	return cmd
}

// DeleteRecordCmd is the CLI command for deleting a record.
func DeleteRecordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			id := args[0]
			msg := types.NewMsgDeleteRecord(clientCtx.GetFromAddress(), id)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	_, _ = flags.AddTxFlags(cmd)
	return cmd
}
