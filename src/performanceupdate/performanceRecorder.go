package performanceupdate

import (
	"encoding/json"
	"go.uber.org/zap"
	"os"
	"supervisor/src/check"
	"supervisor/src/log"
)

type NodePerformance struct {
	Nodetype    string
	Chain       string
	Name        string
	Rpc         string
	Performance []bool
	SLA         float64
}

// UpdatePerformance Get the performance of the node from last check, and update the performance.json.
func UpdatePerformance(performances []check.NewPerformance) {
	//if file is empty or not exist ,write new performance
	if f, err := os.Stat("./performance.json"); os.IsNotExist(err) || f.Size() == 0 {
		log.Log.Info("performance.json not exist or empty,create new performance.json")
		var nodePerformances []NodePerformance
		for _, performance := range performances {
			if performance.Performance == true {
				nodePerformances = append(nodePerformances, NodePerformance{Nodetype: performance.Nodetype, Chain: performance.Chain, Name: performance.Name, Rpc: performance.Rpc, Performance: []bool{performance.Performance}, SLA: 1})
			}
			if performance.Performance == false {
				nodePerformances = append(nodePerformances, NodePerformance{Nodetype: performance.Nodetype, Chain: performance.Chain, Name: performance.Name, Rpc: performance.Rpc, Performance: []bool{performance.Performance}, SLA: 0})
			}
		}
		//create and write the file
		file, err := os.Create("./performance.json")
		if err != nil {
			log.Log.Error("create performance.json failed", zap.Error(err))
			return
		}
		if err := json.NewEncoder(file).Encode(nodePerformances); err != nil {
			log.Log.Error("Update node performance error,write to file failed", zap.Error(err))
			return
		}
		log.Log.Info("Update node performance success")
		return
	}

	//if file is not empty,read old performance and update
	log.Log.Info("performance.json exist and not empty,will update performance.json")
	var nodePerformances []NodePerformance
	file, err := os.Open("./performance.json")
	if err != nil {
		log.Log.Error("open performance.json failed", zap.Error(err))
		return
	}
	if err := json.NewDecoder(file).Decode(&nodePerformances); err != nil {
		log.Log.Error("Update node performance error,decode file failed", zap.Error(err))
		return
	}
	//if node is new,update it

	for _, newperformance := range performances {
		found := false
		for i, nodePerformance := range nodePerformances {
			if nodePerformance.Rpc == newperformance.Rpc {
				nodePerformances[i].Performance = append(nodePerformances[i].Performance, newperformance.Performance)
				//calculate SLA
				var trueCount int
				for _, performance := range nodePerformances[i].Performance {
					if performance == true {
						trueCount++
					}
				}
				nodePerformances[i].SLA = float64(trueCount) / float64(len(nodePerformances[i].Performance))
				found = true
				break
			}
		}
		if found == false {
			if newperformance.Performance == true {
				nodePerformances = append(nodePerformances, NodePerformance{Nodetype: newperformance.Nodetype, Chain: newperformance.Chain, Name: newperformance.Name, Rpc: newperformance.Rpc, Performance: []bool{newperformance.Performance}, SLA: 1})
			}
			if newperformance.Performance == false {
				nodePerformances = append(nodePerformances, NodePerformance{Nodetype: newperformance.Nodetype, Chain: newperformance.Chain, Name: newperformance.Name, Rpc: newperformance.Rpc, Performance: []bool{newperformance.Performance}, SLA: 0})
			}

		}
	}
	//update and write to the file
	newfile, err := os.Create("./performance.json")
	if err != nil {
		log.Log.Error("create performance.json failed", zap.Error(err))
		return
	}
	if err := json.NewEncoder(newfile).Encode(nodePerformances); err != nil {
		log.Log.Error("Update node performance error,write to file failed", zap.Error(err))
		return
	}
	log.Log.Info("Update node performance success")
}
