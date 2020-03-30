package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2"
	"time"
)

type (
	TaskControlService struct {
	}
	Callback func() error
)

func (s TaskControlService) runTask(taskName string, timeInterval int, callback Callback) error {
	if notBeExec, err := s.assetTaskShouldNotBeExecuted(taskName, timeInterval); err != nil {
		return fmt.Errorf("assetTaskShouldNotBeExecuted fail, taskName:%s, err:%s", taskName, err.Error())
	} else {
		if notBeExec {
			logger.Warn("task shouldn't be executed", logger.String("taskName", taskName))
			return nil
		}

		// lock task
		if err := s.lockTask(taskName); err != nil {
			return fmt.Errorf("lockTask fail, taskName:%s, err:%s", taskName, err.Error())
		} else {
			// do task
			var doTaskErr error
			if err := callback(); err != nil {
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
func (s TaskControlService) assetTaskShouldNotBeExecuted(taskName string, timeInterval int) (bool, error) {
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
		if now-v.LatestExecTime < int64(timeInterval) {
			return true, nil
		}
	}

	return false, nil
}

func (s TaskControlService) lockTask(taskName string) error {
	instanceNo := conf.Get().Server.InstanceNo
	return taskControlModel.LockTaskControl(taskName, instanceNo)
}

func (s TaskControlService) unlockTask(taskName string) error {
	return taskControlModel.UnlockTaskControl(taskName)
}
