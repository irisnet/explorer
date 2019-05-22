package utils

import "github.com/irisnet/explorer/backend/logger"

type fn func()
type goroutinePool struct {
	queue chan fn
	exit  chan int
}

var pool *goroutinePool

func init() {
	pool = CreatePool(10)
	pool.Start()
}
func DefaultPool() *goroutinePool {
	return pool
}
func CreatePool(coreSize int) *goroutinePool {
	var queue = make(chan fn, coreSize)
	var exit = make(chan int, 1)
	return &goroutinePool{
		queue: queue,
		exit:  exit,
	}
}

func (pool *goroutinePool) Start() {
	go func() {
		for {
			select {
			case fn := <-pool.queue:
				fn()
			case <-pool.exit:
				logger.Info("goroutinePool exit")
				close(pool.queue)
				close(pool.exit)
				break
			}
		}
	}()
}

func (pool *goroutinePool) Stop() {
	pool.exit <- 1
}

func (pool *goroutinePool) Submit(fn func()) {
	go func() {
		task := func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Error("goroutinePool execute task error")
				}
			}()
			fn()
		}
		pool.queue <- task
	}()
}
