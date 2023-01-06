package check

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"supervisor/src/config"
	"supervisor/src/log"
	"time"
)

type CosmosRpcStatusResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NodeInfo struct {
			ProtocolVersion struct {
				P2P   string `json:"p2p"`
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"protocol_version"`
			ID         string `json:"id"`
			ListenAddr string `json:"listen_addr"`
			Network    string `json:"network"`
			Version    string `json:"version"`
			Channels   string `json:"channels"`
			Moniker    string `json:"moniker"`
			Other      struct {
				TxIndex    string `json:"tx_index"`
				RPCAddress string `json:"rpc_address"`
			} `json:"other"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash     string    `json:"latest_block_hash"`
			LatestAppHash       string    `json:"latest_app_hash"`
			LatestBlockHeight   string    `json:"latest_block_height"`
			LatestBlockTime     time.Time `json:"latest_block_time"`
			EarliestBlockHash   string    `json:"earliest_block_hash"`
			EarliestAppHash     string    `json:"earliest_app_hash"`
			EarliestBlockHeight string    `json:"earliest_block_height"`
			EarliestBlockTime   time.Time `json:"earliest_block_time"`
			CatchingUp          bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

// CosmosSyncCheck Query RPC/status page to check if the node is synced, if not synced, return true
func CosmosSyncCheck(rpc string, timeout int) (bool, error) {
	log.Log.Info("checking RPC sync status", zap.String("rpc", rpc))

	req, err := http.NewRequest("GET", rpc+"/status", nil)
	if err != nil {
		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
	}
	var response CosmosRpcStatusResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return true, fmt.Errorf("rpc sync check failed, err: %v", err)
	}

	if response.Result.SyncInfo.CatchingUp {
		log.Log.Error("rpc is still syncing", zap.String("rpc", rpc))
		return true, nil
	} else {
		log.Log.Info("rpc is fully synced", zap.String("rpc", rpc))
		return false, nil
	}
}

// CosmosGetLatestBlock Query RPC/status page to get the latest block height
func CosmosGetLatestBlock(rpc string, timeout int) (int, error) {
	req, err := http.NewRequest("GET", rpc+"/status", nil)
	if err != nil {
		return 0, fmt.Errorf("rpc check lastest block failed, err: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("rpc check lastest block failed, err: %v ", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("rpc check lastest block failed, err: %v", err)
	}
	var response CosmosRpcStatusResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, fmt.Errorf("rpc check lastest block failed, err: %v", err)
	}
	rpcLatestBlock, _ := strconv.Atoi(response.Result.SyncInfo.LatestBlockHeight)
	return rpcLatestBlock, nil
}

// CosmosCheckLatestBlock  Compare the latest block height of the RPC with the latest block height of the standard rpc , if the difference is less than maxBehindBlocks , return true
func CosmosCheckLatestBlock(rpc string, standardrpc string, maxBehindBlocks int, timeout int) (bool, error) {
	log.Log.Info("checking RPC latest block", zap.String("rpc", rpc))
	rpcLastestBlock, err := CosmosGetLatestBlock(rpc, timeout)
	if err != nil {
		return false, err
	}
	standardRpcLatestBlock, err := CosmosGetLatestBlock(standardrpc, timeout)
	if err != nil {
		log.Log.Error("Failed to perform latest block check, because get standard rpc latest block failed ", zap.String("rpc", rpc), zap.String("standardrpc", standardrpc), zap.Int("timeout", timeout), zap.Error(err))
		return true, nil
	}

	if (rpcLastestBlock + maxBehindBlocks) < standardRpcLatestBlock {
		log.Log.Error("rpc is behind standard rpc too much", zap.String("rpc", rpc), zap.Int("rpcLatestBlock", rpcLastestBlock), zap.Int("standardRpcLatestBlock", standardRpcLatestBlock))
		return false, nil
	}
	log.Log.Info("rpc is fully synced comparing standard Rpc", zap.String("rpc", rpc))
	return true, nil
}

// CosmosCheckAll Perform all checks, if any check fails, return false
func CosmosCheckAll(node config.Node, maxbehindblock int, timeout int) (bool, error) {
	connection, err := CheckConnection(node.Rpc, timeout)
	if err != nil {
		return false, err
	}
	if !connection {
		return false, nil
	}
	sync, err := CosmosSyncCheck(node.Rpc, timeout)
	if err != nil {
		return false, err
	}
	if sync {
		return false, nil
	}
	latest, err := CosmosCheckLatestBlock(node.Rpc, node.Standardrpc, maxbehindblock, timeout)
	if err != nil {
		return false, err
	}
	if !latest {
		return false, nil
	}
	return true, nil
}
