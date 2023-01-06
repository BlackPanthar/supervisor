package check

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"supervisor/src/config"
	"supervisor/src/log"
	"time"
)

// CosmosQueryValidators Query RPC+ "/validators" to get the data of validators
func CosmosQueryValidators(rpc string, timeout int) (string, error) {
	req, err := http.NewRequest("GET", rpc+"/validators", nil)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryValidators check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryValidators check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryValidators check ,read response error: %v", err)
	}
	// Parse response
	return string(body), nil
}

// QueryValidatorsCheck Compare the data of validators from RPC and standard RPC, if they are different, return false
func QueryValidatorsCheck(rpc string, standardrpc string, timeout int) (bool, error) {
	log.Log.Info("start to perform queryValidators check", zap.String("rpc", rpc))
	validators, err := CosmosQueryValidators(rpc, 10)
	if err != nil {
		return false, err
	}
	standardvalidators, err := CosmosQueryValidators(standardrpc, timeout)
	if err != nil {
		return false, err
	}
	if validators == standardvalidators {
		return true, nil
	} else {
		log.Log.Error("rpc queryValidators check failed ", zap.String("rpc", rpc))
		return false, nil
	}
}

// CosmosQueryBlock Query RPC+ +"/block?height="+blockHeight to get the data of block
func CosmosQueryBlock(rpc string, blockHeight int, timeout int) (string, error) {
	blockHeightStr := strconv.Itoa(blockHeight)
	req, err := http.NewRequest("GET", rpc+"/block?height="+blockHeightStr, nil)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryBlock check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryBlock check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform queryBlock check ,read response error: %v", err)
	}
	return string(body), nil

}

// CosmosQueryBlockCheck Compare the data of block(The Latest block - 5) from RPC and standard RPC, if they are different, return false
func CosmosQueryBlockCheck(rpc string, standardrpc string, timeout int) (bool, error) {
	log.Log.Info("start to perform queryBlock check", zap.String("rpc", rpc))
	blockHeight, err := CosmosGetLatestBlock(standardrpc, timeout)
	if err != nil {
		log.Log.Error("failed to perform queryBlock check ,get latest block from standard rpc failed", zap.Error(err), zap.String("standardrpc", standardrpc))
		return true, nil
	}
	block, err := CosmosQueryBlock(rpc, blockHeight-5, timeout)
	if err != nil {
		return false, err
	}
	standardblock, err := CosmosQueryBlock(standardrpc, blockHeight-5, timeout)
	if err != nil {
		return false, err
	}
	if block == standardblock {
		return true, nil
	} else {
		log.Log.Error("rpc queryBlock check failed ", zap.String("rpc", rpc))
		return false, nil
	}
}

// CosmosIntegrityCheck Perform all integrity checks, if any check fails, return false
func CosmosIntegrityCheck(cconfig config.CosmosIntegrityCheckConfig, node config.Node, timeout int) (bool, error) {
	log.Log.Info("start to perform integrity check", zap.String("rpc", node.Rpc))
	if cconfig.QueryValidators {
		validatorsCheck, err := QueryValidatorsCheck(node.Rpc, node.Standardrpc, timeout)
		if err != nil {
			return false, err
		}
		if validatorsCheck == false {
			return false, nil
		}
	}
	if cconfig.QueryBlock {
		queryBlockCheck, err := CosmosQueryBlockCheck(node.Rpc, node.Standardrpc, timeout)
		if err != nil {
			return false, err
		}
		if queryBlockCheck == false {
			return false, nil
		}
	}
	log.Log.Info("integrity check passed", zap.String("rpc", node.Rpc))
	return true, nil
}
