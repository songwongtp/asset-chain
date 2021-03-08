package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

const (
	// FlagDenom is used to specify a certain asset denom
	FlagDenom = "denom"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group asset queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		GetAssetsCmd(),
	)

	return cmd
}

// GetAssetsCmd returns the query commands for retrieving asset info(s)
func GetAssetsCmd() *cobra.Command {
	cmd := &cobra.Command{
			Use: "assets",
			Short: "Query for asset info(s)",
			Long: strings.TrimSpace(
				fmt.Sprintf(`Query the information of all assets or of a specific denomination.
Examples:
	$ %s query %s assets
	$ %s query %s assets --denom=[denom]
`,
					version.AppName, types.ModuleName, version.AppName, types.ModuleName,
				),
			),
			Args: cobra.ExactArgs(0),
			RunE: func(cmd *cobra.Command, args []string) error {
				clientCtx, err := client.GetClientQueryContext(cmd)
				if err != nil {
					return err
				}

				denom, err := cmd.Flags().GetString(FlagDenom)
				if err != nil {
					return err
				}

				queryClient := types.NewQueryClient(clientCtx)

				if denom == "" {
					params := types.NewQueryAllAssetInfosRequest()
					res, err := queryClient.AllAssetInfos(context.Background(), params)
					if err != nil {
						return err
					}
					return clientCtx.PrintProto(res)
				}

				params := types.NewQueryAssetInfoRequest(denom)
				res, err := queryClient.AssetInfo(context.Background(), params)
				if err != nil {
					return err
				}
				return clientCtx.PrintProto(res)
			},
	}

	cmd.Flags().String(FlagDenom, "", "The Specific asset denomination to query for")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}