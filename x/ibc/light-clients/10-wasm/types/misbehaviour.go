package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

var (
	_ exported.Misbehaviour = &Misbehaviour{}
)

func (m *Misbehaviour) ClientType() string {
	return m.Header1.ClientType()
}

func (m *Misbehaviour) GetClientID() string {
	return m.ClientId
}

func (m *Misbehaviour) ValidateBasic() error {
	const ValidateBasicQuery = "misbehaviourvalidatebasic"
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

func (m *Misbehaviour) GetHeight() exported.Height {
	return m.Header1.GetHeight()
}