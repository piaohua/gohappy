package main

import (
	"fmt"

	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct{ Who string }
type parentActor struct{}

func (state *parentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		props := actor.FromProducer(newChildActor)
		child := context.Spawn(props)
		child.Tell(msg)
	}
}

func newParentActor() actor.Actor {
	return &parentActor{}
}

type childActor struct{}

func (state *childActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Starting, initialize actor here")
	case *actor.Stopping:
		fmt.Println("Stopping, actor is about to shut down")
	case *actor.Stopped:
		fmt.Println("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		fmt.Println("Restarting, actor is about to restart")
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
		panic("Ouch")
	}
}

func newChildActor() actor.Actor {
	return &childActor{}
}

func main() {
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		FromProducer(newParentActor).
		WithSupervisor(supervisor)

	pid := actor.Spawn(props)
	pid.Tell(&hello{Who: "Roger"})

	console.ReadLine()
}
