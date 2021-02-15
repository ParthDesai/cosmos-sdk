package types

import (
	"encoding/json"
	"fmt"
	"github.com/CosmWasm/wasmvm/api"
	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)


func NewClientState() {

}

/**
Following are functions that modifies state, so should be part of handle call
 */


func (c *ClientState) Initialize(context sdk.Context, marshaler codec.BinaryMarshaler, store sdk.KVStore, state exported.ConsensusState) error {
	payload := make(map[string]interface{})
	payload["self"] = c
	payload["consensus_state"] = state

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	if !output.Result.IsValid {
		return fmt.Errorf("%s error ocurred while initializing client", output.Result.ErrorMsg)
	}
	// We might have modified client state, so should store updated version of it
	*c = *output.Self
	return nil
}

func (c *ClientState) CheckHeaderAndUpdateState(context sdk.Context, marshaler codec.BinaryMarshaler, store sdk.KVStore, header exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	const CheckHeaderAndUpdateState = "checkandupdateclientstate"
	payload := make(map[string]map[string]interface{})
	payload[CheckHeaderAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckHeaderAndUpdateState]
	inner["self"] = c
	inner["header"] = header

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	return output.NewClientState, output.NewConsensusState, nil
}

func (c *ClientState) CheckMisbehaviourAndUpdateState(context sdk.Context, marshaler codec.BinaryMarshaler, store sdk.KVStore, misbehaviour exported.Misbehaviour) (exported.ClientState, error) {
	const CheckMisbehaviourAndUpdateState = "checkmisbehaviourandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[CheckMisbehaviourAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckMisbehaviourAndUpdateState]
	inner["self"] = c
	inner["misbehaviour"] = misbehaviour

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}

	return output.NewClientState, nil
}

func (c *ClientState) CheckProposedHeaderAndUpdateState(context sdk.Context, marshaler codec.BinaryMarshaler, store sdk.KVStore, header exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	const CheckProposedHeaderAndUpdateState = "checkproposedheaderandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[CheckProposedHeaderAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckProposedHeaderAndUpdateState]
	inner["self"] = c
	inner["header"] = header

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	return output.NewClientState, output.NewConsensusState, nil
}

func (c *ClientState) VerifyUpgradeAndUpdateState(ctx sdk.Context, cdc codec.BinaryMarshaler, store sdk.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient, proofUpgradeConsState []byte) (exported.ClientState, exported.ConsensusState, error) {
	const VerifyUpgradeAndUpdateState = "verifyupgradeandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyUpgradeAndUpdateState] = make(map[string]interface{})
	inner := payload[VerifyUpgradeAndUpdateState]
	inner["self"] = c
	inner["new_client"] = newClient
	inner["new_consensus_state"] = newConsState
	inner["client_upgrade_proof"] = proofUpgradeClient
	inner["consensus_state_upgrade_proof"] = proofUpgradeConsState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	out, err := callContract(c.CodeId, ctx, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	return output.NewClientState, output.NewConsensusState, nil
}

func (c *ClientState) ZeroCustomFields() exported.ClientState {
	const ZeroCustomFields = "zerocustomfields"
	payload := make(map[string]map[string]interface{})
	payload[ZeroCustomFields] = make(map[string]interface{})
	inner := payload[ZeroCustomFields]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}

	gasMeter := sdk.NewGasMeter(0)
	out, err := callContractWithEnvAndMeter(c.CodeId, nil, &FailKVStore{}, api.MockEnv(), gasMeter, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	return output.Self
}


/**
Following functions only queries the state so should be part of query call
 */

func (c *ClientState) ClientType() string {
	const ClientTypeQuery = "clienttype"
	payload := make(map[string]map[string]interface{})
	payload[ClientTypeQuery] = make(map[string]interface{})
	inner := payload[ClientTypeQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(c.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.ClientType
}

func (c *ClientState) ExportMetadata(store sdk.KVStore) []exported.GenesisMetadata {
	const ExportMetadataQuery = "exportmetadata"
	payload := make(map[string]map[string]interface{})
	payload[ExportMetadataQuery] = make(map[string]interface{})
	inner := payload[ExportMetadataQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	genesisMetadata := make([]exported.GenesisMetadata, len(output.GenesisMetadata))
	for i, metadata := range output.GenesisMetadata {
		genesisMetadata[i] = metadata
	}
	return genesisMetadata
}

func (c *ClientState) GetLatestHeight() exported.Height {
	const GetLatestHeightQuery = "getlatestheight"
	payload := make(map[string]map[string]interface{})
	payload[GetLatestHeightQuery] = make(map[string]interface{})
	inner := payload[GetLatestHeightQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(c.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.Height
}

func (c *ClientState) IsFrozen() bool {
	const IsFrozenQuery = "isfrozen"
	payload := make(map[string]map[string]interface{})
	payload[IsFrozenQuery] = make(map[string]interface{})
	inner := payload[IsFrozenQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(c.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.IsFrozen
}

func (c *ClientState) GetFrozenHeight() exported.Height {
	const GetFrozenHeightQuery = "getfrozenheight"
	payload := make(map[string]map[string]interface{})
	payload[GetFrozenHeightQuery] = make(map[string]interface{})
	inner := payload[GetFrozenHeightQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(c.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.FrozenHeight
}

func (c *ClientState) Validate() error {
	if c.Data == nil || len(c.Data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}

	if c.CodeId == nil || len(c.CodeId) == 0 {
		return fmt.Errorf("codeid cannot be empty")
	}

	return nil
}

func (c *ClientState) GetProofSpecs() []*ics23.ProofSpec {
	const GetProofSpecsQuery = "getproofspecs"
	payload := make(map[string]map[string]interface{})
	payload[GetProofSpecsQuery] = make(map[string]interface{})
	inner := payload[GetProofSpecsQuery]
	inner["self"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(c.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.ProofSpecs
}

func (c *ClientState) VerifyClientState(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, prefix exported.Prefix, counterpartyClientIdentifier string, proof []byte, clientState exported.ClientState) error {
	const VerifyClientStateQuery = "verifyclientstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyClientStateQuery] = make(map[string]interface{})
	inner := payload[VerifyClientStateQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["counterparty_client_identifier"] = counterpartyClientIdentifier
	inner["proof"] = proof
	inner["counterparty_client_state"] = clientState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}

}

func (c *ClientState) VerifyClientConsensusState(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, counterpartyClientIdentifier string, consensusHeight exported.Height, prefix exported.Prefix, proof []byte, consensusState exported.ConsensusState) error {
	const VerifyClientConsensusStateQuery = "verifyclientconsensusstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyClientConsensusStateQuery] = make(map[string]interface{})
	inner := payload[VerifyClientConsensusStateQuery]
	inner["self"] = c
	inner["height"] = height
	inner["consensus_height"] = consensusHeight
	inner["commitment_prefix"] = prefix
	inner["counterparty_client_identifier"] = counterpartyClientIdentifier
	inner["proof"] = proof
	inner["counterparty_consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyConnectionState(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, prefix exported.Prefix, proof []byte, connectionID string, connectionEnd exported.ConnectionI) error {
	const VerifyConnectionStateQuery = "verifyconnectionstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyConnectionStateQuery] = make(map[string]interface{})
	inner := payload[VerifyConnectionStateQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["connection_id"] = connectionID
	inner["connection_end"] = connectionEnd

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyChannelState(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, prefix exported.Prefix, proof []byte, portID, channelID string, channel exported.ChannelI) error {
	const VerifyChannelStateQuery = "verifychannelstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyChannelStateQuery] = make(map[string]interface{})
	inner := payload[VerifyChannelStateQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["channel"] = channel

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyPacketCommitment(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64, commitmentBytes []byte) error {
	const VerifyPacketCommitmentQuery = "verifypacketcommitment"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketCommitmentQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketCommitmentQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence
	inner["commitment_bytes"] = commitmentBytes

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}

}

func (c *ClientState) VerifyPacketAcknowledgement(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64, acknowledgement []byte) error {
	const VerifyPacketAcknowledgementQuery = "verifypacketacknowledgement"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketAcknowledgementQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketAcknowledgementQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence
	inner["acknowledgement"] = acknowledgement

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyPacketReceiptAbsence(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64) error {
	const VerifyPacketReceiptAbsenceQuery = "verifypacketreceiptabsence"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketReceiptAbsenceQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketReceiptAbsenceQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyNextSequenceRecv(store sdk.KVStore, cdc codec.BinaryMarshaler, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, nextSequenceRecv uint64) error {
	const VerifyNextSequenceRecvQuery = "verifynextsequencerecv"
	payload := make(map[string]map[string]interface{})
	payload[VerifyNextSequenceRecvQuery] = make(map[string]interface{})
	inner := payload[VerifyNextSequenceRecvQuery]
	inner["self"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["next_sequence_recv"] = nextSequenceRecv

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating", output.Result.ErrorMsg)
	}
}
