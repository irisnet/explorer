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
	UpdateProposalVoters{}.DoTask()
}

func TestHeartBeat(t *testing.T) {

	stop := make(chan bool)
	go HeartBeat("update_gov_params", stop)
	time.Sleep(11 * time.Second)
	stop <- true
}

//func TestHeartBeat1(t *testing.T) {
//
//	context0 := context.Background()
//	go HeartBeat1("update_gov_params", context0)
//	time.Sleep(11 * time.Second)
//}
//
//func TestHeartBeat2(t *testing.T) {
//
//	go HeartBeat2("update_gov_params")
//	time.Sleep(11 * time.Second)
//}
