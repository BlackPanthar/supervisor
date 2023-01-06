package main

import (
	"supervisor/src/config"
	"supervisor/src/log"
	"supervisor/src/scheduling"
)

func main() {
	log.InitLog()
	c := config.GetConfig()
	scheduling.ScheduleNodeCheck(c)
}
