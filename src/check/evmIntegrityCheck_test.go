package check

import (
	"testing"
)

func TestEvmGetBalance(t *testing.T) {
	rpc := "https://goerli.infura.io/v3/28da0689af874bca9a175c66dcd1e151"
	address := "0x7FDCDe37BdCa782092c22E8adc683c9D61e508eE"
	timeout := 60
	expectedBalance := "0x0"

	balance, err := EvmGetBalance(rpc, address, timeout)
	if err != nil {
		t.Errorf("EvmGetBalance returned error: %v ,maybe rpc %s has problem", err, rpc)
	}

	if balance != expectedBalance {
		t.Errorf("Expected balance %s but got %s", expectedBalance, balance)
	}
}

func TestEvmGetBalanceCheck(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	address := "0x0000000000000000000000000000000000000000"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	result, err := EvmGetBalanceCheck(rpc, standardrpc, address, timeout)
	if err != nil {
		t.Errorf("EvmGetBalanceCheck returned error: %v ， maybe rpc %s or standard rpc %s have error", err, rpc, standardrpc)
	}

	if !result {
		t.Error("Expected EvmGetBalanceCheck to return true but got false")
	}

	//If we have unhealthy rpc to test,will be better
	//unhealthyrpc := ""
	//unhealthyResult, err := EvmGetBalanceCheck(unhealthyrpc, standardrpc, address, timeout)
	//if err != nil {
	//	t.Errorf("EvmGetBalanceCheck returned error: %v ， maybe rpc %s or standard rpc %s have error", err,rpc,standardrpc)
	//}
	//
	//if unhealthyResult {
	//	t.Error("Expected EvmGetBalanceCheck to return false but got true")
	//}
}

func TestEvmGetERC20Balance(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	address := "0x7FDCDe37BdCa782092c22E8adc683c9D61e508eE"
	ERC20address := "0x07865c6E87B9F70255377e024ace6630C1Eaa37F"
	timeout := 60
	expectedBalance := "0x0000000000000000000000000000000000000000000000000000000000000000"

	balance, err := EvmGetERC20Balance(rpc, address, ERC20address, timeout)
	if err != nil {
		t.Errorf("EvmGetERC20Balance returned error: %v ，maybe rpc %s has problem", err, rpc)
	}

	if balance != expectedBalance {
		t.Errorf("Expected balance %s but got %s", expectedBalance, balance)
	}
}

func TestEvmGetERC20BalanceCheck(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	address := "0x0000000000000000000000000000000000000000"
	ERC20address := "0x07865c6E87B9F70255377e024ace6630C1Eaa37F"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	result, err := EvmGetERC20BalanceCheck(rpc, standardrpc, address, ERC20address, timeout)
	if err != nil {
		t.Errorf("EvmGetERC20BalanceCheck returned error: %v ， maybe rpc %s or standard rpc %s have error", err, rpc, standardrpc)
	}

	if !result {
		t.Error("Expected EvmGetERC20BalanceCheck to return true but got false")
	}

	//If we have unhealthy rpc to test,will be better
	//unhealthyrpc := ""
	//unhealthyResult, err := EvmGetERC20BalanceCheck(unhealthyrpc, standardrpc, address, ERC20address, timeout)
	//if err != nil {
	//	t.Errorf("EvmGetERC20BalanceCheck returned error: %v ， maybe rpc %s or standard rpc %s have error", err,rpc,standardrpc)
	//}
	//
	//if unhealthyResult {
	//	t.Error("Expected EvmGetERC20BalanceCheck to return false but got true")
	//}
}

func TestEvmGetCode(t *testing.T) {
	rpc := "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	address := "0x07865c6E87B9F70255377e024ace6630C1Eaa37F"
	timeout := 60
	expectedCode := "0x60806040526004361061005a5760003560e01c80635c60da1b116100435780635c60da1b146101315780638f2839701461016f578063f851a440146101af5761005a565b80633659cfe6146100645780634f1ef286146100a4575b6100626101c4565b005b34801561007057600080fd5b506100626004803603602081101561008757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101de565b610062600480360360408110156100ba57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82351691908101906040810160208201356401000000008111156100f257600080fd5b82018360208201111561010457600080fd5b8035906020019184600183028401116401000000008311171561012657600080fd5b509092509050610232565b34801561013d57600080fd5b50610146610309565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017b57600080fd5b506100626004803603602081101561019257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610318565b3480156101bb57600080fd5b50610146610420565b6101cc610466565b6101dc6101d76104fa565b61051f565b565b6101e6610543565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102275761022281610568565b61022f565b61022f6101c4565b50565b61023a610543565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fc5761027683610568565b60003073ffffffffffffffffffffffffffffffffffffffff16348484604051808383808284376040519201945060009350909150508083038185875af1925050503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b50509050806102f657600080fd5b50610304565b6103046101c4565b505050565b60006103136104fa565b905090565b610320610543565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102275773ffffffffffffffffffffffffffffffffffffffff81166103bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806106966036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103e8610543565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301528051918290030190a1610222816105bd565b6000610313610543565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061045e57508115155b949350505050565b61046e610543565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156104f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806106646032913960400191505060405180910390fd5b6101dc6101dc565b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c35490565b3660008037600080366000845af43d6000803e80801561053e573d6000f35b3d6000fd5b7f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b5490565b610571816105e1565b6040805173ffffffffffffffffffffffffffffffffffffffff8316815290517fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b9181900360200190a150565b7f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b55565b6105ea8161042a565b61063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806106cc603b913960400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c35556fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a2646970667358221220119e941d353783c92238fbc4e38a3a0327e471d10cff47c0a5066819d4a4195664736f6c634300060c0033"

	code, err := EvmGetCode(rpc, address, timeout)
	if err != nil {
		t.Errorf("EvmGetCode returned error: %v， maybe rpc %s has problem", err, rpc)
	}

	if code != expectedCode {
		t.Errorf("Expected code %s but got %s", expectedCode, code)
	}
}

func TestEvmGetCodeCheck(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	contractAddress := "0x07865c6E87B9F70255377e024ace6630C1Eaa37F"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	expected := true
	result, err := EvmGetCodeCheck(rpc, standardrpc, contractAddress, timeout)
	if err != nil {
		t.Errorf("EvmGetCodeCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	}
	if result != expected {
		t.Errorf("Expected result %v but got %v", expected, result)
	}

	//If we have unhealthy  rpc to test,will be better
	//rpcUnhealthy:=""
	//expectedUnhealthy:=false
	//resultUnhealthy, err := EvmGetCodeCheck(rpcUnhealthy, standardrpc, contractAddress, timeout)
	//if err != nil {
	//	t.Errorf("EvmGetCodeCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	//}
	//if resultUnhealthy != expectedUnhealthy {
	//	t.Errorf("Expected result %v but got %v", expectedUnhealthy, resultUnhealthy)
	//}
}

func TestEvmGetTransactionCount(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	address := "0x0000000000000000000000000000000000000000"
	timeout := 60
	expected := "0x0"
	count, err := EvmGetTransactionCount(rpc, address, timeout)
	if err != nil {
		t.Errorf("EvmGetTransactionCount returned error: %v， maybe rpc %s has problem", err, rpc)
	}

	if count != expected {
		t.Errorf("Expected count %s but got %s", expected, count)
	}
}

func TestEvmGetTransactionCountCheck(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	address := "0x0000000000000000000000000000000000000000"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	expected := true
	result, err := EvmGetTransactionCountCheck(rpc, standardrpc, address, timeout)
	if err != nil {
		t.Errorf("EvmGetTransactionCountCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	}
	if result != expected {
		t.Errorf("Expected result %v but got %v", expected, result)
	}

	//If we have unhealthy  rpc to test,will be better
	//rpcUnhealthy:=""
	//expectedUnhealthy:=false
	//resultUnhealthy, err := EvmGetTransactionCountCheck(rpcUnhealthy, standardrpc, address, timeout)
	//if err != nil {
	//	t.Errorf("EvmGetTransactionCountCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	//}
	//if resultUnhealthy != expectedUnhealthy {
	//	t.Errorf("Expected result %v but got %v", expectedUnhealthy, resultUnhealthy)
	//}
}

func TestEvmWeb3_sha3(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	data := "aaa"
	timeout := 60
	expected := "0xb9a5dc0048db9a7d13548781df3cd4b2334606391f75f40c14225a92f4cb3537"
	result, err := EvmWeb3_sha3(rpc, data, timeout)
	if err != nil {
		t.Errorf("EvmWeb3_sha3 returned error: %v， maybe rpc %s has problem", err, rpc)
	}
	if result != expected {
		t.Errorf("Expected result %v but got %v", expected, result)
	}
}

func TestEvmWeb3_sha3Check(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	expected := true
	result, err := EvmWeb3_sha3Check(rpc, standardrpc, timeout)
	if err != nil {
		t.Errorf("EvmWeb3_sha3Check returned error: %v， maybe rpc %s has problem", err, rpc)
	}
	if result != expected {
		t.Errorf("Expected result %v but got %v", expected, result)
	}

	//If we have unhealthy rpc to test,will be better
	//rpcUnhealthy:=""
	//expectedUnhealthy:=false
	//resultUnhealthy, err := EvmWeb3_sha3Check(rpcUnhealthy, standardrpc, timeout)
	//if err != nil {
	//	t.Errorf("EvmWeb3_sha3Check returned error: %v， maybe rpc %s has problem", err, rpc)
	//}
	//if resultUnhealthy != expectedUnhealthy {
	//	t.Errorf("Expected result %v but got %v", expectedUnhealthy, resultUnhealthy)
	//}
}

func TestEvmGetTransactionByHash(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	hash := "0x0c63c504901f6170f399920b4937c46abc3d072706d39530a2c661ec1ead0700"
	timeout := 60

	_, err := EvmGetTransactionByHash(rpc, hash, timeout)
	if err != nil {
		t.Errorf("EvmGetBlockByHash returned error: %v， maybe rpc %s has problem", err, rpc)
	}
}

func TestEvmGetTransactionByHashCheck(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	hash := "0x0c63c504901f6170f399920b4937c46abc3d072706d39530a2c661ec1ead0700"
	timeout := 60
	standardrpc := "https://eth-goerli.public.blastapi.io"
	expected := true
	result, err := EvmGetTransactionByHashCheck(rpc, standardrpc, hash, timeout)
	if err != nil {
		t.Errorf("EvmGetTransactionByHashCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	}
	if result != expected {
		t.Errorf("Expected result %v but got %v", expected, result)
	}

	//If we have unhealthy rpc to test,will be better
	//rpcUnhealthy:=""
	//expectedUnhealthy:=false
	//resultUnhealthy, err := EvmGetTransactionByHashCheck(rpcUnhealthy, standardrpc, hash, timeout)
	//if err != nil {
	//	t.Errorf("EvmGetTransactionByHashCheck returned error: %v， maybe rpc %s has problem", err, rpc)
	//}
	//if resultUnhealthy != expectedUnhealthy {
	//	t.Errorf("Expected result %v but got %v", expectedUnhealthy, resultUnhealthy)
	//}
}

func TestEvmIntegrityCheck(t *testing.T) {

}
