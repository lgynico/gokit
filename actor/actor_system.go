package actor

type (
	ActorSystem struct {
		actors map[int]*actor
	}
)

func NewActorSystem(actorCount int) *ActorSystem {
	actors := make(map[int]*actor, actorCount)
	for i := 0; i < actorCount; i++ {
		actors[i] = newActor(i)
	}

	return &ActorSystem{
		actors: actors,
	}
}

func (p *ActorSystem) Start() {
	for _, actor := range p.actors {
		go actor.Start()
	}
}

func (p *ActorSystem) Shutdown() {
	for _, actor := range p.actors {
		actor.Shutdown()
	}
}

func (p *ActorSystem) Send(id int, task Task) {
	id %= len(p.actors)
	if actor, ok := p.actors[id]; ok {
		actor.Send(task)
	}
}

// func (p *ActorSystem) SendSync(name string, id int, task Task) chan any {
// 	ch := make(chan any)

// 	return ch
// }
