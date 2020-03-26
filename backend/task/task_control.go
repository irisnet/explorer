package task

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2"
	"time"
)

type TaskControlService struct {
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
	return taskControlModel.LockTaskControl(taskName)
}

func (s TaskControlService) unlockTask(taskName string) error {
	return taskControlModel.UnlockTaskControl(taskName)
}
