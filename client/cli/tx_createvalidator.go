package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contohdam",
		Short: "Broadcast message contohdam",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr := clientCtx.GetFromAddress()
			pkStr, err := cmd.Flags().GetString(FlagPubKey)

			var pk cryptotypes.PubKey
			if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &pk); err != nil {
				return err
			}

			moniker, _ := cmd.Flags().GetString(FlagMoniker)
			identity, _ := cmd.Flags().GetString(FlagIdentity)
			website, _ := cmd.Flags().GetString(FlagWebsite)
			security, _ := cmd.Flags().GetString(FlagSecurityContact)
			details, _ := cmd.Flags().GetString(FlagDetails)
			description := types.NewDescription(
				moniker,
				identity,
				website,
				security,
				details,
			)

			msg, err := types.NewMsgCreateValidator(
				sdk.ValAddress(valAddr),
				pk,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
