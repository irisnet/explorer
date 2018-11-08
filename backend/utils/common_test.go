package utils

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime/debug"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	_, ok := ParseInt("1")
	assert.True(t, ok)

	_, ok = ParseUint("-1")
	assert.False(t, ok)

	_, ok = ParseInt("sd")
	assert.False(t, ok)
}

func TestParseFloat(t *testing.T) {
	f, ok := RoundFloat(101.1234567, 4)
	assert.True(t, ok)
	fmt.Println(f)
}

func TestGetNodes(t *testing.T) {
	bz := GetNodes()
	fmt.Printf(string(bz))
}

func TestGetBalance(t *testing.T) {
	amt := GetBalance("faa10jv6pkdtjc39pwpjxnurqend0p09gphl3xg5yc")
	bz, _ := json.Marshal(amt)
	fmt.Printf(string(bz))
}

func TestDefer(t *testing.T) {
	for {
		TryCatch(func() {
			fmt.Println("do something...")
			time.Sleep(time.Second * 4)
			panic(1)
		}, func() {
			fmt.Println("panic...")
		}, func() {
			fmt.Println("finally")
		})
	}
}

func TestDefer2(t *testing.T) {
	i := f2()
	fmt.Println(i)
}

func f2() int {
	defer fmt.Println("finish")
	i := 0
	panic(1)
	i++
	return i
}

func TryCatch(try func(), catch func(), finally func()) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			catch()
		}
		finally()
	}()
	try()
}

type TaskTicker struct {
	*time.Ticker
	Cmd func()
}

func (task *TaskTicker) Start() {
	for {
		select {
		case t := <-task.C:
			fmt.Print(t.String() + " : ")
			task.Cmd()
		}
	}
}

func (task *TaskTicker) Stop() {
	task.Ticker.Stop()
}

type Cron struct {
	taskQueue []*TaskTicker
}

func New() *Cron {
	var queue []*TaskTicker
	return &Cron{
		taskQueue: queue,
	}
}

func (c *Cron) AddFunc(spec time.Duration, cmd func()) {
	t := &TaskTicker{
		Ticker: time.NewTicker(spec),
		Cmd:    cmd,
	}
	c.taskQueue = append(c.taskQueue, t)
}

func (c *Cron) Start() {
	for _, task := range c.taskQueue {
		go task.Start()
	}
}

func (c *Cron) Stop() {
	for _, task := range c.taskQueue {
		go task.Stop()
	}
}

var i = 0

func TestTicker(t *testing.T) {

	cron := New()
	cron.AddFunc(1*time.Second, testTimer1)
	cron.AddFunc(1*time.Second, testTimer2)

	cron.Start()

	for {
		select {
		case <-time.After(20 * time.Second):
			cron.Stop()
			fmt.Println(i)
			return
		}
	}

}

func testTimer2() {
	i++
	fmt.Println("testTimer2")
	time.Sleep(time.Second * 10)
}

func testTimer1() {
	fmt.Println("================testTimer1")
	fmt.Println(i)
}
