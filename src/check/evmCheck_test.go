package check

import (
	"supervisor/src/config"
	"testing"
)

//func TestEvmSyncCheck(t *testing.T) {
//	log.InitLog()
//	rpc := "https://eth-goerli.public.blastapi.io"
//	timeout := 60
//
//	expected := false
//	result, err := EvmSyncCheck(rpc, timeout)
//	if err != nil {
//		t.Errorf("unexpected error: %v ,maybe rpc %s is invalid", err, rpc)
//	}
//	if result != expected {
//		t.Errorf("incorrect result: expected %v, got %v", expected, result)
//	}
//	//If we have unhealthy not synced rpc to test,will be better
//	//rpcNotsynced:=""
//	//notsyncedexpected := false
//	//notsyncedresult, err := EvmSyncCheck(rpcNotsynced, timeout)
//	//if err != nil {
//	//	t.Errorf("unexpected error: %v,maybe rpc %s is invalid", err, rpc)
//	//}
//	//if notsyncedresult != notsyncedexpected {
//	//	t.Errorf("incorrect result: expected %v, got %v", notsyncedexpected, notsyncedresult)
//	//}
//}

func TestEvmCheckLatestBlock(t *testing.T) {
	rpc := "https://eth-goerli.public.blastapi.io"
	standardrpc := "https://goerli.infura.io/v3/28da0689af874bca9a175c66dcd1e151"
	maxBehindBlocks := 5
	timeout := 60

	expected := true
	result, err := EvmCheckLatestBlock(rpc, standardrpc, maxBehindBlocks, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpc, standardrpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
	//If we have unhealthy not synced rpc to test,will be better
	//rpcNotsynced:=""
	//notsyncedexpected := false
	//notsyncedresult, err := EvmCheckLatestBlock(rpcNotsynced, standardrpc,10,timeout)
	//if err != nil {
	// t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpc, standardrpc)
	//}
	//if notsyncedresult != notsyncedexpected {
	//	t.Errorf("incorrect result: expected %v, got %v", notsyncedexpected, notsyncedresult)
	//}

}

func TestEvmCheckAll(t *testing.T) {
	var node config.Node
	node.Rpc = "https://eth-goerli.public.blastapi.io"
	node.Standardrpc = "https://goerli.infura.io/v3/28da0689af874bca9a175c66dcd1e151"
	node.Chain = "goerli"
	node.Name = "test"
	node.Nodetype = "evm"

	expected := true
	result, err := EvmCheckAll(node, 10, 60)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s  is invalid", err, node.Rpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}

	//if we have unhealthy rpc to test,will be better
	//unhealthyRpc:=""
	//node.Rpc = unhealthyRpc
	//unhealthyExpected := false
	//unhealthyResult, err := EvmCheckAll(node,10,60)
	//if err != nil {
	//	t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, node.Rpc, node.Standardrpc)
	//}
	//if unhealthyResult != unhealthyExpected {
	//	t.Errorf("incorrect result: expected %v, got %v", unhealthyExpected, unhealthyResult)
	//}

}
