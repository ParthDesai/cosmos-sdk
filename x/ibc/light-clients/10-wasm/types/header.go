package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

var _ exported.Header = (*Header)(nil)

func (m *Header) ClientType() string {
	const ClientTypeQuery = "headerclienttype"
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

func (m *Header) GetHeight() exported.Height {
	const GetHeightQuery = "headergetheight"
	payload := make(map[string]map[string]interface{})
	payload[GetHeightQuery] = make(map[string]interface{})
	inner := payload[GetHeightQuery]
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

	return output.Height

}

func (m *Header) ValidateBasic() error {
	const ValidateBasicQuery = "headervalidatebasic"
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
