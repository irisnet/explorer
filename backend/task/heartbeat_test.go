package task

import (
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	time.Sleep(1 * time.Hour)
}

func TestStartTask(t *testing.T) {
	TxNumGroupByDayTask{}.Start()
	time.Sleep(1 * time.Hour)
}

func TestUpdateProposalVoters_DoTask(t *testing.T) {
	UpdateProposalVoters{}.DoTask(HeartBeat)
}

func TestHeartBeat(t *testing.T) {

	stop := make(chan bool)
	go HeartBeat("update_gov_params")
	time.Sleep(11 * time.Second)
	stop <- true
}


