package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"

	wasm "github.com/CosmWasm/wasmvm"

	"github.com/cosmos/cosmos-sdk/x/ibc/core/28-wasm/types"

    "crypto/sha256"
	"encoding/hex"
)

func generateWASMCodeID(code []byte) string {
	hash := sha256.Sum256(code)
	return hex.EncodeToString(hash[:])
}

// Keeper will have a reference to Wasmer with it's own data directory.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler
	wasmer *wasm.VM
	wasmValidator *WASMValidator
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, validationConfig *WASMValidationConfig) Keeper {
	// TODO: Make this configurable
	vm, err := wasm.NewVM("wasm_data", "staking", 8, true, 8)
	if err != nil {
		panic(err)
	}

	wasmValidator, err := NewWASMValidator(validationConfig, func () (*wasm.VM, error) {
		return wasm.NewVM("wasm_test_data", "staking", 8, true, 8)
	})
	if err != nil {
		panic(err)
	}

	return Keeper{
		wasmer: vm,
		cdc: cdc,
		storeKey: key,
		wasmValidator: wasmValidator,
	}
}

func (k Keeper) PushNewWASMCode(ctx sdk.Context, clientType string, code []byte) (string, error) {
	store := ctx.KVStore(k.storeKey)
	codeId := generateWASMCodeID(code)

	latestVersionKey := host.LatestWASMCode(clientType)
	codekey := host.WASMCode(clientType, codeId)
	entryKey := host.WASMCodeEntry(clientType, codeId)

	if isValidWASMCode, err := k.wasmValidator.validateWASMCode(code); err != nil {
		return "", fmt.Errorf("unable to validate wasm code, error: %s", err)
	} else {
		if !isValidWASMCode {
			return "", fmt.Errorf("invalid wasm code")
		}
	}

	latestVersionCodeId := store.Get(latestVersionKey)

	// More careful management of doubly linked list can lift this constraint
	// But we do not see any significant advantage of it.
	if store.Has(entryKey) {
		return "", fmt.Errorf("wasm code already exists")
	} else {
		codeEntry := types.WasmCodeEntry{
			PreviousCodeId: string(latestVersionCodeId),
			NextCodeId: "",
		}

		previousVersionEntryKey := host.WASMCodeEntry(clientType, string(latestVersionCodeId))
		previousVersionEntryBz := store.Get(previousVersionEntryKey)
		if len(previousVersionEntryBz) != 0 {
			var previousEntry types.WasmCodeEntry
			k.cdc.MustUnmarshalBinaryBare(previousVersionEntryBz, &previousEntry)
			previousEntry.NextCodeId = codeId
			store.Set(previousVersionEntryKey, k.cdc.MustMarshalBinaryBare(&previousEntry))
		}

		store.Set(entryKey, k.cdc.MustMarshalBinaryBare(&codeEntry))
		store.Set(latestVersionKey, []byte(codeId))
		store.Set(codekey, code)
	}
	return codeId, nil
}

