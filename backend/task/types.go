package task

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/robfig/cron/v3"
)

var (
	engine = Engine{
		task: []TimerTask{},
	}

	taskControlModel document.TaskControl
	tcService        TaskControlService

	cstZone = time.FixedZone("CST", 8*3600)
	// adapt multiple asset
	rewardsDenom = []string{"iris-atto"}
)

func init() {
	engine.AppendTask(UpdateValidator{})
	engine.AppendTask(UpdateGovParams{})
	engine.AppendTask(UpdateValidatorIcons{})
	engine.AppendTask(UpdateAssetTokens{})
	engine.AppendTask(ValidatorStaticInfo{})
	engine.AppendTask(UpdateProposalVoters{})
	engine.AppendTask(UpdateAccount{})

	taskControlMonitor := TaskControlMonitor{}
	taskControlMonitor.unlockAllTasks()
	engine.AppendTask(taskControlMonitor)
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

	// tasks manager by cron job
	c := cron.New(cron.WithLocation(cstZone))
	c.Start()

	txNumTask := TxNumGroupByDayTask{}
	txNumTask.init()
	c.AddFunc("01 0 * * *", func() {
		txNumTask.Start()
		new(UpdateValidatorIcons).Start()
	})
	c.AddFunc("59 23 * * *", func() {
		new(StaticDelegatorTask).Start()
		new(StaticValidatorTask).Start()
	})

	c.AddFunc("0 0 01 * *", func() { //每月1号0点0分
		delegatortask := new(StaticDelegatorByMonthTask)
		validatortask := new(StaticValidatorByMonthTask)
		delegatortask.Start()
		validatortask.SetAddressCoinMapData(delegatortask.AddressCoin, delegatortask.AddrPeriodCommission,
			delegatortask.AddrBeginCommission, delegatortask.AddrTerminalCommission)
		validatortask.Start()
	})
}
