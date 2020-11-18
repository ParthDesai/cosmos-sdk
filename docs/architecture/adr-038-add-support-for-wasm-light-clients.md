# ADR 38: Add support for Wasm based light client
## Changelog
- {date}: {changelog}
## Status
{DRAFT | PROPOSED} Not Implemented
> Please have a look at the [PROCESS](./PROCESS.md#adr-status) page.
> Use DRAFT if the ADR is in a draft stage (draft PR) or PROPOSED if it's in review.
## Abstract
> "If you can't explain it simply, you don't understand it well enough." Provide a simplified and layman-accessible explanation of the ADR.
> A short (~200 word) description of the issue being addressed.

Currently in the SDK, light clients are part of Cosmos SDK codebase. What this means that, anytime a blockchain built upon Cosmos SDK need to add
support for new light client, Cosmos SDK need to be modified and all validator nodes of that blockchain need to be updated with custom version of
Cosmos SDK. 

To remedy these shortcomings, we are proposing a WASM VM to host light client bytecode, which allows easier upgrading of existing light clients
as well as adding support for new light clients without requiring a code release and corresponding hard-fork event.
## Context
> This section describes the forces at play, including technological, political, social, and project local. These forces are probably in tension, and should be called out as such. The language in this section is value-neutral. It is simply describing facts. It should clearly explain the problem and motivation that the proposal aims to resolve.
> {context body}
>
Currently in the SDK, light clients are defined as part of the codebase and are implemented as submodules under `/x/ibc/light-clients/`. To add support for new light client or
update an existing light client in the event of security issue or consensus update, we need to modify the codebase and integrate it in *numerous* places. In addition, individual blockchains built upon Cosmos SDK
need to update all their validator nodes to latest version to add support for this light client. This entire process is tedious and time consuming.
In the event a blockchain wants to add support for niche light client, it need to fork Cosmos SDK and modify codebase to add support for it. This creates overhead of
maintaining that fork against mainstream Cosmos SDK release.

We are proposing simplifying this workflow by integrating a WASM light client module which makes adding support for new light client a simple transaction.
The light client bytecode, written in Wasm-compilable Rust runs inside a Wasmer VM. The Wasm light client submodule exposes a proxy light client interface that routes incoming messages to the appropriate handler function, inside the Wasm VM for execution.
## Decision
> This section describes our response to these forces. It is stated in full sentences, with active voice. "We will ..."
> {decision body}

We decided to use WASM light client module as a generic light client which will interface with the actual light client uploaded as WASM bytecode.
This will require changing client selection method to allow any client if the client type has prefix of `wasm/`.
```go
// IsAllowedClient checks if the given client type is registered on the allowlist.
func (p Params) IsAllowedClient(clientType string) bool {
	for _, allowedClient := range p.AllowedClients {
		if allowedClient == clientType {
			return true
		}
	}
	return false || isWASMClient(clientType)
}
```
Inside Wasm light client `ClientState`, appropriate Wasm bytecode will be executed depending upon `ClientType`.
```go
func (cs ClientState) Validate() error {
    wasmRegistry = getWASMRegistry()
	clientType := cs.ClientType()
    codeHandle := wasmRegistry.getCodeHandle(clientType)
    return codeHandle.validate(cs)
}
```
To upload new light client, user need to create a transaction with Wasm byte code and unique client id, which will be processed by IBC Wasm module.
```go
func (k Keeper) UploadLightClient (WasmCode: []byte, id: String, description: String) {
    wasmRegistry = getWASMRegistry()
    assert(!wasmRegistry.exists(id))
    assert(wasmRegistry.validateAndStoreCode(id, description, wasmCode, false))
}
```
As name implies, Wasm registry is a registry which stores set of Wasm client code indexed by its unique id and allows client code to retrieve latest code uploaded.
## Consequences
> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.
### Positive
- Adding support for new light client or upgrading existing light client is way easier than before and only requires single transaction.
- Improves maintainability of Cosmos SDK, since no change in codebase is required to support new client or upgrade it.
### Negative
- Light clients need to be written in subset of rust which could compile in Wasm.
- Introspecting light client code is difficult as only compiled bytecode exists in the blockchain.
