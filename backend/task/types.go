package task

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

var engine = Engine{
	task: []TimerTask{},
}

func init() {
	engine.AppendTask(TxNumGroupByDayTask{})
	engine.AppendTask(UpdateValidator{})
	engine.AppendTask(UpdateGovParams{})
	engine.AppendTask(UpdateValidatorIcons{})
	engine.AppendTask(UpdateAssetTokens{})
	engine.AppendTask(UpdateAssetGateways{})
	engine.AppendTask(ValidatorStaticInfo{})
}

type TimerTask interface {
	Start()
	Name() string
	DoTask() error
}

type Engine struct {
	task []TimerTask
}

func (e *Engine) Start() {
	if len(e.task) == 0 {
		return
	}
	for _, t := range e.task {
		var taskId = fmt.Sprintf("%s[%s]", t.Name(), utils.FmtTime(time.Now(), utils.DateFmtYYYYMMDD))
		logger.Info("timerTask begin to work", logger.String("taskId", taskId))
		go t.Start()
	}
}

func (e *Engine) AppendTask(task TimerTask) {
	e.task = append(e.task, task)
}

func Start() {
	engine.Start()
}
