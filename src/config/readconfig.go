package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"supervisor/src/log"
)

type Config struct {
	Checkconfig    checkconfig `json:"checkconfig"`
	Integritycheck bool        `json:"Integritycheck"`
	Nodes          []Node      `json:"nodes"`
}

type checkconfig struct {
	Granularity       int
	Connect_timeout   int
	Max_behind_blocks int
}

type Node struct {
	Nodetype    string
	Chain       string
	Name        string
	Rpc         string
	Standardrpc string
}

// GetConfig read config.toml and return in config struct
func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Log.Error("read config file failed", zap.Error(err))
		panic(err)
	}
	var c Config
	c.Checkconfig.Granularity = viper.GetInt("check.granularity")
	c.Checkconfig.Connect_timeout = viper.GetInt("check.connect_timeout")
	c.Checkconfig.Max_behind_blocks = viper.GetInt("check.max_behind_blocks")
	c.Integritycheck = viper.GetBool("integritycheck.enabled")
	c.Nodes = make([]Node, 0)
	for k, v := range viper.GetStringMap("cosmossdk") {
		for k1, v1 := range v.(map[string]interface{}) {
			if k1 == "standard_rpc" {
				//should not be empty
				if v1.(string) == "" {
					log.Log.Error(" standard rpc is empty", zap.String("chain", k))
					panic(fmt.Errorf("chain %s standard rpc is empty", k))
				}
			}
			if k1 == "node" {
				for _, v2 := range v1.([]interface{}) {
					var n Node
					n.Nodetype = "cosmossdk"
					n.Chain = k
					//should not be empty
					if v2.([]interface{})[0].(string) == "" || v2.([]interface{})[1].(string) == "" {
						log.Log.Error("node name or rpc is empty", zap.String("chain", k))
						panic(fmt.Errorf("chain %s node name or rpc is empty", k))
					}
					n.Name = v2.([]interface{})[0].(string)
					n.Rpc = v2.([]interface{})[1].(string)
					n.Standardrpc = viper.GetString("cosmossdk." + k + ".standard_rpc")
					c.Nodes = append(c.Nodes, n)
				}
			}
		}
	}
	//evm
	for k, v := range viper.GetStringMap("evm") {
		for k1, v1 := range v.(map[string]interface{}) {
			if k1 == "standard_rpc" {
				//should not be empty
				if v1.(string) == "" {
					log.Log.Error(" standard rpc is empty", zap.String("chain", k))
					panic(fmt.Errorf("chain %s standard rpc is empty", k))
				}
			}
			if k1 == "node" {
				for _, v2 := range v1.([]interface{}) {
					var evmnode Node
					evmnode.Standardrpc = viper.GetString("evm." + k + ".standard_rpc")
					evmnode.Nodetype = "evm"
					evmnode.Chain = k
					//should not be empty
					if v2.([]interface{})[0].(string) == "" || v2.([]interface{})[1].(string) == "" {
						log.Log.Error("node name or rpc is empty", zap.String("chain", k))
						panic(fmt.Errorf("chain %s node name or rpc is empty", k))
					}
					evmnode.Name = v2.([]interface{})[0].(string)
					evmnode.Rpc = v2.([]interface{})[1].(string)
					c.Nodes = append(c.Nodes, evmnode)
				}
			}
		}
	}
	return c
}

type EvmIntegrityCheckConfig struct {
	Chain                                      string
	Getbalance                                 bool   `toml:"getbalance"`
	GetbalanceDefaultaddress                   string `toml:"getbalanceDefaultaddress"`
	GetERC20Balance                            bool   `toml:"getERC20balance"`
	DefaultERC20Address                        string `toml:"defaultERC20address"`
	GetERC20BalanceDefaultaddress              string `toml:"getERC20balanceDefaultaddress"`
	GetCode                                    bool   `toml:"getCode"`
	GetCodeDefaultaddress                      string `toml:"getCodeDefaultaddress"`
	GetTransactionCount                        bool   `toml:"getTransactionCount"`
	GetTransactionCountDefaultaddress          string `toml:"getTransactionCountDefaultaddress"`
	Web3Sha3                                   bool   `toml:"web3_sha3"`
	GetTransactionByHash                       bool   `toml:"getTransactionByHash"`
	GetTransactionByHashDefaulttransactionhash string `toml:"getTransactionByHashDefaulttransactionhash"`
}

// GetEvmIntegrityCheckConfig Read evm integrity check config file and return in evmIntegrityCheckConfig array
func GetEvmIntegrityCheckConfig() []EvmIntegrityCheckConfig {
	viper.SetConfigName("integritycheck")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Log.Error("read integrity check config file failed", zap.Error(err))
		panic(err)
	}
	var c []EvmIntegrityCheckConfig
	for k, v := range viper.GetStringMap("evm") {
		var singleconfig EvmIntegrityCheckConfig
		singleconfig.Chain = k
		for k1, v1 := range v.(map[string]interface{}) {
			if k1 == "getbalance" {
				singleconfig.Getbalance = v1.(bool)
			}
			if k1 == "getbalancedefaultaddress" {
				singleconfig.GetbalanceDefaultaddress = v1.(string)
			}
			if k1 == "geterc20balance" {
				singleconfig.GetERC20Balance = v1.(bool)
			}
			if k1 == "defaulterc20address" {
				singleconfig.DefaultERC20Address = v1.(string)
			}
			if k1 == "geterc20balancedefaultaddress" {
				singleconfig.GetERC20BalanceDefaultaddress = v1.(string)
			}
			if k1 == "getcode" {
				singleconfig.GetCode = v1.(bool)
			}
			if k1 == "getcodedefaultaddress" {
				singleconfig.GetCodeDefaultaddress = v1.(string)
			}
			if k1 == "gettransactioncount" {
				singleconfig.GetTransactionCount = v1.(bool)
			}
			if k1 == "gettransactioncountdefaultaddress" {
				singleconfig.GetTransactionCountDefaultaddress = v1.(string)
			}
			if k1 == "web3_sha3" {
				singleconfig.Web3Sha3 = v1.(bool)
			}
			if k1 == "gettransactioncount" {
				singleconfig.GetTransactionByHash = v1.(bool)
			}
			if k1 == "gettransactionbyhashdefaulttransactionhash" {
				singleconfig.GetTransactionByHashDefaulttransactionhash = v1.(string)
			}
		}
		c = append(c, singleconfig)
	}
	return c
}

type CosmosIntegrityCheckConfig struct {
	Chain           string
	QueryValidators bool `toml:"queryValidators"`
	QueryBlock      bool `toml:"queryBlock"`
}

// GetCosmosIntegrityCheckConfig Read cosmos integrity check config file and return in cosmosIntegrityCheckConfig array
func GetCosmosIntegrityCheckConfig() []CosmosIntegrityCheckConfig {
	viper.SetConfigName("integritycheck")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Log.Error("read integrity check config file failed", zap.Error(err))
		panic(err)
	}
	var c []CosmosIntegrityCheckConfig
	for k, v := range viper.GetStringMap("cosmossdk") {
		var singleconfig CosmosIntegrityCheckConfig
		singleconfig.Chain = k
		for k1, v1 := range v.(map[string]interface{}) {
			if k1 == "queryValidators" {
				singleconfig.QueryValidators = v1.(bool)
			}
			if k1 == "queryBlock" {
				singleconfig.QueryBlock = v1.(bool)

			}
		}
		c = append(c, singleconfig)
	}
	return c
}
