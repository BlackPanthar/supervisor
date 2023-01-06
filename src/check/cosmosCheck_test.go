package check

import (
	"supervisor/src/config"
	"supervisor/src/log"
	"testing"
)

func TestCosmosSyncCheck(t *testing.T) {
	log.InitLog()

	rpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60

	expected := false
	result, err := CosmosSyncCheck(rpc, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s is invalid", err, rpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
	//If we have unhealthy not synced rpc to test,will be better
	//rpcNotsynced:=""
	//notsyncedexpected := false
	//notsyncedresult, err :=CosmosSyncCheck(rpcNotsynced, timeout)
	//if err != nil {
	//	t.Errorf("unexpected error: %v,maybe rpc %s is invalid", err, rpc)
	//}
	//if notsyncedresult != notsyncedexpected {
	//	t.Errorf("incorrect result: expected %v, got %v", notsyncedexpected, notsyncedresult)
	//}
}

func TestCosmosGetLatestBlock(t *testing.T) {
	log.InitLog()
	rpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60

	_, err := CosmosGetLatestBlock(rpc, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCosmosCheckLatestBlock(t *testing.T) {
	log.InitLog()
	rpc := "https://rpc.stride.silentvalidator.com"
	standardrpc := "https://rpc.stride.silentvalidator.com"
	maxBehindBlocks := 5
	timeout := 60

	expected := true
	result, err := CosmosCheckLatestBlock(rpc, standardrpc, maxBehindBlocks, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpc, standardrpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
	//If we have unhealthy not synced rpc to test,will be better
	//rpcNotsynced:=""
	//notsyncedexpected := false
	//notsyncedresult, err := CosmosCheckLatestBlock(rpcNotsynced, standardrpc, maxBehindBlocks, timeout)
	//if err != nil {
	//	t.Errorf("unexpected error: %v,maybe rpc %s or standard rpc %s is invalid", err, rpcNotsynced,standardrpc)
	//}
	//if notsyncedresult != notsyncedexpected {
	//	t.Errorf("incorrect result: expected %v, got %v", notsyncedexpected, notsyncedresult)
	//}
}

func TestCosmosCheckAll(t *testing.T) {
	log.InitLog()
	var node config.Node
	node.Rpc = "https://rpc.stride.silentvalidator.com"
	node.Standardrpc = "https://rpc.stride.silentvalidator.com"
	node.Chain = "stride"
	node.Nodetype = "cosmossdk"
	node.Name = "test"

	maxBehindBlocks := 10
	timeout := 60
	result, err := CosmosCheckAll(node, maxBehindBlocks, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != true {
		t.Errorf("incorrect result: expected %v, got %v, or rpc %s has problem", true, result, node.Rpc)
	}

	//If we have unhealthy not synced rpc to test,will be better
	//node.Rpc=""
	//nodeUnhealthyexpected := false
	//nodeUnhealthyresult, err := CosmosCheckAll(node, maxBehindBlocks, timeout)
	//if err != nil {
	//	t.Errorf("unexpected error: %v", err)
	//}
	//if nodeUnhealthyresult != nodeUnhealthyexpected {
	//	t.Errorf("incorrect result: expected %v, got %v, or rpc %s has problem", nodeUnhealthyexpected, nodeUnhealthyresult, node.Rpc)

}
