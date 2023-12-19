// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PushOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/PushOracle.sol:PushOracle\",\"label\":\"data\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_uint256,t_uint256)\"}],\"types\":{\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PushOracleStorageLayout = new(solc.StorageLayout)

var PushOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c80631ab06ee514610046578063e591b2821461005b578063f0ba8440146100a0575b600080fd5b610059610054366004610187565b6100ce565b005b61007673deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100c06100ae3660046101a9565b60006020819052908152604090205481565b604051908152602001610097565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f507573684f7261636c653a206f6e6c7920746865206465706f7369746f72206160448201527f63636f756e740000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b60009182526020829052604090912055565b6000806040838503121561019a57600080fd5b50508035926020909101359150565b6000602082840312156101bb57600080fd5b503591905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(PushOracleStorageLayoutJSON), PushOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PushOracle"] = PushOracleStorageLayout
	deployedBytecodes["PushOracle"] = PushOracleDeployedBin
	immutableReferences["PushOracle"] = false
}
