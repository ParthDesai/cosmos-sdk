package cli

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/light-clients/10-wasm/types"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func NewCreateClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [CodeID in hex] [path/to/consensus_state.bin] [path/to/client_state.bin]",
		Short: "create new wasm client",
		Long: "Create a new wasm IBC client",
		Example: fmt.Sprintf("%s tx ibc %s create [CodeID in hex] [path/to/consensus_state.bin] [path/to/client_state.bin]", version.AppName, types.SubModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := hex.DecodeString(args[0])
			if err != nil {
				// TODO: Handle error
			}

			clientStateBytes, err := ioutil.ReadFile(args[1])
			if err != nil {
				// TODO: Handle error
			}

			consensusStateBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				// TODO: Handle error
			}

			clientState := types.ClientState{
				Data: clientStateBytes,
				CodeId: codeID,
			}

			consensusState := types.ConsensusState{
				Data: consensusStateBytes,
				CodeId: codeID,
			}

			msg, err := clienttypes.NewMsgCreateClient(
				&clientState, &consensusState, clientCtx.GetFromAddress(),
			)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewUpdateClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [client-id] [Code Id in hex] [path/to/header.bin]",
		Short: "update existing client with a header",
		Long:  "update existing wasm client with a header",
		Example: fmt.Sprintf(
			"$ %s tx ibc %s update [client-id] [path/to/header.bin] --from node0 --home ../node0/<app>cli --chain-id $CID",
			version.AppName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientID := args[0]

			codeID, err := hex.DecodeString(args[1])
			if err != nil {
				// TODO: Handle error
			}

			headerBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				// TODO: Handle error
			}

			header := types.Header{
				Data: headerBytes,
				CodeId: codeID,
			}

			msg, err := clienttypes.NewMsgUpdateClient(clientID, &header, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSubmitMisbehaviourCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misbehaviour [client-Id] [codeId in hex] [path/to/header1.bin] [path/to/header2.bin]",
		Short: "submit a client misbehaviour",
		Long:  "submit a client misbehaviour to invalidate to invalidate previous state roots and prevent future updates",
		Example: fmt.Sprintf(
			"$ %s tx ibc %s misbehaviour [client-Id] [codeId in hex] [path/to/header1.bin] [path/to/header2.bin] --from node0 --home ../node0/<app>cli --chain-id $CID",
			version.AppName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientID := args[0]

			codeID, err := hex.DecodeString(args[1])
			if err != nil {
				// TODO: Handle error
			}

			header1Bytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				// TODO: Handle error
			}

			header2Bytes, err := ioutil.ReadFile(args[3])
			if err != nil {
				// TODO: Handle error
			}

			header1 := types.Header{
				Data:   header1Bytes,
				CodeId: codeID,
			}

			header2 := types.Header{
				Data:   header2Bytes,
				CodeId: codeID,
			}

			misbehaviour := types.Misbehaviour{
				CodeId:     codeID,
				InstanceId: 0,
				ClientId:   clientID,
				Header1:    &header1,
				Header2:    &header2,
			}

			msg, err := clienttypes.NewMsgSubmitMisbehaviour(misbehaviour.ClientId, &misbehaviour, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
