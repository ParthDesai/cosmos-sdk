package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/light-clients/10-wasm/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func NewCreateClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [CodeID in hex] [path/to/consensus_state.bin] [path/to/client_state.bin]",
		Short: "create new wasm client",
		Long: "Create a new wasm IBC client",
		Example: fmt.Sprintf("%s tx ibc %s create [path/to/consensus_state.json] [path/to/client_state.json]", version.AppName, types.SubModuleName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			clientStateBytes, err := ioutil.ReadFile(args[1])
			if err != nil {
				return errors.Wrap(err, "error reading client state from file")
			}

			clientState := types.ClientState{}
			if err := json.Unmarshal(clientStateBytes, &clientState); err != nil {
				return errors.Wrap(err, "error unmarshalling client state")
			}

			consensusStateBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return errors.Wrap(err, "error reading consensus state from file")
			}

			consensusState := types.ConsensusState{}
			if err := json.Unmarshal(consensusStateBytes, &consensusState); err != nil {
				return errors.Wrap(err, "error unmarshalling consensus state")
			}

			if bytes.Compare(clientState.CodeId, consensusState.CodeId) != 0 {
				return fmt.Errorf("CodeId mismatch between client state and consensus state")
			}

			msg, err := clienttypes.NewMsgCreateClient(
				&clientState, &consensusState, clientCtx.GetFromAddress(),
			)
			if err != nil {
				return errors.Wrap(err, "error composing MsgCreateClient")
			}

			if err := msg.ValidateBasic(); err != nil {
				return errors.Wrap(err, "error validating MsgCreateClient")
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
			"$ %s tx ibc %s update [client-id] [path/to/header.json] --from node0 --home ../node0/<app>cli --chain-id $CID",
			version.AppName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientID := args[0]

			headerBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return errors.Wrap(err, "error reading header from file")
			}

			header := types.Header{}
			if err := json.Unmarshal(headerBytes, &header); err != nil {
				return errors.Wrap(err, "error unmarshalling header")
			}

			msg, err := clienttypes.NewMsgUpdateClient(clientID, &header, clientCtx.GetFromAddress())
			if err != nil {
				return errors.Wrap(err, "error composing MsgUpdateClient")
			}

			if err := msg.ValidateBasic(); err != nil {
				return errors.Wrap(err, "error validating MsgUpdateClient")
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSubmitMisbehaviourCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misbehaviour [client-Id] [codeId in hex] [path/to/header1.json] [path/to/header2.json]",
		Short: "submit a client misbehaviour",
		Long:  "submit a client misbehaviour to invalidate to invalidate previous state roots and prevent future updates",
		Example: fmt.Sprintf(
			"$ %s tx ibc %s misbehaviour [client-Id] [path/to/header1.json] [path/to/header2.json] --from node0 --home ../node0/<app>cli --chain-id $CID",
			version.AppName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientID := args[0]

			header1Bytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return errors.Wrap(err, "error reading header1 from file")
			}
			header1 := types.Header{}
			if err := json.Unmarshal(header1Bytes, &header1); err != nil {
				return errors.Wrap(err, "error unmarshalling header1")
			}

			header2Bytes, err := ioutil.ReadFile(args[3])
			if err != nil {
				return errors.Wrap(err, "error reading header2 from file")
			}
			header2 := types.Header{}
			if err := json.Unmarshal(header2Bytes, &header2); err != nil {
				return errors.Wrap(err, "error unmarshalling header2")
			}

			if bytes.Compare(header1.CodeId, header2.CodeId) != 0 {
				return fmt.Errorf("CodeId mismatch between two headers")
			}


			misbehaviour := types.Misbehaviour{
				CodeId:     header1.CodeId,
				ClientId:   clientID,
				Header1:    &header1,
				Header2:    &header2,
			}

			msg, err := clienttypes.NewMsgSubmitMisbehaviour(misbehaviour.ClientId, &misbehaviour, clientCtx.GetFromAddress())
			if err != nil {
				return errors.Wrap(err, "error composing MsgSubmitMisbehaviour")
			}

			if err := msg.ValidateBasic(); err != nil {
				return errors.Wrap(err, "error validating MsgSubmitMisbehaviour")
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
