package check

import (
	"supervisor/src/config"
	"supervisor/src/log"
	"testing"
)

func TestCosmosQueryValidators(t *testing.T) {
	log.InitLog()
	rpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60

	_, err := CosmosQueryValidators(rpc, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestQueryValidatorsCheck(t *testing.T) {
	log.InitLog()
	rpc := "https://rpc.stride.silentvalidator.com"
	standardrpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60
	expected := true
	result, err := QueryValidatorsCheck(rpc, standardrpc, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpc, standardrpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}

	//if we have unhealthy rpc to test,will be better
	//rpcNotsynced:=""
	//notsyncedexpected := false
	//notsyncedresult, err := QueryValidatorsCheck(rpcNotsynced, standardrpc, timeout)
	//if err != nil {
	//	t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpcNotsynced, standardrpc)
	//}
	//if notsyncedresult != notsyncedexpected {
	//	t.Errorf("incorrect result: expected %v, got %v", notsyncedexpected, notsyncedresult)
	//}

}

func TestCosmosQueryBlock(t *testing.T) {
	rpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60
	//get the latest block
	latestBlock, err := CosmosGetLatestBlock(rpc, timeout)
	if err != nil {
		t.Errorf("Get latest block failed: %v", err)
	}
	_, err = CosmosQueryBlock(rpc, latestBlock-5, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCosmosQueryBlockCheck(t *testing.T) {
	log.InitLog()
	rpc := "https://rpc.stride.silentvalidator.com"
	standardrpc := "https://rpc.stride.silentvalidator.com"
	timeout := 60
	expected := true
	result, err := CosmosQueryBlockCheck(rpc, standardrpc, timeout)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s is invalid", err, rpc)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}

	//if we have unhealthy rpc to test,will be better
	//rpcUnhealthy:=""
	//rpcUnhealthyExpected := false
	//rpcUnhealthyResult, err := CosmosQueryBlockCheck(rpcUnhealthy, standardrpc, timeout)
	//if err != nil {
	//	t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpcUnhealthyResult, standardrpc)
	//}
	//if rpcUnhealthyResult != rpcUnhealthyExpected {
	//	t.Errorf("incorrect result: expected %v, got %v", rpcUnhealthyExpected, rpcUnhealthyResult)
	//}
}

func TestCosmosIntegrityCheck(t *testing.T) {
	log.InitLog()
	var c config.CosmosIntegrityCheckConfig
	c.Chain = "stride"
	c.QueryBlock = true
	c.QueryValidators = true
	expected := true
	var node config.Node
	node.Rpc = "https://rpc.stride.silentvalidator.com"
	node.Standardrpc = "https://rpc.stride.silentvalidator.com"
	node.Chain = "stride"
	node.Nodetype = "cosmossdk"
	node.Name = "test"
	result, err := CosmosIntegrityCheck(c, node, 60)
	if err != nil {
		t.Errorf("unexpected error: %v ,maybe rpc %s is invalid", err, node.Chain)
	}
	if result != expected {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}

	//if we have unhealthy rpc to test,will be better
	//unhealthyRpc:=""
	//node.Rpc=unhealthyRpc
	//rpcUnhealthyExpected := false
	//rpcUnhealthyResult, err := CosmosIntegrityCheck(c,node , 60)
	//if err != nil {
	//	t.Errorf("unexpected error: %v ,maybe rpc %s or standard rpc %s is invalid", err, rpcUnhealthyResult, standardrpc)
	//}
	//if rpcUnhealthyResult != rpcUnhealthyExpected {
	//	t.Errorf("incorrect result: expected %v, got %v", rpcUnhealthyExpected, rpcUnhealthyResult)
	//}

}
