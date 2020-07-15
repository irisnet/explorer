package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2"
	"time"
)

type (
	TaskControlMonitor struct {
		controlModel document.TaskControl
	}
	TaskControlService struct {
	}
	Callback func(func(string) chan bool) error
)

func (s TaskControlService) runTask(taskName string, timeInterval int, callback Callback) error {
	cronjob := false
	if _, ok := CronJob[taskName]; ok {
		cronjob = true
	}
	if notBeExec, err := s.assetTaskShouldNotBeExecuted(taskName, timeInterval, cronjob); err != nil {
		return fmt.Errorf("assetTaskShouldNotBeExecuted fail, taskName:%s, err:%s", taskName, err.Error())
	} else {
		if notBeExec {
			logger.Warn("task shouldn't be executed", logger.String("taskName", taskName))
			return nil
		}

		// lock task
		if err := s.lockTask(taskName, timeInterval); err != nil {
			return fmt.Errorf("lockTask fail, taskName:%s, err:%s", taskName, err.Error())
		} else {
			// do task
			var doTaskErr error
			if err := callback(HeartBeat); err != nil {
				doTaskErr = fmt.Errorf("doTask fail, taskName:%s, err:%s", taskName, err.Error())
			}

			// unlock task
			if err := s.unlockTask(taskName); err != nil {
				return fmt.Errorf("unLockTask fail, taskName:%s, err:%s", taskName, err.Error())
			}

			if doTaskErr != nil {
				return doTaskErr
			} else {
				return nil
			}
		}
	}
}

// task should not be executed as follow:
// - query record failed in task control table
// - task is executing by other server instance
// - now sub task latest update time less than task time interval
func (s TaskControlService) assetTaskShouldNotBeExecuted(taskName string, timeInterval int, isCronjob bool) (bool, error) {
	if v, err := taskControlModel.QueryOneByTaskName(taskName); err != nil {
		if err == mgo.ErrNotFound {
			// record not exist, need insert it
			record := document.TaskControl{
				TaskName:     taskName,
				TimeInterval: int64(timeInterval),
				IsInProcess:  false,
			}
			if err := taskControlModel.Save(record); err != nil {
				return true, err
			}
		} else {
			return true, err
		}
	} else {
		if v.IsInProcess {
			return true, nil
		}

		now := time.Now().Unix()
		if now-v.LatestExecTime < int64(timeInterval) && !isCronjob {
			return true, nil
		}
	}

	return false, nil
}

func (s TaskControlService) lockTask(taskName string, timeInterval int) error {
	instanceNo := conf.Get().Server.InstanceNo
	return taskControlModel.LockTaskControl(taskName, instanceNo, timeInterval)
}

func (s TaskControlService) unlockTask(taskName string) error {
	return taskControlModel.UnlockTaskControl(taskName)
}

func (s TaskControlMonitor) Name() string {
	return types.TaskConTrolMonitor
}
func (s TaskControlMonitor) Start() {
	timeInterval := conf.Get().Server.CronTimeControlTask
	taskName := s.Name()

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, s.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (s TaskControlMonitor) DoTask(fn func(string) chan bool) error {
	stop := fn(s.Name())
	defer HeartQuit(stop)
	runValue := true
	skip := 0
	limit := 20
	for runValue {
		res, err := s.controlModel.List(skip, limit)
		if err != nil {
			logger.Error(" list have error", logger.String("err", err.Error()))
		}
		s.checkAndChangeWorkOn(res)
		total := len(res)
		if total < limit {
			runValue = false
		} else {
			skip = skip + total
		}
	}
	return nil
}

func (s TaskControlMonitor) checkAndChangeWorkOn(list []document.TaskControl) {

	for _, one := range list {
		if err := s.CheckAndUpdate(one); err != nil {
			logger.Error("CheckAndUpdate have error",
				logger.String("serviceName", one.TaskName),
				logger.String("err", err.Error()))
		}
	}
	return
}

func (s TaskControlMonitor) CheckAndUpdate(one document.TaskControl) error {
	currentTimeInterval := time.Now().Unix() - one.LatestExecTime
	timeHeartBeat := int64(conf.Get().Server.CronTimeHeartBeat)

	if timeHeartBeat > currentTimeInterval {
		return nil
	}
	if currentTimeInterval >= 3*timeHeartBeat && one.IsInProcess && one.TaskName != types.TaskConTrolMonitor {
		one.IsInProcess = false
		one.LatestExecTime = time.Now().Unix()
		if err := s.controlModel.UpdateByPK(one); err != nil {
			return err
		}
	}
	return nil
}

//// unlock all tasks which task is unlocked
//func (s TaskControlMonitor) unlockAllTasks() error {
//	if err := taskControlModel.UnlockAllTasks(); err != nil {
//		if err != mgo.ErrNotFound {
//			return err
//		}
//	}
//
//	return nil
//}
