package scheduling

import (
	"go.uber.org/zap"
	"math/rand"
	"supervisor/src/check"
	"supervisor/src/config"
	"supervisor/src/log"
	"supervisor/src/performanceupdate"
	"time"
)

// ScheduleNodeCheck schedule a node check in the configured interval(Granularity)
func ScheduleNodeCheck(c config.Config) {

	// schedule a node check at least once, and random in the configured interval
	granularity := c.Checkconfig.Granularity
	//must check at least once in the configured interval
	for {
		passduration := false
		go func() {
			log.Log.Info("The check interval is", zap.Int("second", granularity))
			time.Sleep(time.Duration(granularity) * time.Second)
			passduration = true
		}()
		go func() {
			duration := time.Duration(rand.Int63n(int64(granularity))) * time.Second
			time.Sleep(duration)
			log.Log.Info("Performing node check")
			p := check.CheckAll(c)
			performanceupdate.UpdatePerformance(p)
			log.Log.Info("Node check finished")
		}()
		//wait for pass duration
		for {
			if passduration {
				break
			}
		}
	}
}
