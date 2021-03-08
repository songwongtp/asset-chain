package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		GetBuyAssetCmd(),
		GetSellAssetCmd(),
	)

	return cmd
}

// GetBuyAssetCmd returns the message command for buying asset
func GetBuyAssetCmd() *cobra.Command {
	cmd := &cobra.Command{
			Use: "buy [buyer] [denom] [amount]",
			Short: `buy asset of the denom type for the given amount. Note, the'--from' flag is
ignored as it is implied from [buyer].`,
			Args: cobra.ExactArgs(3),
			RunE: func(cmd *cobra.Command, args []string) error {
				cmd.Flags().Set(flags.FlagFrom, args[0])

				clientCtx, err := client.GetClientTxContext(cmd)
				if err != nil {
					return err
				}

				denom := args[1]

				amount, err := strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}
				
				msg := types.NewMsgBuyAsset(clientCtx.GetFromAddress().String(), denom, amount)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}

				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetSellAssetCmd returns the message command for selling asset
func GetSellAssetCmd() *cobra.Command {
	cmd := &cobra.Command{
			Use: "sell [seller] [denom] [amount]",
			Short: `sell asset of the denom type for the given amount. Note, the'--from' flag is
ignored as it is implied from [buyer].`,
			Args: cobra.ExactArgs(3),
			RunE: func(cmd *cobra.Command, args []string) error {
				cmd.Flags().Set(flags.FlagFrom, args[0])

				clientCtx, err := client.GetClientTxContext(cmd)
				if err != nil {
					return err
				}

				denom := args[1]

				amount, err := strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				msg := types.NewMsgSellAsset(clientCtx.GetFromAddress().String(), denom, amount)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}

				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}