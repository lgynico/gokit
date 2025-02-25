package workerx

import (
	"sync"

	"github.com/lgynico/gokit/runtimex"
)

type Single struct {
	ch          chan func()
	once        sync.Once
	goroutineId int
}

func NewSingle(buf int) *Single {
	return &Single{
		ch: make(chan func(), buf),
	}
}

func (p *Single) Start() {
	p.once.Do(func() {
		go func() {
			p.goroutineId = runtimex.GetGoroutineId()
			for f := range p.ch {
				f()
			}
		}()
	})
}

func (p *Single) Shutdown() {
	close(p.ch)
}

func (p *Single) Submit(task func()) {
	p.ch <- task
}

func (p *Single) Id() int {
	return p.goroutineId
}
