package actor

type actor struct {
	id      int
	mailbox chan Task
	exitC   chan struct{}
}

func newActor(id int) *actor {
	return &actor{
		id:      id,
		mailbox: make(chan Task, 1024),
		exitC:   make(chan struct{}),
	}
}

func (p *actor) Start() {
	for {
		select {
		case f := <-p.mailbox:
			f()
		case <-p.exitC:
			return
		}
	}
}

func (p *actor) Shutdown() {
	close(p.exitC)
}

func (p *actor) Send(task Task) {
	p.mailbox <- task
}
