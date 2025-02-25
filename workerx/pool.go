package workerx

import (
	"sync"

	"github.com/lgynico/gokit/mathx"
	"github.com/lgynico/gokit/runtimex"
)

type Pool struct {
	workers    sync.Map
	workerSize int
}

func NewPool(workerSize int) *Pool {
	return &Pool{
		workerSize: workerSize,
	}
}

func (p *Pool) Submit(task func()) {
	var (
		workerId   = mathx.RandInRange(0, p.workerSize-1)
		worker, ok = p.workers.LoadOrStore(workerId, NewSingle(64))
		single     = worker.(*Single)
	)

	if !ok {
		single.Start()
	}

	single.Submit(task)
}

func (p *Pool) SubmitBinding(workerId int, task func()) {
	workerId = p.ensureWorkerId(workerId)
	worker, ok := p.workers.LoadOrStore(workerId, NewSingle(64))
	single := worker.(*Single)
	if !ok {
		single.Start()
	}

	if single.Id() == runtimex.GetGoroutineId() {
		task()
	} else {
		single.Submit(task)
	}

}

func (p *Pool) Shutdown() {
	p.workers.Range(func(_, worker any) bool {
		worker.(*Single).Shutdown()
		return true
	})
}

func (p *Pool) ensureWorkerId(workerId int) int {
	return workerId % p.workerSize
}
