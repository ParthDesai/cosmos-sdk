package keeper

import cosmwasm "github.com/CosmWasm/wasmvm"

// Basic validation config can be extended to add other configuration later
type WASMValidationConfig struct {
	MaxSizeAllowed int
}

func NewWASMValidator(config *WASMValidationConfig, vm *cosmwasm.VM) (*WASMValidator, error) {
	return &WASMValidator{
		config: config,
		testVm: vm,
	}, nil
}

type WASMValidator struct {
	testVm *cosmwasm.VM
	config *WASMValidationConfig
}

func (v *WASMValidator) validateWASMCode(code []byte) bool {
	if len(code) > v.config.MaxSizeAllowed {
		return false
	}

	_, err := v.testVm.Create(code)
	if err != nil {
		return false
	}

	// Validation start

	// Validation ends

	v.testVm.Cleanup()
	return true
}


