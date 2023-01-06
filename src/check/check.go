package check

import (
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"supervisor/src/config"
	"supervisor/src/log"
	"time"
)

// CheckConnection check TCP handshake, if the node is reachable, return true, else return false
func CheckConnection(rpc string, timeout int) (bool, error) {
	log.Log.Info("checking rpc TCP handshake", zap.String("rpc", rpc))
	//remove http:// or https://
	u, err := url.Parse(rpc)
	if err != nil {
		return false, fmt.Errorf("parse rpc %s url failed, err: %v", rpc, err)
	}
	var rpc1 string
	if u.Port() == "" {
		if u.Scheme == "http" {
			rpc1 = rpc + ":80"
		} else if u.Scheme == "https" {
			rpc1 = rpc + ":443"
		}
	}
	re := regexp.MustCompile(`^http(s)?://`)
	rpcwithoutprefix := re.ReplaceAllString(rpc1, "")
	//check TCP handshake
	conn, err := net.DialTimeout("tcp", rpcwithoutprefix, time.Duration(timeout)*time.Second)
	if err != nil {
		return false, fmt.Errorf(" rpc %s TCP handshake failed, err: %v", rpc, err)
	}
	defer conn.Close()
	log.Log.Info("rpc check TCP handshake success", zap.String("rpc", rpc))
	return true, nil
}

type NewPerformance struct {
	Nodetype    string
	Chain       string
	Name        string
	Rpc         string
	Performance bool
}

// CheckAll Performance all the check and return the result in newPerformance array
func CheckAll(Config config.Config) []NewPerformance {
	timeout := Config.Checkconfig.Connect_timeout
	maxBehindBlocks := Config.Checkconfig.Max_behind_blocks
	var performance []NewPerformance
	integrityCheck := Config.Integritycheck
	for _, v := range Config.Nodes {
		if v.Nodetype == "cosmossdk" {
			nodecheckResult, err := CosmosCheckAll(v, maxBehindBlocks, timeout)
			if err != nil {
				log.Log.Error("node check failed", zap.String("rpc", v.Rpc), zap.Error(err))
				performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
				continue
			}
			if !nodecheckResult {
				log.Log.Error("node check failed", zap.String("rpc", v.Rpc))
				performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
				continue

			}
			if integrityCheck {
				//flip a coin to  perform integrity check
				if rand.Intn(2) == 0 {
					// find the right integrity check config
					var cconfig config.CosmosIntegrityCheckConfig
					integrityCheckConfigs := config.GetCosmosIntegrityCheckConfig()
					for _, integrityCheckConfig := range integrityCheckConfigs {
						if integrityCheckConfig.Chain == v.Chain {
							cconfig = integrityCheckConfig
						}
					}
					integrityCheckResult, err := CosmosIntegrityCheck(cconfig, v, timeout)
					if err != nil {
						log.Log.Error("integrity check failed", zap.String("rpc", v.Rpc), zap.Error(err))
						performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
						continue
					}
					if !integrityCheckResult {
						log.Log.Error("integrity check failed", zap.String("rpc", v.Rpc))
						performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
						continue
					}
				}
			}
			log.Log.Info("node check succeed", zap.String("rpc", v.Rpc))
			performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: true})
		}
		if v.Nodetype == "evm" {
			nodecheckResult, err := EvmCheckAll(v, maxBehindBlocks, timeout)
			if err != nil {
				log.Log.Error("node check failed", zap.String("rpc", v.Rpc), zap.Error(err))
				performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
				continue
			}
			if !nodecheckResult {
				log.Log.Error("node check failed", zap.String("rpc", v.Rpc))
				performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
				continue
			}
			if integrityCheck {
				//flip a coin to  perform integrity check
				if rand.Intn(2) == 0 {
					// find the right integrity check config
					var econfig config.EvmIntegrityCheckConfig
					integrityCheckConfigs := config.GetEvmIntegrityCheckConfig()

					for _, integrityCheckConfig := range integrityCheckConfigs {
						if integrityCheckConfig.Chain == v.Chain {
							econfig = integrityCheckConfig
						}
					}
					integrityCheckResult, err := EvmIntegrityCheck(econfig, v, timeout)
					if err != nil {
						log.Log.Error("integrity check failed", zap.String("rpc", v.Rpc), zap.Error(err))
						performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
						continue
					}
					if !integrityCheckResult {
						log.Log.Error("integrity check failed", zap.String("rpc", v.Rpc))
						performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: false})
						continue
					}
				}
			}
			log.Log.Info("node check succeed", zap.String("rpc", v.Rpc))
			performance = append(performance, NewPerformance{Nodetype: v.Nodetype, Chain: v.Chain, Name: v.Name, Rpc: v.Rpc, Performance: true})
		}
	}
	return performance
}
