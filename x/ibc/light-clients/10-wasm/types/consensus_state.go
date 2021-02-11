package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

func (m *ConsensusState) ClientType() string {
	const ClientTypeQuery = "consensusclienttype"
	payload := make(map[string]map[string]interface{})
	payload[ClientTypeQuery] = make(map[string]interface{})
	inner := payload[ClientTypeQuery]
	inner["self"] = m

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(m.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.ClientType
}

func (m *ConsensusState) GetRoot() exported.Root {
	const GetRootQuery = "consensusgetroot"
	payload := make(map[string]map[string]interface{})
	payload[GetRootQuery] = make(map[string]interface{})
	inner := payload[GetRootQuery]
	inner["self"] = m

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(m.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.Root
}

func (m *ConsensusState) GetTimestamp() uint64 {
	const GetTimeStampQuery = "consensusgettimestamp"
	payload := make(map[string]map[string]interface{})
	payload[GetTimeStampQuery] = make(map[string]interface{})
	inner := payload[GetTimeStampQuery]
	inner["self"] = m

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(m.CodeId, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	return output.Timestamp
}

func (m *ConsensusState) ValidateBasic() error {
	const ValidateBasicQuery = "consensusvalidatebasic"
	payload := make(map[string]map[string]interface{})
	payload[ValidateBasicQuery] = make(map[string]interface{})
	inner := payload[ValidateBasicQuery]
	inner["self"] = m

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContract(m.CodeId, encodedData)
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
