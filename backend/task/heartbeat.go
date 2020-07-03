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
		doc, err := taskcontrol.QueryOneByTaskName(taskName)
		if err != nil {
			logger.Error("QueryOneByTaskName have error",
				logger.String("taskName", taskName),
				logger.String("tableName", taskcontrol.Name()),
				logger.String("err", err.Error()))
			return
		}
		curtime := time.Now().Unix()
		if curtime > doc.LatestExecTime+timeInterval {
			doc.LatestExecTime = time.Now().Unix()
			if err := taskcontrol.UpdateByPK(doc); err != nil {
				logger.Error("update segment have error",
					logger.String("taskName", taskName),
					logger.String("tableName", doc.Name()),
					logger.String("err", err.Error()))
				return
			}
		}
	}
	doTask()
	for {
		select {
		case <-ticker.C:
			doTask()
			logger.Info("HeartBeat is ok", logger.String("taskName", taskName))
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

//
//func HeartBeat1(taskName string, ctx context.Context) {
//	var taskcontrol document.TaskControl
//	timeInterval := int64(conf.Get().Server.CronTimeHeartBeat)
//	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
//	defer ticker.Stop()
//	for {
//		select {
//		case <-ticker.C:
//			doc, err := taskcontrol.QueryOneByTaskName(taskName)
//			if err != nil {
//				logger.Error("QueryOneByTaskName have error",
//					logger.String("taskName", taskName),
//					logger.String("tableName", taskcontrol.Name()),
//					logger.String("err", err.Error()))
//				return
//			}
//			curtime := time.Now().Unix()
//			if curtime > doc.LatestExecTime+timeInterval {
//				doc.LatestExecTime = time.Now().Unix()
//				if err := taskcontrol.UpdateByPK(doc); err != nil {
//					logger.Error("update segment have error",
//						logger.String("taskName", taskName),
//						logger.String("tableName", doc.Name()),
//						logger.String("err", err.Error()))
//					return
//				}
//			}
//			logger.Info("HeartBeat is ok", logger.String("taskName", taskName))
//		case <-ctx.Done():
//			logger.Info("HeartBeat Quit...", logger.String("taskName", taskName))
//			return
//
//		}
//
//	}
//}
//
//func HeartBeat2(taskName string) {
//	var taskcontrol document.TaskControl
//	timeInterval := int64(conf.Get().Server.CronTimeHeartBeat)
//	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
//	stop := make(chan os.Signal)
//	signal.Notify(stop, os.Interrupt)
//	defer ticker.Stop()
//	for {
//		select {
//		case <-ticker.C:
//			doc, err := taskcontrol.QueryOneByTaskName(taskName)
//			if err != nil {
//				logger.Error("QueryOneByTaskName have error",
//					logger.String("taskName", taskName),
//					logger.String("tableName", taskcontrol.Name()),
//					logger.String("err", err.Error()))
//				return
//			}
//			curtime := time.Now().Unix()
//			if curtime > doc.LatestExecTime+timeInterval {
//				doc.LatestExecTime = time.Now().Unix()
//				if err := taskcontrol.UpdateByPK(doc); err != nil {
//					logger.Error("update segment have error",
//						logger.String("taskName", taskName),
//						logger.String("tableName", doc.Name()),
//						logger.String("err", err.Error()))
//					return
//				}
//			}
//			logger.Info("HeartBeat is ok", logger.String("taskName", taskName))
//		case <-stop:
//			logger.Info("HeartBeat Quit...", logger.String("taskName", taskName))
//			return
//
//		}
//
//	}
//}