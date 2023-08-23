package cli

import (
	"context"

	"belkyc/x/belkyc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListValidatorKYC() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-validator-kyc",
		Short: "list all validatorKYC",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorKYCRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorKYCAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowValidatorKYC() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-kyc [address]",
		Short: "shows a validatorKYC",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetValidatorKYCRequest{
				Address: argAddress,
			}

			res, err := queryClient.ValidatorKYC(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
