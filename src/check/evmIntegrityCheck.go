package check

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"supervisor/src/config"
	"supervisor/src/log"
	"time"
)

// EvmGetBalance use eth_getBalance method to get balance
func EvmGetBalance(rpc string, address string, timeout int) (string, error) {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBalance",
		"params":  []string{address, "latest"},
		"id":      1,
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to perform getBalance check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform getBalance check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform getBalance check ,read response error: %v", err)
	}
	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to perform getBalance check ,parse response error: %v", err)
	}

	if result["result"] != nil {
		return result["result"].(string), nil
	}
	return "", fmt.Errorf("failed to perform getBalance check, result is nil")
}

// EvmGetBalanceCheck use eth_getBalance method to get balance from rpc and standard rpc, if they are not equal, return false
func EvmGetBalanceCheck(rpc string, standardrpc string, address string, timeout int) (bool, error) {
	log.Log.Info("start to perform getBalance check", zap.String("rpc", rpc))
	balance, err := EvmGetBalance(rpc, address, timeout)
	if err != nil {
		return false, err
	}
	standardbalance, err := EvmGetBalance(standardrpc, address, timeout)
	if err != nil {
		log.Log.Error("failed to perform getBalance check because use standard rpc to get balance failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if balance == standardbalance {
		return true, nil
	}
	log.Log.Error("rpc getBalance check failed", zap.String("rpc", rpc))
	return false, nil

}

// EvmGetCodeCheck use eth_call method to get ERC20 token balance
func EvmGetERC20Balance(rpc string, address string, ERC20address string, timeout int) (string, error) {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_call",
		"params": []interface{}{
			map[string]interface{}{
				"to":   ERC20address,
				"data": "0x70a08231000000000000000000000000" + address[2:],
			},
			"latest",
		},
		"id": 1,
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to perform get ERC20 balance check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform get ERC20 balance check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform get ERC20 balance check ,read response error: %v", err)
	}
	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to perform get ERC20 balance check ,parse response error: %v", err)
	}
	if result["result"] != nil {
		return result["result"].(string), nil
	}
	return "", fmt.Errorf("failed to perform get ERC20 balance check, result is nil")
}

// EvmGetERC20BalanceCheck use eth_call method to get ERC20 token balance from rpc and standard rpc, if they are not equal, return false
func EvmGetERC20BalanceCheck(rpc string, standardrpc string, address string, ERC20address string, timeout int) (bool, error) {
	log.Log.Info("start to perform get ERC20 balance check", zap.String("rpc", rpc))
	balance, err := EvmGetERC20Balance(rpc, address, ERC20address, timeout)
	if err != nil {
		return false, err
	}
	standardbalance, err := EvmGetERC20Balance(standardrpc, address, ERC20address, timeout)
	if err != nil {
		log.Log.Error("failed to perform get ERC20 balance check because use standard rpc to get balance failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if balance == standardbalance {
		return true, nil
	}
	log.Log.Error("rpc get ERC20 balance check failed", zap.String("rpc", rpc))
	return false, nil
}

// EvmGetCode use eth_getCode method to get code
func EvmGetCode(rpc, address string, timeout int) (string, error) {
	// Set up request
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getCode",
		"params":  []string{address, "latest"},
		"id":      1,
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to perform Getcode check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Set up context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform Getcode check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform Getcode check ,read response error: %v", err)
	}
	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to perform Getcode check ,parse response error: %v", err)
	}
	if result["result"] != nil {
		return result["result"].(string), nil
	}
	return "", fmt.Errorf("failed to perform Getcode check, result is nil")
}

// EvmGetCodeCheck use eth_getCode method to get code from rpc and standard rpc, if they are not equal, return false
func EvmGetCodeCheck(rpc string, standardrpc string, address string, timeout int) (bool, error) {
	log.Log.Info("start to perform get code check", zap.String("rpc", rpc))
	code, err := EvmGetCode(rpc, address, timeout)
	if err != nil {
		return false, err
	}
	standardcode, err := EvmGetCode(standardrpc, address, timeout)
	if err != nil {
		log.Log.Error("failed to perform Getcode check because use standard rpc to perform Getcode failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if code == standardcode {
		return true, nil
	}
	log.Log.Error("rpc Getcode check failed", zap.String("rpc", rpc))
	return false, nil
}

// EvmGetTransactionCount use eth_getTransactionCount method to get transaction count
func EvmGetTransactionCount(rpc, address string, timeout int) (string, error) {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionCount",
		"params":  []string{address, "latest"},
		"id":      1,
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to perform getTransactionCount check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform getTransactionCount check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform getTransactionCount check ,read response error: %v", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to perform getTransactionCount check ,parse response error: %v", err)
	}
	if result["result"] != nil {
		return result["result"].(string), nil
	}
	return "", fmt.Errorf("failed to perform getTransactionCount check, result is nil")
}

// EvmGetTransactionCountCheck use eth_getTransactionCount method to get transaction count from rpc and standard rpc, if they are not equal, return false
func EvmGetTransactionCountCheck(rpc string, standardrpc string, address string, timeout int) (bool, error) {
	log.Log.Info("start to perform get getTransactionCount check", zap.String("rpc", rpc))
	transactionCount, err := EvmGetTransactionCount(rpc, address, timeout)
	if err != nil {
		return false, err
	}
	standardtransactionCount, err := EvmGetTransactionCount(standardrpc, address, timeout)
	if err != nil {
		log.Log.Error("failed to perform getTransactionCount check because use standard rpc to perform getTransactionCount failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if transactionCount == standardtransactionCount {
		return true, nil
	}
	log.Log.Error("rpc getTransactionCount check failed", zap.String("rpc", rpc))
	return false, nil
}

// EvmWeb3_sha3 use web3_sha3 method to get sha3
func EvmWeb3_sha3(rpc, data string, timeout int) (string, error) {
	// Set up request
	//convert data to hex
	data = "0x" + hex.EncodeToString([]byte(data))
	reqData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "web3_sha3",
		"params":  []string{data},
		"id":      1,
	}
	payload, _ := json.Marshal(reqData)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to perform web3_sha3 check ,make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform web3_sha3 check ,http request failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to perform web3_sha3 check ,read response error: %v", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to perform web3_sha3 check ,parse response error: %v", err)
	}
	if result["result"] != nil {
		return result["result"].(string), nil
	}
	return "", fmt.Errorf("failed to perform web3_sha3 check, result is nil")

}

// EvmWeb3_sha3Check use web3_sha3 method to get sha3 of random string from rpc and standard rpc, if they are not equal, return false
func EvmWeb3_sha3Check(rpc string, standardrpc string, timeout int) (bool, error) {
	log.Log.Info("start to perform web3_sha3 check", zap.String("rpc", rpc))
	//get random string as data
	data, _ := uuid.NewRandom()
	sha3, err := EvmWeb3_sha3(rpc, data.String(), timeout)
	if err != nil {
		return false, err
	}
	standardsha3, err := EvmWeb3_sha3(standardrpc, data.String(), timeout)
	if err != nil {
		log.Log.Error("failed to perform web3_sha3 check because use standard rpc to perform web3_sha3 failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if sha3 == standardsha3 {
		return true, nil
	}
	log.Log.Error("rpc web3_sha3 check failed", zap.String("rpc", rpc))
	return false, nil
}

// EvmGetTransactionByHash use eth_getTransactionByHash method to get transaction data by hash
func EvmGetTransactionByHash(rpc, hash string, timeout int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		"params":  []string{hash},
		"id":      1,
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rpc, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to perform getTransactionByHash check, make new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform getTransactionByHash check, http request failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to perform getTransactionByHash check, read response error: %v", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to perform getTransactionByHash check, parse response error: %v", err)
	}
	if result["result"] != nil {
		return result["result"].(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("failed to perform getTransactionByHash check, result is nil")
}

// EvmGetTransactionByHashCheck use eth_getTransactionByHash method to get transaction data by hash from rpc and standard rpc, if they are not equal, return false
func EvmGetTransactionByHashCheck(rpc string, standardrpc string, hash string, timeout int) (bool, error) {
	log.Log.Info("start to perform getTransactionByHash check", zap.String("rpc", rpc))
	tx, err := EvmGetTransactionByHash(rpc, hash, timeout)
	if err != nil {
		return false, err
	}
	standardtx, err := EvmGetTransactionByHash(standardrpc, hash, timeout)
	if err != nil {
		log.Log.Error("failed to perform getTransactionByHash check because use standard rpc to perform getTransactionByHash failed", zap.String("rpc", rpc), zap.String("standard rpc", standardrpc), zap.Error(err))
		return true, nil
	}
	if reflect.DeepEqual(tx, standardtx) {
		return true, nil
	}
	log.Log.Error("rpc getTransactionByHash check failed", zap.String("rpc", rpc))
	return false, nil
}

// IsRightHash Check if the hash format is correct
func IsRightHash(hash string) bool {
	var blockHashRegexp = regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	return blockHashRegexp.MatchString(hash)
}

// EvmIntegrityCheck Perform all evm integrity check, if any check failed, return false. If integrity config have problem, panic
func EvmIntegrityCheck(integrityCheckConfig config.EvmIntegrityCheckConfig, node config.Node, timeout int) (bool, error) {
	log.Log.Info("start to perform integrity check", zap.String("rpc", node.Rpc))
	if integrityCheckConfig.Getbalance == true {
		//ensure the address  format is right
		if common.IsHexAddress(integrityCheckConfig.GetbalanceDefaultaddress) == false {
			panic("getbalanceDefaultaddress in integritycheck.toml is not right")
		}

		getbalance, err := EvmGetBalanceCheck(node.Rpc, node.Standardrpc, integrityCheckConfig.GetbalanceDefaultaddress, timeout)
		if err != nil {
			return false, err
		}
		if getbalance == false {
			return false, nil
		}

	}
	if integrityCheckConfig.GetERC20Balance == true {
		if !common.IsHexAddress(integrityCheckConfig.DefaultERC20Address) {
			panic("the defaultERC20address in integritycheck.toml is not a valid ERC20 address")
		}
		if !common.IsHexAddress(integrityCheckConfig.GetERC20BalanceDefaultaddress) {
			panic("the getERC20balanceDefaultaddress in integritycheck.toml is not a valid ERC20 address")
		}
		getERC20Balance, err := EvmGetERC20BalanceCheck(node.Rpc, node.Standardrpc, integrityCheckConfig.GetERC20BalanceDefaultaddress, integrityCheckConfig.DefaultERC20Address, timeout)
		if err != nil {
			return false, err
		}
		if getERC20Balance == false {
			return false, nil
		}
	}
	if integrityCheckConfig.GetCode == true {
		if !common.IsHexAddress(integrityCheckConfig.GetCodeDefaultaddress) {
			panic("the getCodeDefaultaddress in integritycheck.toml is not a valid address")
		}
		getCode, err := EvmGetCodeCheck(node.Rpc, node.Standardrpc, integrityCheckConfig.GetCodeDefaultaddress, timeout)
		if err != nil {
			return false, err
		}
		if getCode == false {
			return false, nil
		}
	}
	if integrityCheckConfig.GetTransactionCount == true {
		if !common.IsHexAddress(integrityCheckConfig.GetTransactionCountDefaultaddress) {
			panic("the getTransactionCountDefaultaddress in integritycheck.toml is not a valid address")
		}
		getTransactionCount, err := EvmGetTransactionCountCheck(node.Rpc, node.Standardrpc, integrityCheckConfig.GetTransactionCountDefaultaddress, timeout)
		if err != nil {
			return false, err
		}
		if getTransactionCount == false {
			return false, nil
		}
	}
	if integrityCheckConfig.Web3Sha3 == true {
		web3Sha3, err := EvmWeb3_sha3Check(node.Rpc, node.Standardrpc, timeout)
		if err != nil {
			return false, err
		}
		if web3Sha3 == false {
			return false, nil
		}
	}
	if integrityCheckConfig.GetTransactionByHash == true {
		getTransactionByHash, err := EvmGetTransactionByHashCheck(node.Rpc, node.Standardrpc, integrityCheckConfig.GetTransactionByHashDefaulttransactionhash, timeout)
		if err != nil {
			return false, err
		}
		if getTransactionByHash == false {
			return false, nil
		}
	}
	log.Log.Info("integrity check passed", zap.String("rpc", node.Rpc))
	return true, nil
}
