package task

import (
	"time"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/conf"
	"os"
	"os/signal"
)

func HeartBeat(taskName string) chan bool {
	//do heartbeat
	stop := make(chan bool)
	go heartBeat(taskName, stop)
	return stop
}

func HeartQuit(stop chan bool) {
	stop <- true
}

// heart beat of task not be executed as follow:
// - query task failed in task control table
// - task is not running
// - now sub task latest update time less than heart beat time interval
func heartBeat(taskName string, stop chan bool) {
	logger.Info("HeartBeat Starting...", logger.String("taskName", taskName))
	signalstop := make(chan os.Signal)
	signal.Notify(signalstop, os.Interrupt)
	var taskcontrol document.TaskControl
	timeInterval := int64(conf.Get().Server.CronTimeHeartBeat)
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	defer ticker.Stop()

	doTask := func() {
		err := taskcontrol.UpdateByTaskName(taskName)
		if err != nil {
			logger.Error("UpdateByTaskName have error",
				logger.String("taskName", taskName),
				logger.String("tableName", taskcontrol.Name()),
				logger.String("err", err.Error()))
			return
		}
	}
	doTask()
	for {
		select {
		case <-ticker.C:
			doTask()
			//logger.Info("HeartBeat is ok", logger.String("taskName", taskName))
		case <-stop:
			close(stop)
			logger.Info("HeartBeat Quit...", logger.String("taskName", taskName))
			return
		case <-signalstop:
			close(signalstop)
			logger.Info("Interrupt HeartBeat Quit...", logger.String("taskName", taskName))
			return

		}

	}
}
