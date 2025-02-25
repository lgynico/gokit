package syncx

type Wait struct {
	waitC chan struct{}
}

func NewWait() *Wait {
	return &Wait{
		waitC: make(chan struct{}),
	}
}

func (p *Wait) Wait() {
	<-p.waitC
}

func (p *Wait) Done() {
	close(p.waitC)
}
