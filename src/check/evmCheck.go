package check

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"supervisor/src/config"
	"supervisor/src/log"
	"time"
)

// EvmSyncCheck use ethclient.SyncProgress to check if the node is syncing, if the node is syncing, return true, else return false
// Since some RFC do not support this method , stop using it

//func EvmSyncCheck(rpc string, timeout int) (bool, error) {
//	log.Log.Info("checking RPC sync status", zap.String("rpc", rpc))
//	client, err := ethclient.Dial(rpc)
//	if err != nil {
//		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
//	defer cancel()
//	sync, err := client.SyncProgress(ctx)
//	if err != nil {
//		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
//	}
//
//	if sync == nil {
//		log.Log.Info("rpc is fully synced", zap.String("rpc", rpc))
//		return false, nil
//	}
//	log.Log.Error("rpc is still syncing", zap.String("rpc", rpc))
//	return true, nil
//}

// EvmCheckLatestBlock use ethclient.BlockNumber to get the latest block height, if the latest block height is behind the standard node over maxBehindBlocks , return false, else return true
func EvmCheckLatestBlock(rpc string, standardrpc string, maxBehindBlocks int, timeout int) (bool, error) {
	log.Log.Info("checking RPC latest block", zap.String("rpc", rpc))
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return false, fmt.Errorf("rpc check lastest block failed, err: %v", err)
	}

	standardClient, err := ethclient.Dial(standardrpc)
	if err != nil {
		log.Log.Error("rpc check latest block failed, because check standard rpc failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("rpc check lastest block failed, err: %v", err)
	}
	standardctx, standardcancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer standardcancel()
	standardBlockNumber, err := standardClient.BlockNumber(standardctx)
	if err != nil {
		log.Log.Error("rpc check latest block failed, because check standard rpc failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if blockNumber+uint64(maxBehindBlocks) < standardBlockNumber {
		log.Log.Error("rpc is behind standard rpc too much", zap.String("rpc", rpc), zap.Uint64("RPCBlockNumber", blockNumber), zap.Uint64("standardRpcBlockNumber", standardBlockNumber))
		return false, nil
	}
	log.Log.Info("rpc is fully synced comparing standard Rpc", zap.String("rpc", rpc))
	return true, nil
}

// EvmCheckAll Perform all evm checks, if any check fails, return false, else return true
func EvmCheckAll(node config.Node, maxbehindblock int, timeout int) (bool, error) {
	connection, err := CheckConnection(node.Rpc, timeout)
	if err != nil {
		return false, err
	}
	if !connection {
		return false, nil
	}
	//sync, err := EvmSyncCheck(node.Rpc, timeout)
	//if err != nil {
	//	return false, err
	//}
	//if sync {
	//	return false, nil
	//}
	latestBlock, err := EvmCheckLatestBlock(node.Rpc, node.Standardrpc, maxbehindblock, timeout)
	if err != nil {
		return false, err
	}
	if !latestBlock {
		return false, nil
	}
	return true, nil
}
